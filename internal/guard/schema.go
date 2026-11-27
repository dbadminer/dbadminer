package guard

import (
	"errors"
	"fmt"
	"strings"
)

// SchemaSQL reports whether sql stays inside schema.
// An empty schema allows every statement. A qualified relation (schema.table)
// and any three-part name (schema.table.column) must use schema. USE, ATTACH,
// DETACH, search_path changes, and CREATE, DROP, or ALTER DATABASE or SCHEMA
// are rejected because they leave the pinned schema for later statements.
func SchemaSQL(sql, schema string) error {
	if strings.TrimSpace(schema) == "" {
		return nil
	}
	if err := validateStatement(sql); err != nil {
		return err
	}
	stripped, err := stripSQL(sql)
	if err != nil {
		return err
	}
	atoms, err := scanAtoms(stripped)
	if err != nil {
		return err
	}
	if err := rejectSessionChange(atoms); err != nil {
		return err
	}
	for i := 0; i < len(atoms); i++ {
		a := atoms[i]
		if a.kind == "ident" {
			parts, _ := readQual(atoms, i)
			if len(parts) >= 3 {
				if err := matchSchema(parts[0], schema); err != nil {
					return err
				}
			}
		}
		if a.kind != "ident" || a.quoted {
			continue
		}
		if _, ok := relationKeywords[a.text]; !ok {
			continue
		}
		if skipRelationKeyword(atoms, i) {
			continue
		}
		next, err := consumeRelations(atoms, i+1, schema)
		if err != nil {
			return err
		}
		if next > i+1 {
			i = next - 1
		}
	}
	return nil
}

type atom struct {
	kind   string
	text   string
	quoted bool
}

func scanAtoms(s string) ([]atom, error) {
	var out []atom
	for i := 0; i < len(s); {
		c := s[i]
		if c == ' ' || c == '\n' || c == '\t' || c == '\r' || c == '\f' || c == '\v' {
			i++
			continue
		}
		switch c {
		case ',':
			out = append(out, atom{kind: "comma"})
			i++
			continue
		case '(':
			out = append(out, atom{kind: "lparen"})
			i++
			continue
		case ')':
			out = append(out, atom{kind: "rparen"})
			i++
			continue
		case '.':
			out = append(out, atom{kind: "dot"})
			i++
			continue
		case '"', '`':
			text, next, err := unquoteIdent(s, i, c)
			if err != nil {
				return nil, err
			}
			out = append(out, atom{kind: "ident", text: text, quoted: true})
			i = next
			continue
		}
		if isIdentStart(c) {
			j := i + 1
			for j < len(s) && isIdentByte(s[j]) {
				j++
			}
			out = append(out, atom{kind: "ident", text: strings.ToLower(s[i:j])})
			i = j
			continue
		}
		i++
	}
	return out, nil
}

func unquoteIdent(s string, i int, quote byte) (string, int, error) {
	var b strings.Builder
	i++
	for i < len(s) {
		if s[i] == quote {
			if i+1 < len(s) && s[i+1] == quote {
				b.WriteByte(quote)
				i += 2
				continue
			}
			return b.String(), i + 1, nil
		}
		b.WriteByte(s[i])
		i++
	}
	return "", 0, errors.New("dbadminer: unterminated string")
}

func isIdentStart(c byte) bool {
	return c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func rejectSessionChange(atoms []atom) error {
	i := firstKeyword(atoms)
	if i < 0 {
		return nil
	}
	switch atoms[i].text {
	case "use", "attach", "detach":
		return errors.New("dbadminer: statement would leave the configured schema")
	case "set", "reset":
		for _, a := range atoms[i+1:] {
			if a.kind == "ident" && !a.quoted && a.text == "search_path" {
				return errors.New("dbadminer: statement would leave the configured schema")
			}
		}
	case "create", "drop", "alter":
		j := i + 1
		for j < len(atoms) && isPrefixNoise(atoms[j]) {
			j++
		}
		if j < len(atoms) && atoms[j].kind == "ident" && !atoms[j].quoted &&
			(atoms[j].text == "database" || atoms[j].text == "schema") {
			return errors.New("dbadminer: statement would leave the configured schema")
		}
	}
	return nil
}

func firstKeyword(atoms []atom) int {
	for i, a := range atoms {
		if a.kind == "ident" && !a.quoted {
			return i
		}
	}
	return -1
}

func isPrefixNoise(a atom) bool {
	if a.kind != "ident" || a.quoted {
		return false
	}
	switch a.text {
	case "or", "replace", "if", "not", "exists", "only", "temporary", "temp", "unlogged":
		return true
	default:
		return false
	}
}

func skipRelationKeyword(atoms []atom, i int) bool {
	word := atoms[i].text
	if word == "update" && prevKeyword(atoms, i) == "for" {
		return true
	}
	if word == "into" {
		next := nextIdent(atoms, i+1)
		if next == "outfile" || next == "dumpfile" {
			return true
		}
	}
	if word == "on" {
		next := nextIdent(atoms, i+1)
		if next == "delete" || next == "update" || next == "conflict" {
			return true
		}
		if prevKeyword(atoms, i) == "join" {
			return true
		}
	}
	return false
}

func prevKeyword(atoms []atom, i int) string {
	for j := i - 1; j >= 0; j-- {
		a := atoms[j]
		if a.kind == "ident" && !a.quoted {
			if _, ok := keywords[a.text]; ok {
				return a.text
			}
		}
	}
	return ""
}

func nextIdent(atoms []atom, i int) string {
	if i < len(atoms) && atoms[i].kind == "ident" && !atoms[i].quoted {
		return atoms[i].text
	}
	return ""
}

func consumeRelations(atoms []atom, i int, schema string) (int, error) {
	for {
		i = skipNoise(atoms, i)
		if i >= len(atoms) || atoms[i].kind != "ident" {
			return i, nil
		}
		if !atoms[i].quoted && (isBoundary(atoms[i].text) || isRelationKeyword(atoms[i].text)) {
			return i, nil
		}
		parts, next := readQual(atoms, i)
		if len(parts) >= 2 {
			if err := matchSchema(parts[0], schema); err != nil {
				return next, err
			}
		}
		i = skipAlias(atoms, next)
		if i < len(atoms) && atoms[i].kind == "comma" {
			i++
			continue
		}
		return i, nil
	}
}

func skipNoise(atoms []atom, i int) int {
	for i < len(atoms) && isPrefixNoise(atoms[i]) {
		i++
	}
	return i
}

func readQual(atoms []atom, i int) ([]string, int) {
	if i >= len(atoms) || atoms[i].kind != "ident" {
		return nil, i
	}
	parts := []string{atoms[i].text}
	i++
	for i+1 < len(atoms) && atoms[i].kind == "dot" && atoms[i+1].kind == "ident" {
		parts = append(parts, atoms[i+1].text)
		i += 2
	}
	return parts, i
}

func skipAlias(atoms []atom, i int) int {
	if i >= len(atoms) || atoms[i].kind != "ident" {
		return i
	}
	if atoms[i].quoted {
		return i + 1
	}
	if atoms[i].text == "as" {
		i++
		if i < len(atoms) && atoms[i].kind == "ident" {
			return i + 1
		}
		return i
	}
	if isBoundary(atoms[i].text) || isRelationKeyword(atoms[i].text) || isPrefixNoise(atoms[i]) {
		return i
	}
	return i + 1
}

func matchSchema(name, schema string) error {
	if strings.EqualFold(name, schema) {
		return nil
	}
	if len(name) > 63 {
		return errors.New("dbadminer: statement references another schema")
	}
	return fmt.Errorf("dbadminer: statement references schema %q", name)
}

var relationKeywords = map[string]struct{}{
	"from": {}, "join": {}, "straight_join": {}, "into": {}, "update": {},
	"table": {}, "view": {}, "references": {}, "truncate": {}, "function": {},
	"procedure": {}, "sequence": {}, "trigger": {}, "on": {}, "to": {},
}

var keywords = map[string]struct{}{
	"add": {}, "after": {}, "all": {}, "alter": {}, "and": {}, "as": {},
	"attach": {}, "before": {}, "begin": {}, "by": {}, "case": {}, "column": {},
	"conflict": {}, "constraint": {}, "create": {}, "cross": {}, "database": {},
	"delete": {}, "detach": {}, "drop": {}, "dumpfile": {}, "each": {}, "else": {},
	"end": {}, "exists": {}, "for": {}, "foreign": {}, "from": {}, "full": {},
	"function": {}, "grant": {}, "group": {}, "having": {}, "if": {}, "index": {},
	"inner": {}, "insert": {}, "instead": {}, "into": {}, "join": {}, "key": {},
	"left": {}, "limit": {}, "natural": {}, "not": {}, "offset": {}, "on": {},
	"only": {}, "or": {}, "order": {}, "outer": {}, "outfile": {}, "primary": {},
	"procedure": {}, "references": {}, "rename": {}, "replace": {}, "reset": {},
	"returning": {}, "right": {}, "row": {}, "schema": {}, "search_path": {},
	"select": {}, "sequence": {}, "set": {}, "show": {}, "statement": {},
	"straight_join": {}, "table": {}, "temporary": {}, "then": {}, "to": {},
	"trigger": {}, "truncate": {}, "union": {}, "unique": {}, "unlogged": {},
	"update": {}, "use": {}, "using": {}, "values": {}, "view": {}, "when": {},
	"where": {}, "window": {}, "with": {},
}

func isRelationKeyword(word string) bool {
	_, ok := relationKeywords[word]
	return ok
}

func isBoundary(word string) bool {
	switch word {
	case "where", "set", "group", "order", "limit", "having", "union", "except",
		"intersect", "window", "returning", "values", "select", "using", "and",
		"or", "when", "then", "else", "end", "left", "right", "inner", "outer",
		"cross", "full", "natural", "offset", "fetch", "by", "add", "column",
		"for", "before", "after", "instead", "each", "row", "statement", "begin",
		"grant", "with", "recursive", "primary", "unique", "constraint", "foreign",
		"key", "index", "show", "explain", "describe", "rename":
		return true
	default:
		return false
	}
}
