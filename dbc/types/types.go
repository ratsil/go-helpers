package types

import (
	"math"
	sc "strconv"
	t "time"
)

//IDNull .
const IDNull = ID(math.MaxInt64)

//ID .
type ID int64

//MarshalText implements TextMarshaler
func (th ID) MarshalText() (text []byte, err error) {
	return []byte(sc.FormatInt(int64(th), 10)), nil
}

//IRecord .
type IRecord interface {
	IdGet() ID
}

//Record .
type Record struct {
	Id ID `json:"id"`
}

//IdGet .
func (th *Record) IdGet() ID {
	return th.Id
}

//Named .
type Named struct {
	Name string `json:"sName"`
}

//ITimed .
type ITimed interface {
	DtGet() *t.Time
}

//Timed .
type Timed struct {
	Dt t.Time `json:"dt"`
}

//DtGet .
func (th *Timed) DtGet() *t.Time {
	return &th.Dt
}

//Dictionary .
type Dictionary struct {
	Record
	Named
}
