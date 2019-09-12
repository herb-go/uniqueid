package uniqueid

import "testing"
import "strings"

func TestUniqueID(t *testing.T) {
	UnregisterAll()
	defer func() {
		UnregisterAll()
		Register("simpleid", SimpleIDFactory)
	}()
	f := Factories()
	if len(f) != 0 {
		t.Fatal(f)
	}
	Register("simpleid", SimpleIDFactory)
	f = Factories()
	if len(f) != 1 {
		t.Fatal(f)
	}
}

func TestNotexistedDriver(t *testing.T) {
	d, err := NewDriver("notexist", nil, "")
	if d != nil {
		t.Fatal(d)
	}
	if err == nil || !strings.Contains(err.Error(), "unknown driver") {
		t.Fatal(err)
	}
}

func TestRegisterExistedDriver(t *testing.T) {
	UnregisterAll()
	defer func() {
		UnregisterAll()
		Register("simpleid", SimpleIDFactory)
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
		err := r.(error)
		if err == nil || !strings.Contains(err.Error(), "twice") {
			t.Fatal(err)
		}
	}()
	Register("simpleid", SimpleIDFactory)
	Register("simpleid", SimpleIDFactory)
}

func TestRegisterNilDriver(t *testing.T) {
	UnregisterAll()
	defer func() {
		UnregisterAll()
		Register("simpleid", SimpleIDFactory)
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
		err := r.(error)
		if err == nil || !strings.Contains(err.Error(), "nil") {
			t.Fatal(err)
		}
	}()
	Register("nil", nil)
}

func TestDefaultGenerator(t *testing.T) {
	defer func() {
		DefaultGenerator = nil
	}()
	g := newSimpleIDGenerator()
	DefaultGenerator = g
	id1, err := GenerateID()
	if err != nil {
		t.Fatal(err)
	}
	id2 := MustGenerateID()
	if id1 == id2 {
		t.Fatal(id1, id2)
	}
}
