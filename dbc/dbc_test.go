package dbc

import (
	"database/sql"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock.v1"
	"github.com/stretchr/testify/assert"
)

var _nTestsQty int

type Tests struct {
	Assert *assert.Assertions
	DBMock sqlmock.Sqlmock
	DBC    *DBController
}
type TestDriver struct {
	t  *testing.T
	DB *sql.DB
}

func Setup(t *testing.T) (pTests *Tests) {
	pTests = &Tests{Assert: assert.New(t)}

	pDB, oDBMock, err := sqlmock.New()
	pTests.Assert.Nil(err)
	pTests.Assert.NotNil(pDB)
	pTests.Assert.NotNil(oDBMock)
	pTests.DBMock = oDBMock
	pTests.DBC = &DBController{
		Driver: &TestDriver{t: t, DB: pDB},
	}
	return
}

//---------------------------------------------------------------------------
func (th *TestDriver) Open() (pRetVal *sql.DB, err error) {
	return th.DB, nil
}
func (th *TestDriver) Close(pDB *sql.DB) {}

func Next() {
	_nTestsQty++
	fmt.Printf("Test #%d\n", _nTestsQty)
}
