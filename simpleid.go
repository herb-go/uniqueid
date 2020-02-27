package uniqueid

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"sync/atomic"
	"time"
)

//SimpleID simple id driver
type SimpleID struct {
	Current *uint32
	Suff    string
}

//GenerateID generate unique id.
//Return  generated id and any error if rasied.
func (i *SimpleID) GenerateID() (string, error) {
	buf1 := bytes.NewBuffer(nil)
	buf2 := bytes.NewBuffer(nil)
	ts := time.Now().UnixNano()
	err := binary.Write(buf1, binary.BigEndian, ts)
	if err != nil {
		return "", err
	}
	c := atomic.AddUint32(i.Current, 1)
	err = binary.Write(buf2, binary.BigEndian, c)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf1.Bytes()) + "-" + hex.EncodeToString(buf2.Bytes()) + i.Suff, nil
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
	i.Suff = conf.Suff
	return i, nil
}

func init() {
	Register("simpleid", SimpleIDFactory)
}
