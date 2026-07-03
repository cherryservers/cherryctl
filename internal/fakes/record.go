package fakes

import (
	"reflect"
	"testing"
)

type CallRecord struct {
	params []any
	method string
}

func (cr CallRecord) AssertParams(t *testing.T, want ...any) {
	t.Helper()
	got := cr.params

	if len(got) != len(want) {
		t.Errorf("want %d params for %q, got %d", len(want), cr.method, len(got))
	}

	l := min(len(want), len(got))
	for i := range l {
		if reflect.TypeOf(want[i]) != reflect.TypeOf(got[i]) {
			t.Errorf("want %T type for param number %d for %q, got %T", want[i], i, cr.method, got[i])
		} else if !reflect.DeepEqual(want[i], got[i]) {
			t.Errorf("want %#v for param number %d for %q, got %#v", want[i], i, cr.method, got[i])
		}
	}
}

func (cr CallRecord) AssertMethod(t *testing.T, want string) {
	t.Helper()
	if want != cr.method {
		t.Errorf("want method %q, got %q", want, cr.method)
	}
}

type OutputRecord struct {
	in any
	th []string
	td [][]string
}

func (or OutputRecord) Assert(t *testing.T, wantIn any, wantTh []string, wantTd [][]string) {
	t.Helper()

	if reflect.TypeOf(wantIn) != reflect.TypeOf(or.in) {
		t.Errorf("want %T type for output val, got %T", wantIn, or.in)
	} else if !reflect.DeepEqual(wantIn, or.in) {
		t.Errorf("want %#v output val, got %#v", wantIn, or.in)
	}

	if !reflect.DeepEqual(wantTh, or.th) {
		t.Errorf("want %#v output table header, got %#v", wantTh, or.th)
	}

	if !reflect.DeepEqual(wantTd, or.td) {
		t.Errorf("want %#v output table data, got %#v", wantTd, or.td)
	}
}
