package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	t "time"
)

var (
	_pDefault  *Logger
	_oMessages chan *message
)

//Logger .
type Logger struct {
	nDay    int
	sPath   string
	sPrefix string
	sFile   string
	pFile   *os.File
}
type message struct {
	Logger *Logger
	Value  string
}

func init() {
	_oMessages = make(chan *message, 1048576)
	go writer()
}

//New .
func New(sPath, sPrefix string) *Logger {
	return &Logger{
		nDay:    t.Now().Day(),
		sPath:   sPath,
		sPrefix: sPrefix,
	}
}

//Default .
func Default(sPath, sPrefix string) {
	_pDefault = New(sPath, sPrefix)
}
func writer() {
	var aLoggers []*Logger
	fClose := func() {
		for _, o := range aLoggers {
			o.close()
		}
	}
	defer fClose()
	var oMessage *message
	for {
		select {
		case oMessage = <-_oMessages:
		default:
			fClose()
			aLoggers = []*Logger{}
			oMessage = <-_oMessages
		}
		if nil == oMessage.Logger.pFile {
			oMessage.Logger.open()
			aLoggers = append(aLoggers, oMessage.Logger)
		}
		oMessage.Logger.write(oMessage.Value)
	}
}

func (th *Logger) fileNameGet() string {
	if 1 > len(th.sFile) || th.nDay != t.Now().Day() {
		th.nDay = t.Now().Day()
		th.sFile = filepath.Join(th.sPath, th.sPrefix+"_"+t.Now().Format("060102")+".log")
	}
	return th.sFile
}

//Notice .
func (th *Logger) Notice(sMessage string) {
	_oMessages <- &message{
		Logger: th,
		Value:  sMessage,
	}
}

//Warning .
func (th *Logger) Warning(err error) {
	if nil != err {
		th.Notice("WARNING:" + err.Error())
	}
}

//Error .
func (th *Logger) Error(err error) {
	if nil != err {
		th.Notice("ERROR:" + err.Error())
	}
}

func (th *Logger) open() (err error) {
	th.close()
	th.pFile, err = os.OpenFile(th.fileNameGet(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	return
}
func (th *Logger) write(sMessage string) (err error) {
	if nil == th.pFile {
		if err = th.open(); nil != err {
			return
		}
	}
	_, err = th.pFile.WriteString(t.Now().Format("2006-01-02 15:04:05") + ": " + sMessage + "\n")
	return
}
func (th *Logger) close() {
	if nil != th.pFile {
		th.pFile.Close()
		th.pFile = nil
	}
}

//Notice .
func Notice(sMessage string) {
	if nil == _pDefault {
		panic("no default log")
	}
	_oMessages <- &message{
		Logger: _pDefault,
		Value:  sMessage,
	}
}

//Printf .
func Printf(sMessage string, aArgs ...interface{}) {
	Notice(fmt.Sprintf(sMessage, aArgs...))
}

//Print .
func Print(iMessage interface{}) {
	Notice(fmt.Sprintf("%v", iMessage))
}

//Debug .
func Debug(iMessage interface{}) {
	DebugWithArgs("%v", iMessage)
}

//DebugWithArgs .
func DebugWithArgs(sMessage string, aArgs ...interface{}) {
	var aBuf []byte
	runtime.Stack(aBuf, false)
	Notice("DEBUG:" + fmt.Sprintf(sMessage, aArgs...) + "\n" + string(debug.Stack()))
}

//Warning .
func Warning(err error) {
	if nil != err {
		Notice("WARNING:" + err.Error())
	}
}

//Error .
func Error(err error) {
	if nil != err {
		var aBuf []byte
		runtime.Stack(aBuf, false)
		Notice("ERROR:" + err.Error() + "\n" + string(aBuf))
	}
}

//ErrorWithStack .
func ErrorWithStack(err error) {
	if nil != err {
		Notice("ERROR:" + err.Error() + "\n" + string(debug.Stack()))
	}
}

//Recover .
func Recover() (iRetVal interface{}) {
	if iRetVal := recover(); nil != iRetVal {
		Notice("RECOVER:" + fmt.Sprintf("%v", iRetVal) + "\n" + string(debug.Stack()))
	}
	return
}

//Fatal .
func Fatal(err error) {
	if nil != _pDefault {
		if nil == _pDefault.pFile {
			_pDefault.open()
		}
		_pDefault.write(err.Error())
		_pDefault.close()
	}
	panic(err.Error())
}
