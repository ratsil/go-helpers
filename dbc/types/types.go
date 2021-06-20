package types

import (
	"encoding/json"
	"fmt"
	"math"
	sc "strconv"
	s "strings"
	t "time"
)

//IDNull .
const IDNull = ID(math.MaxInt64)

//ID .
type ID int64

//MarshalJSON .
func (th ID) MarshalJSON() (aRetVal []byte, err error) {
	aRetVal, err = json.Marshal(int64(th))
	return
}

//MarshalText implements TextMarshaler
func (th ID) MarshalText() (text []byte, err error) {
	return []byte(sc.FormatInt(int64(th), 10)), nil
}

//IsEmpty .
func (th ID) IsEmpty() bool {
	return 1 > th || IDNull == th
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

//TimedNullable .
type TimedNullable struct {
	Dt *t.Time `json:"dt"`
}

//DtGet .
func (th *TimedNullable) DtGet() *t.Time {
	return th.Dt
}

//Dictionary .
type Dictionary struct {
	Record
	Named
}

//Point .
type Point struct {
	X float64
	Y float64
}

//MarshalJSON .
func (th *Point) MarshalJSON() ([]byte, error) {
	if nil == th {
		return []byte("null"), nil
	}
	return []byte(`"(` + fmt.Sprintf("%f", th.X) + "," + fmt.Sprintf("%f", th.Y) + `)"`), nil
}

//UnmarshalJSON .
func (th *Point) UnmarshalJSON(b []byte) (err error) {
	a := s.Split(s.Trim(string(b), `()"`), ",")
	if th.X, err = sc.ParseFloat(a[0], 64); err == nil {
		th.Y, err = sc.ParseFloat(a[1], 64)
	}
	return
}
