package helpers

import (
	crypto "crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
	"os"
	"regexp"
	s "strings"
	t "time"
)

var TimeMax = t.Unix(1<<63-62135596801, 999999999)
var DTNull = TimeMax

//MD5 returns MD5 hash as a string
func MD5(s string) string {
	var aHash = crypto.Sum([]byte(s))
	return hex.EncodeToString(aHash[:])
}

//PasswordGenerator generates new password
func PasswordGenerator() (string, error) {
	aHash := make([]byte, 32)
	_, err := rand.Read(aHash)
	if nil != err {
		return "", err
	}
	aValue := make([]byte, base64.StdEncoding.EncodedLen(len(aHash)))
	base64.StdEncoding.Encode(aValue, aHash)
	aValue = regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAll(aValue, nil)
	return string(aValue[:8]), nil
}

//FileExists file's existance test
func FileExists(s string) (bool, error) {
	_, err := os.Stat(s)
	if nil == err {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//FileWriteAllBytes writes all bytes to a file
func FileWriteAllBytes(sFile string, aBytes []byte) error {
	pFile, err := os.OpenFile(sFile, os.O_CREATE|os.O_WRONLY, 0666)
	if nil != err {
		return err
	}
	defer pFile.Close()
	_, err = pFile.Write(aBytes) //UNDONE maybe we need to check returned bytes qty
	return err
}

//FileReadAllBytes Reads all bytes from a file
func FileReadAllBytes(sFile string) (aBytes []byte, err error) {
	pFile, err := os.OpenFile(sFile, os.O_CREATE|os.O_RDONLY, 0666)
	if nil != err {
		return
	}
	defer pFile.Close()
	_, err = pFile.Read(aBytes) //UNDONE maybe we need to check returned bytes qty
	return aBytes, err
}

//IsEmpty tests for nil/zero length/etc
func IsEmpty(oValue interface{}) bool {
	if nil == oValue {
		return true
	}
	if dt, b := oValue.(t.Time); b {
		return TimeMax == dt
	}
	if n, b := oValue.(uint64); b {
		return math.MaxUint64 == n
	}
	if s, b := oValue.(string); b {
		return 1 > len(s)
	}
	return false
}

//E shorthand for errors.New
func E(s string) error {
	return errors.New(s)
}

//ReplaceAll .
func ReplaceAll(sSource string, aSearch []string, oReplace interface{}) (sRetVal string) {
	sRetVal = sSource
	if aTarget, b := oReplace.([]string); b {
		for n, sSource := range aSearch {
			sRetVal = s.Replace(sRetVal, sSource, aTarget[n], -1)
		}
	} else {
		sSource = oReplace.(string)
		for _, sSource := range aSearch {
			sRetVal = s.Replace(sRetVal, sSource, sSource, -1)
		}
	}
	return
}
