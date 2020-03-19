package uniqueid

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"sync/atomic"
	"time"
)

func encodeTimestamp(ts int64) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.BigEndian, ts)
	if err != nil {
		return "", err
	}

	hexstr := hex.EncodeToString(buf.Bytes())
	lengthbuf := bytes.NewBuffer(nil)
	err = binary.Write(buf, binary.BigEndian, len(hexstr))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(lengthbuf.Bytes()) + hexstr, nil
}

//SimpleID simple id driver
type SimpleID struct {
	Current *uint32
	Suff    string
}

//GenerateID generate unique id.
//Return  generated id and any error if rasied.
func (i *SimpleID) GenerateID() (string, error) {
	buf2 := bytes.NewBuffer(nil)
	ts, err := encodeTimestamp(time.Now().UnixNano())
	if err != nil {
		return "", err
	}
	c := atomic.AddUint32(i.Current, 1)
	err = binary.Write(buf2, binary.BigEndian, c)
	if err != nil {
		return "", err
	}
	return ts + hex.EncodeToString(buf2.Bytes()) + i.Suff, nil
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
	if len(conf.Suff) > 7 {
		return nil, ErrSuffTooLong
	}
	i.Suff = conf.Suff
	return i, nil
}

func init() {
	Register("simpleid", SimpleIDFactory)
}
