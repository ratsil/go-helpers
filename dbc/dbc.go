package dbc

import (
	"database/sql"
	"errors"
	"reflect"
	//t "time"

	. "github.com/ratsil/go-helpers/common"
	. "github.com/ratsil/go-helpers/dbc/types"

	_ "github.com/lib/pq"
)

var Cache *DBController

type IDriver interface {
	Open() (*sql.DB, error)
	Close(*sql.DB)
}
type Driver struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (this *Driver) Open() (*sql.DB, error) {
	return sql.Open("postgres", "host="+this.Host+" port="+this.Port+" dbname="+this.Name+" user="+this.User+" password="+this.Password+" connect_timeout=10 sslmode=disable")
}
func (this *Driver) Close(pDB *sql.DB) {
	if nil != pDB {
		pDB.Close()
	}
}

type DBController struct {
	Driver       IDriver
	pDB          *sql.DB
	pTransaction *sql.Tx
	IsOpen       bool
}

func (this *DBController) Close() {
	this.Driver.Close(this.pDB)
	this.IsOpen = false
	this.pTransaction = nil
}
func (this *DBController) Open() (err error) {
	this.IsOpen = false
	this.Driver.Close(this.pDB)
	if this.pDB, err = this.Driver.Open(); nil != err {
		return
	}
	this.IsOpen = true
	Cache = this
	return
}
func (this *DBController) Begin() (err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	this.pTransaction, err = this.pDB.Begin()
	return
}
func (this *DBController) Commit() error {
	if nil == this.pTransaction {
		return errors.New("no active transaction")
	}
	p := this.pTransaction
	this.pTransaction = nil
	return p.Commit()
}
func (this *DBController) Rollback() error {
	if nil == this.pTransaction {
		return errors.New("no active transaction")
	}
	p := this.pTransaction
	this.pTransaction = nil
	return p.Rollback()
}
func (this *DBController) IsTransacted() bool {
	return nil != this.pTransaction
}
func (this *DBController) exec(sSQL string, args ...interface{}) (oRetVal sql.Result, err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	if nil == this.pTransaction {
		oRetVal, err = this.pDB.Exec(sSQL, args...)
	} else {
		oRetVal, err = this.pTransaction.Exec(sSQL, args...)
	}
	return
}
func (this *DBController) QueryForID(sSQL string, args ...interface{}) (nRetVal ID, err error) {
	aDBValues, err := this.Query(sSQL, args...)
	if err != nil {
		return
	}
	for _, mValues := range aDBValues {
		for _, v := range mValues {
			nRetVal = ToID(v)
			return
		}
	}

	return
}
func (this *DBController) QueryForValue(sSQL, sName string, args ...interface{}) (sRetVal string, err error) {
	aDBValues, err := this.Query(sSQL, args...)
	if nil != err {
		return
	}
	if 0 < len(aDBValues) {
		sRetVal = aDBValues[0][sName]
	} else {
		err = errors.New("no value has been found")
	}
	return
}
func (this *DBController) QueryForRow(sSQL string, args ...interface{}) (mRetVal map[string]string, err error) {
	aDBValues, err := this.Query(sSQL, args...)
	if nil != err {
		return
	}
	if 0 < len(aDBValues) {
		mRetVal = aDBValues[0]
	} else {
		err = errors.New("no row has been found")
	}
	return
}
func (this *DBController) Query(sSQL string, args ...interface{}) (aRetVal []map[string]string, err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	var pRows *sql.Rows
	if nil == this.pTransaction {
		pRows, err = this.pDB.Query(sSQL, args...)
	} else {
		pRows, err = this.pTransaction.Query(sSQL, args...)
	}
	if nil != err {
		return
	}

	aColumns, err := pRows.Columns()
	if nil != err {
		return
	}

	aValues := make([]sql.RawBytes, len(aColumns))
	aScanArgs := make([]interface{}, len(aValues))
	for i := range aValues {
		aScanArgs[i] = &aValues[i]
	}
	aRetVal = make([]map[string]string, 0, 0)
	for pRows.Next() {
		if err = pRows.Scan(aScanArgs...); nil != err {
			return
		}
		mValues := map[string]string{}

		for i, col := range aValues {
			mValues[aColumns[i]] = string(col)
		}

		aRetVal = append(aRetVal, mValues)
	}
	err = pRows.Err()
	return
}
func (this *DBController) Perform(sSQL string, args ...interface{}) (err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	_, err = this.exec(sSQL, args...)
	return
}

func getID(o IRecord) *ID {
	defer func() { recover() }()
	if nil == o || reflect.ValueOf(o).IsNil() || IDNull == o.IdGet() {
		return nil
	}
	n := o.IdGet()
	return &n
}

/*
func getDT(o ITimed) *t.Time {
	defer func() { recover() }()
	if nil == o {
		return nil
	}
	return parseDT(o.DtGet())
}
func parseDT(dt t.Time) *t.Time {
	if t.Unix(1<<63-62135596801, 999999999) == dt {
		return nil
	}
	return &dt
}*/
