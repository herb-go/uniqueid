package uniqueid

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

func encodeu32(data uint32) string {
	return encode64(int64(data))
}
func encode64(data int64) string {
	hexstr := strconv.FormatInt(data, 32)
	return strconv.FormatInt(int64(len(hexstr)), 32) + hexstr
}

//SimpleID simple id driver
type SimpleID struct {
	Current *uint32
	Suff    string
}

//GenerateID generate unique id.
//Return  generated id and any error if rasied.
func (i *SimpleID) GenerateID() (string, error) {
	ts := encode64(time.Now().UnixNano())
	current := encodeu32(atomic.AddUint32(i.Current, 1))
	return ts + current + i.Suff, nil
}

// NewSimpleID create new simpleid driver
func NewSimpleID() *SimpleID {
	time.Sleep(time.Millisecond)
	var c = rand.Uint32()
	return &SimpleID{
		Current: &c,
		Suff:    "",
	}
}

type SimpleIDConfig struct {
	Suff string
}

//SimpleIDFactory simple id driver factory
func SimpleIDFactory(loader func(v interface{}) error) (Driver, error) {
	i := NewSimpleID()
	conf := SimpleIDConfig{}
	err := loader(&conf)
	if err != nil {
		return nil, err
	}
	if len(conf.Suff) > 6 {
		return nil, ErrSuffTooLong
	}
	i.Suff = conf.Suff
	return i, nil
}

func init() {
	Register("simpleid", SimpleIDFactory)
}
