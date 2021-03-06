package template

import (
	"bytes"
	"strings"
	"testing"
)

func TestNameFromPath(t *testing.T) {

	if nameFromPath("test.grm") != "test" {
		t.Fatal("'test.grm' != 'test'")
	}
	if nameFromPath("test.txt.grm") != "test" {
		t.Fatal("'test.txt.grm' != 'test'")
	}
	if nameFromPath("./a/test.grm") != "test" {
		t.Fatal("'./a/test.grm' != 'test'")
	}
	if nameFromPath("/a/b/test.grm") != "test" {
		t.Fatal("'/a/b/test.grm' != 'test'")
	}
}

func TestParseImports(t *testing.T) {

	tmpl := `{{/* Regular comment */}}
{{/* +import a "./a" */}}
{{/* Another comment */}}
{{/* +import b "./b" */}}
{{/* +import "./c" */}}
<div>Hello World</div>
{{/* Ignored import */}}
{{/* +import "./d" */}}
`
	imps, err := parseImports(tmpl)
	if err != nil {
		t.Fatal(err)
	}
	if imps["a"] != "./a.grm" {
		t.Fatal("import 'a' missing")
	}
	if imps["b"] != "./b.grm" {
		t.Fatal("import 'b' missing")
	}
	if imps["c"] != "./c.grm" {
		t.Fatal("import 'c' missing")
	}
	if _, ok := imps["d"]; ok {
		t.Fatal("import 'd' not ignored")
	}
}

func TestTemplate1(t *testing.T) {

	tmpl, err := New(nil, false).ParseFile("tmpl1", "./test/tmpl1.grm")
	if err != nil {
		t.Fatal(err)
	}

	data := map[string]string{
		"Greeting": "Hello World",
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, data); err != nil {
		t.Fatal(err)
	}

	result := `<html>
    <body>Hello World</body>
</html>
`
	if bytes.Compare(buf.Bytes(), []byte(result)) != 0 {
		t.Fatal("template does not match expected:\n", result)
	}
}

func TestTemplate2(t *testing.T) {

	tmpl, err := New(nil, false).ParseFile("tmpl2", "./test/tmpl2.grm")
	if err != nil {
		t.Fatal(err)
	}

	data := map[string]string{
		"Greeting": "Hello World",
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, data); err != nil {
		t.Fatal(err)
	}

	result := `<html>
    <body>Hello World</body>
</html>
`
	if bytes.Compare(buf.Bytes(), []byte(result)) != 0 {
		t.Fatalf("template result: `%s`\ndoes not match expected: `%s`\n", buf.String(), result)
	}
}

func TestTemplate3(t *testing.T) {

	funcs := FuncMap{
		"upper": strings.ToUpper,
	}

	tmpl, err := New(funcs, false).ParseFile("tmpl3", "./test/tmpl3.grm")
	if err != nil {
		t.Fatal(err)
	}

	data := map[string]string{
		"Greeting": "Hello World",
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, data); err != nil {
		t.Fatal(err)
	}

	result := `<html>
    <body>HELLO WORLD</body>
</html>
`
	if bytes.Compare(buf.Bytes(), []byte(result)) != 0 {
		t.Fatalf("template result: `%s`\ndoes not match expected: `%s`\n", buf.String(), result)
	}
}
