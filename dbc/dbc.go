package dbc

import (
	"database/sql"

	t "time"
	//. "types"

	. "github.com/ratsil/go-helpers"

	_ "github.com/lib/pq"
)

var Cache *DBController

type IDriver interface {
	Open() (*sql.DB, error)
	Close(*sql.DB)
}
type Driver struct {
	Host     *string
	Port     *string
	Name     *string
	User     *string
	Password *string
}

func (this *Driver) Open() (*sql.DB, error) {
	return sql.Open("postgres", "host="+*this.Host+" port="+*this.Port+" dbname="+*this.Name+" user="+*this.User+" password="+*this.Password+" connect_timeout=10 sslmode=disable")
}
func (this *Driver) Close(pDB *sql.DB) {
	if nil != pDB {
		pDB.Close()
	}
}

type DBController struct {
	Driver IDriver
	pDB    *sql.DB
	IsOpen bool
}

func (this *DBController) Close() {
	this.Driver.Close(this.pDB)
	this.IsOpen = false
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
func (this *DBController) exec(sSQL string, args ...interface{}) (oRetVal sql.Result, err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	defer this.Close()

	oRetVal, err = this.pDB.Exec(sSQL, args...)
	return
}
func (this *DBController) QueryForID(sSQL string, args ...interface{}) (nRetVal ID, err error) {
	nSQLRes, err := this.exec(sSQL, args...)
	if err == nil {
		var nRes int64
		nRes, err = nSQLRes.LastInsertId()
		nRetVal = ID(nRes)
	}
	return
}
func (this *DBController) Query(sSQL string, args ...interface{}) (aRetVal []map[string]string, err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	defer this.Close()

	oRows, err := this.pDB.Query(sSQL, args...)
	if nil != err {
		return
	}

	aColumns, err := oRows.Columns()
	if nil != err {
		return
	}

	aValues := make([]sql.RawBytes, len(aColumns))

	aScanArgs := make([]interface{}, len(aValues))
	for i := range aValues {
		aScanArgs[i] = &aValues[i]
	}
	aRetVal = make([]map[string]string, 0, 0)
	for oRows.Next() {
		if err = oRows.Scan(aScanArgs...); nil != err {
			return
		}
		mValues := map[string]string{}
		for i, col := range aValues {
			mValues[aColumns[i]] = string(col)
		}

		aRetVal = append(aRetVal, mValues)
	}
	err = oRows.Err()
	return
}
func (this *DBController) Perform(sSQL string, args ...interface{}) (err error) {
	if !this.IsOpen {
		if err = this.Open(); nil != err {
			return
		}
	}
	defer this.Close()
	_, err = this.exec(sSQL, args...)
	return
}

func getID(o IRecord) *ID {
	defer func() { recover() }()
	if nil == o || IDNull == o.IdGet() {
		return nil
	}
	n := o.IdGet()
	return &n
}
func getDT(o ITimed) *t.Time {
	defer func() { recover() }()
	if nil == o {
		return nil
	}
	return parseDT(o.DtGet())
}
func parseDT(dt t.Time) *t.Time {
	if DTNull == dt {
		return nil
	}
	return &dt
}
