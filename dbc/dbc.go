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

//IDriver .
type IDriver interface {
	Open() (*sql.DB, error)
	Close(*sql.DB)
}

//Driver .
type Driver struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

//Open .
func (th *Driver) Open() (*sql.DB, error) {
	return sql.Open("postgres", "host="+th.Host+" port="+th.Port+" dbname="+th.Name+" user="+th.User+" password="+th.Password+" connect_timeout=10 sslmode=disable")
}

//Close .
func (th *Driver) Close(pDB *sql.DB) {
	if nil != pDB {
		pDB.Close()
	}
}

//Cache .
var Cache *DBController

//DBController .
type DBController struct {
	Driver       IDriver
	pDB          *sql.DB
	pTransaction *sql.Tx
	IsOpen       bool
}

//Open .
func (th *DBController) Open() (err error) {
	th.IsOpen = false
	th.Driver.Close(th.pDB)
	if th.pDB, err = th.Driver.Open(); nil != err {
		return
	}
	th.IsOpen = true
	Cache = th
	return
}

//Close .
func (th *DBController) Close() {
	th.Driver.Close(th.pDB)
	th.IsOpen = false
	th.pTransaction = nil
}

//Begin .
func (th *DBController) Begin() (err error) {
	if !th.IsOpen {
		if err = th.Open(); nil != err {
			return
		}
	}
	th.pTransaction, err = th.pDB.Begin()
	return
}

//Commit .
func (th *DBController) Commit() error {
	if nil == th.pTransaction {
		return errors.New("no active transaction")
	}
	p := th.pTransaction
	th.pTransaction = nil
	return p.Commit()
}

//Rollback .
func (th *DBController) Rollback() error {
	if nil == th.pTransaction {
		return errors.New("no active transaction")
	}
	p := th.pTransaction
	th.pTransaction = nil
	return p.Rollback()
}

//IsTransacted .
func (th *DBController) IsTransacted() bool {
	return nil != th.pTransaction
}

//QueryForID .
func (th *DBController) QueryForID(sSQL string, args ...interface{}) (nRetVal ID, err error) {
	aDBValues, err := th.Query(sSQL, args...)
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

//QueryForValue .
func (th *DBController) QueryForValue(sSQL, sName string, args ...interface{}) (sRetVal string, err error) {
	aDBValues, err := th.Query(sSQL, args...)
	if nil != err {
		return
	}
	if 0 < len(aDBValues) {
		sRetVal = aDBValues[0][sName]
	}
	return
}

//QueryForRow .
func (th *DBController) QueryForRow(sSQL string, args ...interface{}) (mRetVal map[string]string, err error) {
	aDBValues, err := th.Query(sSQL, args...)
	if nil != err {
		return
	}
	if 0 < len(aDBValues) {
		mRetVal = aDBValues[0]
	}
	return
}

//Query .
func (th *DBController) Query(sSQL string, args ...interface{}) (aRetVal []map[string]string, err error) {
	if !th.IsOpen {
		if err = th.Open(); nil != err {
			return
		}
	}
	var pRows *sql.Rows
	if nil == th.pTransaction {
		pRows, err = th.pDB.Query(sSQL, args...)
	} else {
		pRows, err = th.pTransaction.Query(sSQL, args...)
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

//Perform .
func (th *DBController) Perform(sSQL string, args ...interface{}) (err error) {
	if !th.IsOpen {
		if err = th.Open(); nil != err {
			return
		}
	}
	_, err = th.exec(sSQL, args...)
	return
}

func (th *DBController) exec(sSQL string, args ...interface{}) (oRetVal sql.Result, err error) {
	if !th.IsOpen {
		if err = th.Open(); nil != err {
			return
		}
	}
	if nil == th.pTransaction {
		oRetVal, err = th.pDB.Exec(sSQL, args...)
	} else {
		oRetVal, err = th.pTransaction.Exec(sSQL, args...)
	}
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
