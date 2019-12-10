package snowflake

import (
	"bytes"
	"encoding/json"

	"github.com/herb-go/uniqueid"

	"testing"
)

func newSnowFlakeGenerator() *uniqueid.Generator {
	g := uniqueid.NewGenerator()
	o := uniqueid.NewOptionConfig()
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	decoder := json.NewDecoder(buf)
	err := encoder.Encode(SnowFlakeConfig{})
	if err != nil {
		panic(err)
	}
	o.Driver = "snowflake"
	o.Config = decoder.Decode
	err = o.ApplyTo(g)
	if err != nil {
		panic(err)
	}
	return g
}

func TestSnowFlake(t *testing.T) {
	generator := newSnowFlakeGenerator()
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

func BenchmarkSnowFlake(b *testing.B) {
	generator := newSnowFlakeGenerator()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			generator.GenerateID()
		}
	})
}
