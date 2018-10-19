package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	starlark "github.com/google/skylark"
	starlarktest "github.com/google/skylark/skylarktest"
	dataset "github.com/qri-io/dataset"
)

func TestAsString(t *testing.T) {
	cases := []struct {
		in       starlark.Value
		got, err string
	}{
		{starlark.String("foo"), "foo", ""},
		{starlark.String("\"foo'"), "\"foo'", ""},
		{starlark.Bool(true), "", "invalid syntax"},
	}

	for i, c := range cases {
		got, err := AsString(c.in)
		if !(err == nil && c.err == "" || err != nil && err.Error() == c.err) {
			t.Errorf("case %d error mismatch. expected: '%s', got: '%s'", i, c.err, err)
			continue
		}

		if c.got != got {
			t.Errorf("case %d. expected: '%s', got: '%s'", i, c.got, got)
		}
	}
}

func TestNewModule(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Date", "Mon, 01 Jun 2000 00:00:00 GMT")
		w.Write([]byte(`{"hello":"world"}`))
	}))
	starlark.Universe["test_server_url"] = starlark.String(ts.URL)

	ds := &dataset.Dataset{
		Transform: &dataset.Transform{
			Syntax: "starlark",
			Config: map[string]interface{}{
				"foo": "bar",
			},
		},
	}
	thread := &starlark.Thread{Load: newLoader(ds)}
	starlarktest.SetReporter(thread, t)

	// Execute test file
	_, err := starlark.ExecFile(thread, "testdata/test.sky", nil, nil)
	if err != nil {
		t.Error(err)
	}
}

// load implements the 'load' operation as used in the evaluator tests.
func newLoader(ds *dataset.Dataset) func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	return func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
		switch module {
		case ModuleName:
			return LoadModule()
		case "assert.sky":
			return starlarktest.LoadAssertModule()
		}

		return nil, fmt.Errorf("invalid module")
	}
}

// okay, testing a private function is silly, but this just makes things easier
func TestSetBody(t *testing.T) {
	fd := &starlark.Dict{}
	fd.Set(starlark.String("foo"), starlark.String("bar"))

	cases := []struct {
		rawBody  starlark.String
		formData *starlark.Dict
		jsonData starlark.Value
		body     string
		err      string
	}{
		{starlark.String("hallo"), nil, nil, "hallo", ""},
		// TODO - this should check request form data is being set
		{starlark.String(""), fd, nil, "", ""},
		{starlark.String(""), nil, &starlark.Tuple{starlark.Bool(true), starlark.MakeInt(1), starlark.String("der")}, "[true,1,\"der\"]", ""},
	}

	for i, c := range cases {
		req := httptest.NewRequest("get", "https://example.com", nil)
		err := setBody(req, c.rawBody, c.formData, c.jsonData)
		if !(err == nil && c.err == "" || (err != nil && err.Error() == c.err)) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
			continue
		}
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != c.body {
			t.Errorf("case %d body mismatch. expected: %s, got: %s", i, c.body, string(body))
		}
	}
}
