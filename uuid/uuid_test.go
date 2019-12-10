package uuid

import (
	"bytes"
	"encoding/json"

	"github.com/herb-go/uniqueid"

	"testing"
)

func newUUIDGenerator() *uniqueid.Generator {
	g := uniqueid.NewGenerator()
	o := uniqueid.NewOptionConfig()
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	decoder := json.NewDecoder(buf)
	err := encoder.Encode(UUIDConfig{})
	if err != nil {
		panic(err)
	}

	o.Driver = "uuid"
	o.Config = decoder.Decode
	err = o.ApplyTo(g)
	if err != nil {
		panic(err)
	}
	return g
}

func newUUIDGeneratorV4() *uniqueid.Generator {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	decoder := json.NewDecoder(buf)
	err := encoder.Encode(UUIDConfig{
		Version: 4,
	})
	if err != nil {
		panic(err)
	}

	g := uniqueid.NewGenerator()
	o := uniqueid.NewOptionConfig()
	o.Config = decoder.Decode
	o.Driver = "uuid"
	err = o.ApplyTo(g)
	if err != nil {
		panic(err)
	}
	return g
}

func TestUUID(t *testing.T) {
	generator := newUUIDGenerator()
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

		last = id
	}
}

func TestUUIDV4(t *testing.T) {
	generator := newUUIDGeneratorV4()
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
		last = id
	}
}

func BenchmarkUUID(b *testing.B) {
	generator := newUUIDGenerator()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generator.GenerateID()
		}
	})
}

func BenchmarkUUIDV4(b *testing.B) {
	generator := newUUIDGeneratorV4()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generator.GenerateID()
		}
	})
}
