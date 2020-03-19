package uniqueid

import (
	"bytes"
	"encoding/json"
	"strconv"
	"testing"
)

func TestTooLong(t *testing.T) {
	g := NewGenerator()
	o := NewOptionConfig()
	conf := SimpleIDConfig{
		Suff: "1234567",
	}
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	decoder := json.NewDecoder(buf)
	err := encoder.Encode(conf)
	if err != nil {
		panic(err)
	}
	o.Config = decoder.Decode
	o.Driver = "simpleid"
	err = o.ApplyTo(g)
	if err != ErrSuffTooLong {
		t.Fatal(err)
	}
}
func newSimpleIDGenerator() *Generator {
	g := NewGenerator()
	o := NewOptionConfig()
	conf := SimpleIDConfig{
		Suff: "-test",
	}
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	decoder := json.NewDecoder(buf)
	err := encoder.Encode(conf)
	if err != nil {
		panic(err)
	}
	o.Config = decoder.Decode
	o.Driver = "simpleid"
	err = o.ApplyTo(g)
	if err != nil {
		panic(err)
	}
	return g
}

func TestEncode(t *testing.T) {
	v1, err := strconv.ParseInt("21", 32, 64)
	if err != nil {
		panic(err)
	}
	v2, err := strconv.ParseInt("101", 32, 64)
	if err != nil {
		panic(err)
	}
	if !(encodeu32(uint32(v1)) < encodeu32(uint32(v2))) {
		t.Fatal(v1, v2)
	}
}

func TestSimpleID(t *testing.T) {
	generator := newSimpleIDGenerator()
	var last = ""
	var usedmap = map[string]bool{}
	for i := 0; i < 1000; i++ {
		id, err := generator.GenerateID()
		if err != nil {
			t.Fatal(err)
		}
		if usedmap[id] {
			t.Fatal(id)
		}
		usedmap[id] = true
		if last == id {
			t.Fatal(id)
		}
		if last >= id {
			t.Fatal(id)
		}
		last = id
	}
}

func BenchmarkSimpleID(b *testing.B) {
	generator := newSimpleIDGenerator()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generator.GenerateID()
		}
	})
}
