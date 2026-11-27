package guard

import (
	"strings"
	"testing"
)

func TestIsReadOnlySQL(t *testing.T) {
	cases := []struct {
		sql string
		ok  bool
	}{
		{sql: "select 1", ok: true},
		{sql: "  SeLeCt 1;  ", ok: true},
		{sql: "select 1; select 2", ok: false},
		{sql: "select 1; drop table t", ok: false},
		{sql: "select ';'", ok: true},
		{sql: "select 'delete from t'", ok: true},
		{sql: "/* drop table t */ select 1", ok: true},
		{sql: "-- drop\nselect 1", ok: true},
		{sql: "select 1 /*! drop table t */", ok: false},
		{sql: "select 1 /*!12345 drop table t */", ok: false},
		{sql: "with x as (select 1) select * from x", ok: true},
		{sql: "with x as (delete from t returning id) select id from x", ok: false},
		{sql: "delete from t", ok: false},
		{sql: "insert into t values (1)", ok: false},
		{sql: "update t set a = 1", ok: false},
		{sql: "drop table t", ok: false},
		{sql: "select * from t for update", ok: false},
		{sql: "select replace(name, 'a', 'b') from t", ok: true},
		{sql: "replace into t values (1)", ok: false},
		{sql: "explain select 1", ok: true},
		{sql: "explain analyze select 1", ok: true},
		{sql: "explain analyze delete from t", ok: false},
		{sql: "select a into outfile '/tmp/x' from t", ok: false},
		{sql: "show create table t", ok: true},
		{sql: "show databases", ok: true},
		{sql: "values (1)", ok: true},
		{sql: "table t", ok: true},
		{sql: "pragma table_info(t)", ok: true},
		{sql: "pragma journal_mode=delete", ok: false},
		{sql: "pragma writable_schema=on", ok: false},
		{sql: "select $tag$; drop table t;$tag$", ok: true},
		{sql: "select 'unterminated", ok: false},
		{sql: "", ok: false},
		{sql: "/* only a comment */", ok: false},
		{sql: "select users.set from t", ok: true},
		{sql: "# mysql comment\nselect 1", ok: true},
		{sql: "select 1\x00", ok: false},
	}
	for _, tc := range cases {
		ok, err := isReadOnlySQL(tc.sql)
		if ok != tc.ok {
			t.Errorf("isReadOnlySQL(%q) = %v, %v; want ok=%v", tc.sql, ok, err, tc.ok)
		}
		if ok && err != nil {
			t.Errorf("isReadOnlySQL(%q) ok with error %v", tc.sql, err)
		}
	}
}

func TestSchemaSQL(t *testing.T) {
	allow := []string{
		"select 1",
		"select * from users",
		"select users.id from users",
		"select * from main.users",
		"select * from main.users u where u.id = 1",
		"select * from users u join posts p on u.id = p.user_id",
		"select * from users, posts",
		"select * from users u, posts p",
		"select * from (select * from main.users) s",
		"select * from users for update",
		"create table users (id integer)",
		"create table if not exists main.users (id integer)",
		`create table "main"."users" (id integer)`,
		"create table `main`.`users` (id integer)",
		"insert into users values (1)",
		"insert into main.users values (1)",
		"update users set name = 'a'",
		"create view v as select * from users",
		"create view main.v as select * from main.users",
		"create index i on main.users (id)",
		"alter table main.users add constraint fk foreign key (id) references main.parent (id)",
		"select * from users where name = 'test.evil'",
		"select * from users -- from test.evil\n",
		"set names utf8",
		"extract(year from created_at)",
	}
	deny := []string{
		"create table test.users (id integer)",
		"create table `test`.`users` (id int)",
		`create table "test"."users" (id integer)`,
		"select * from test.users",
		"select * from users u join test.posts p on u.id = p.user_id",
		"select * from users, test.posts",
		"select * from (select * from test.users) s",
		"insert into test.users values (1)",
		"update test.users set a = 1",
		"delete from test.users",
		"drop table test.users",
		"alter table test.users add column a int",
		"create view main.v as select * from test.users",
		"select test.users.email",
		"use test",
		"use main",
		"attach database 'x' as other",
		"create database test",
		"drop schema other",
		"set search_path to other",
		"create index i on test.users (id)",
		"alter table main.users add constraint fk foreign key (id) references test.parent (id)",
		"rename table users to test.users",
		"select 1 /*! from test.users */",
	}
	for _, sql := range allow {
		if err := SchemaSQL(sql, "main"); err != nil {
			t.Errorf("SchemaSQL(%q) = %v", sql, err)
		}
	}
	for _, sql := range deny {
		if err := SchemaSQL(sql, "main"); err == nil {
			t.Errorf("SchemaSQL(%q) allowed", sql)
		}
	}
	if err := SchemaSQL("create table test.users (id integer)", ""); err != nil {
		t.Errorf("empty schema: %v", err)
	}
}

func TestReturnsRows(t *testing.T) {
	cases := []struct {
		sql string
		ok  bool
	}{
		{"select 1", true},
		{"with x as (select 1) select * from x", true},
		{"explain select 1", true},
		{"values (1)", true},
		{"table t", true},
		{"show tables", true},
		{"pragma table_info(t)", true},
		{"delete from t", false},
		{"insert into t values (1)", false},
		{"", false},
	}
	for _, tc := range cases {
		if got := returnsRows(tc.sql); got != tc.ok {
			t.Errorf("returnsRows(%q) = %v, want %v", tc.sql, got, tc.ok)
		}
	}
}

func TestIsLoopbackAddr(t *testing.T) {
	ok := []string{"127.0.0.1", "127.0.0.1:8080", "localhost", "localhost:80", "::1", "[::1]:8080", "LOCALHOST:9"}
	bad := []string{"", "0.0.0.0:8080", "8.8.8.8", "example.com", "example.com:80", "[::]:8080", "127.0.0.1.evil.com"}
	for _, addr := range ok {
		if !isLoopbackAddr(addr) {
			t.Errorf("%q should be loopback", addr)
		}
	}
	for _, addr := range bad {
		if isLoopbackAddr(addr) {
			t.Errorf("%q should not be loopback", addr)
		}
	}
}

func TestCheckRequestHost(t *testing.T) {
	if err := checkRequestHost("127.0.0.1:8080", ""); err != nil {
		t.Fatal(err)
	}
	if err := checkRequestHost("127.0.0.1:8080", "http://127.0.0.1:8080"); err != nil {
		t.Fatal(err)
	}
	if err := checkRequestHost("evil.example:8080", ""); err == nil {
		t.Fatal("expected public host to fail")
	}
	if err := checkRequestHost("127.0.0.1:8080", "http://evil.example"); err == nil {
		t.Fatal("expected foreign origin to fail")
	}
	if err := checkRequestHost("127.0.0.1:8080", "http://127.0.0.1:9"); err == nil {
		t.Fatal("expected origin port mismatch to fail")
	}
	if err := checkRequestHost("localhost:8080", "http://127.0.0.1:8080"); err == nil {
		t.Fatal("expected different loopback names to fail")
	}
}

func FuzzIsReadOnlySQL(f *testing.F) {
	seeds := []string{
		"select 1",
		"select 1; drop table t",
		"/* c */ select 'delete'",
		"with x as (select 1) select * from x",
		"delete from t",
		"select 1;/*",
		"select 'unterminated",
		";;;",
		"SeLeCt 1",
		"explain select 1",
		"pragma table_info(t)",
		"pragma journal_mode=delete",
		"show create table t",
		"select replace(a,'b','c')",
		"--\nselect 1",
		"select $tag$ drop table t $tag$",
		"select 'drop table t'",
		"insert into t values (1)",
		"select 1 union select 2",
	}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, q string) {
		ok, err := isReadOnlySQL(q)
		if ok && err != nil {
			t.Fatalf("ok with error: %v", err)
		}
		if !ok {
			return
		}
		ok2, err2 := isReadOnlySQL(q)
		if !ok2 || err2 != nil {
			t.Fatal("not stable")
		}
		stripped, err := stripSQL(q)
		if err != nil {
			t.Fatal(err)
		}
		n := 0
		for _, part := range strings.Split(stripped, ";") {
			if strings.TrimSpace(part) != "" {
				n++
			}
		}
		if n > 1 {
			t.Fatal("read-only allowed multiple statements")
		}
	})
}
