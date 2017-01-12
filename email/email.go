package email

import (
	"bytes"
	"encoding/base64"
	"net/smtp"
	"net/url"
	"strings"
	"text/template"

	dkim "github.com/toorop/go-dkim"
	//	"log"
)

//DKIM .
type DKIM struct {
	Public   string `json:"public"`
	Private  string `json:"private"`
	Domain   string `json:"domain"`
	Selector string `json:"selector"`
}

//IMailer .
type IMailer interface {
	SourceGet() string
	SourceSet(string)
	DKIMGet() *DKIM
	Send([]string, []byte) error
}

//Mailer .
type Mailer struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DKIM     *DKIM  `json:"dkim,omitempty"`
}

//Send .
func (th *Mailer) Send(aRecipients []string, aBytes []byte) error {
	return smtp.SendMail(th.Host+":"+th.Port,
		smtp.PlainAuth("",
			th.User,
			th.Password,
			th.Host,
		),
		th.User,
		aRecipients,
		aBytes)
}

//SourceGet .
func (th *Mailer) SourceGet() string {
	return th.User
}

//SourceSet .
func (th *Mailer) SourceSet(sSource string) {
	th.User = sSource
}

//DKIMGet .
func (th *Mailer) DKIMGet() *DKIM {
	return th.DKIM
}

//SMTPTemplateData .
type SMTPTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

//SMTPController .
type SMTPController struct {
	Mailer IMailer
}

//Send .
func (th *SMTPController) Send(aRecipients []string, sSubject string, sBody string, sBcc string, oUnsubscribe url.URL) (err error) {
	sHeaders := "From: {{.From}}\r\nReply-To: {{.From}}\r\nTo: {{.To}}\r\nMIME-Version: 1.0\r\nContent-type: text/html;charset=utf-8\r\n"
	if 0 < len(sBcc) {
		sHeaders += "Bcc: " + sBcc + "\r\n"
	}
	if 0 < len(oUnsubscribe.RawQuery) {
		sHeaders += "List-Unsubscribe: <" + oUnsubscribe.String() + ">\r\n"
	}
	sSubjectBased := make([]byte, base64.StdEncoding.EncodedLen(len(sSubject)))
	base64.StdEncoding.Encode(sSubjectBased, []byte(sSubject))
	sSubject = string(sSubjectBased)
	pSmtpTemplateData := &SMTPTemplateData{
		th.Mailer.SourceGet(),
		strings.Join(aRecipients, ","),
		sSubject,
		sBody,
	}

	var pBuffer *bytes.Buffer
	pBuffer = new(bytes.Buffer)
	oEmail := template.New("email")
	if oEmail, err = oEmail.Parse(sHeaders); nil != err {
		return
	}
	if err = oEmail.Execute(pBuffer, pSmtpTemplateData); nil != err {
		return
	}
	sHeaders = pBuffer.String()

	pBuffer = new(bytes.Buffer)
	oEmail = template.New("email")
	if oEmail, err = oEmail.Parse("Subject: =?utf-8?B?{{.Subject}}?=\r\n\r\n"); nil != err {
		return
	}
	if err = oEmail.Execute(pBuffer, pSmtpTemplateData); nil != err {
		return
	}
	sSubject = pBuffer.String()

	pBuffer = new(bytes.Buffer)
	oEmail = template.New("email")
	if oEmail, err = oEmail.Parse("<html><body>{{.Body}}</body></html>\r\n"); nil != err {
		return
	}
	if err = oEmail.Execute(pBuffer, pSmtpTemplateData); nil != err {
		return
	}
	sBody = pBuffer.String()
	aBytes := []byte(sHeaders + sSubject + sBody)

	if pDKIM := th.Mailer.DKIMGet(); nil != pDKIM {
		oDKIMOptions := dkim.NewSigOptions()
		oDKIMOptions.PrivateKey = []byte(pDKIM.Private)
		oDKIMOptions.Domain = pDKIM.Domain
		oDKIMOptions.Selector = pDKIM.Selector
		oDKIMOptions.SignatureExpireIn = 3600
		oDKIMOptions.BodyLength = uint(len(sBody))
		oDKIMOptions.Headers = []string{"from", "date", "mime-version", "received", "received"}
		oDKIMOptions.AddSignatureTimestamp = true
		oDKIMOptions.Canonicalization = "relaxed/relaxed"
		if err = dkim.Sign(&aBytes, oDKIMOptions); nil != err {
			return
		}
	}
	return th.Mailer.Send(aRecipients, aBytes)
}

//SendTemplate .
func (th *SMTPController) SendTemplate(aRecipients []string, sSubject, sBody string, sBcc string, oUnsubscribe url.URL, mTemplateData map[string]interface{}) (err error) {
	var oBuffer bytes.Buffer
	var pEmail *template.Template
	if nil == mTemplateData {
		mTemplateData = make(map[string]interface{})
	}
	if strings.Contains(sSubject, "{{") {
		pEmail, err = template.New("email").Parse(sSubject)
		if nil != err {
			return
		}
		if err = pEmail.Execute(&oBuffer, mTemplateData); nil != err {
			return err
		}
		sSubject = oBuffer.String()
		oBuffer.Reset()
	}
	if strings.Contains(sBody, "{{") {
		pEmail, err = template.New("email").Parse(sBody)
		if nil != err {
			return
		}
		if err = pEmail.Execute(&oBuffer, mTemplateData); nil != err {
			return err
		}
		sBody = oBuffer.String()
		oBuffer.Reset()
	}
	return th.Send(aRecipients, sSubject, sBody, sBcc, oUnsubscribe)
}
