package email

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestMailer struct {
	Raw  []byte
	DKIM *DKIM
}

func (this *TestMailer) Send(aRecipients []string, aBytes []byte) error {
	this.Raw = aBytes
	//	this.t.Logf("%s", aBytes)
	return nil
}
func (this *TestMailer) SourceGet() string {
	return "from@unit.test"
}
func (this *TestMailer) SourceSet(sSource string) {
}
func (this *TestMailer) DKIMGet() *DKIM {
	return this.DKIM
}

func TestSend(t *testing.T) {
	pAssert := assert.New(t)
	pMailer := &TestMailer{}
	oSMTP := SMTPController{
		Mailer: pMailer,
	}
	pAssert.Nil(oSMTP.Send([]string{"to@unit.test"}, "[bspc][alerts]", "test", "", url.URL{}))
	pAssert.Equal("From: from@unit.test\r\nReply-To: from@unit.test\r\nTo: to@unit.test\r\nMIME-Version: 1.0\r\nContent-type: text/html;charset=utf-8\r\nSubject: =?utf-8?B?W2JzcGNdW2FsZXJ0c10=?=\r\n\r\n<html><body>test</body></html>\r\n",
		string(pMailer.Raw))
}
