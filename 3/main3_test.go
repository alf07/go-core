package main

import (
	"reflect"
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	if !m.Exists("one") {
		t.Fatal("one should exist")
	}
	if value, ok := m.Get("two"); !ok || value != 2 {
		t.Fatalf("got (%d, %v), want (2, true)", value, ok)
	}

	m.Remove("one")
	if m.Exists("one") {
		t.Fatal("one should not exist after remove")
	}
}

func TestStringIntMapGetMissing(t *testing.T) {
	m := NewStringIntMap()
	if value, ok := m.Get("missing"); ok || value != 0 {
		t.Fatalf("got (%d, %v), want (0, false)", value, ok)
	}
}

func TestStringIntMapCopyIsIndependent(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)

	copied := m.Copy()
	copied["one"] = 99
	copied["two"] = 2

	want := map[string]int{"one": 1}
	if !reflect.DeepEqual(m.Copy(), want) {
		t.Fatalf("original map changed: %v", m.Copy())
	}
}
