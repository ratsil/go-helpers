package types

import (
	"math"
	t "time"
)

const IDNull = ID(math.MaxInt64)

type ID int64

type IRecord interface {
	IdGet() ID
}
type Record struct {
	Id ID `json:"id"`
}

func (this *Record) IdGet() ID {
	return this.Id
}

type Named struct {
	Name string `json:"sName"`
}
type ITimed interface {
	DtGet() t.Time
}
type Timed struct {
	Dt t.Time `json:"dt"`
}

func (this *Timed) DtGet() t.Time {
	return this.Dt
}

type Dictionary struct {
	Record
	Named
}
