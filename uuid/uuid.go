package uuid

import (
	"github.com/herb-go/uniqueid"
	uuid "github.com/satori/go.uuid"
)

//UUID uuid driver
type UUID struct {
	creator func() (uuid.UUID, error)
}

//NewUUID create new uuid driver
func NewUUID() *UUID {
	return &UUID{}
}

//V1 generate unique id by uuid version1.
//Return  generated id and any error if rasied.
func V1() (string, error) {
	uid, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

//GenerateID generate unique id.
//Return  generated id and any error if rasied.
func (u *UUID) GenerateID() (string, error) {
	uid, err := u.creator()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

//Factory uuid driver factory
func Factory(conf uniqueid.Config, prefix string) (uniqueid.Driver, error) {
	i := NewUUID()
	var version int
	conf.Get(prefix+"Version", &version)
	switch version {
	case 4:
		i.creator = uuid.NewV4
	default:
		i.creator = uuid.NewV1
	}
	return i, nil

}

func init() {
	uniqueid.Register("uuid", Factory)
}