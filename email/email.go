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

type DKIM struct {
	Public   string `json:"public"`
	Private  string `json:"private"`
	Domain   string `json:"domain"`
	Selector string `json:"selector"`
}
type IMailer interface {
	SourceGet() string
	SourceSet(string)
	DKIMGet() *DKIM
	Send([]string, []byte) error
}
type Mailer struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DKIM     *DKIM  `json:"dkim,omitempty"`
}

func (this *Mailer) Send(aRecipients []string, aBytes []byte) error {
	return smtp.SendMail(this.Host+":"+this.Port,
		smtp.PlainAuth("",
			this.User,
			this.Password,
			this.Host,
		),
		this.User,
		aRecipients,
		aBytes)
}
func (this *Mailer) SourceGet() string {
	return this.User
}
func (this *Mailer) SourceSet(sSource string) {
	this.User = sSource
}
func (this *Mailer) DKIMGet() *DKIM {
	return this.DKIM
}

type SmtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

type SMTPController struct {
	Mailer IMailer
}

func (this *SMTPController) Send(aRecipients []string, sSubject string, sBody string, sBcc string, oUnsubscribe url.URL) (err error) {
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
	pSmtpTemplateData := &SmtpTemplateData{
		this.Mailer.SourceGet(),
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

	if pDKIM := this.Mailer.DKIMGet(); nil != pDKIM {
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
	return this.Mailer.Send(aRecipients, aBytes)
}
func (this *SMTPController) SendTemplate(aRecipients []string, sSubject, sBody string, sBcc string, oUnsubscribe url.URL, pTemplateData interface{}) (err error) {
	var oBuffer bytes.Buffer
	var pEmail *template.Template
	var aTemplateData []interface{}
	switch pTemplateData.(type) {
	case []interface{}:
		aTemplateData = (pTemplateData).([]interface{})
	default:
		aTemplateData = append(aTemplateData, pTemplateData)
	}
	if strings.Contains(sSubject, "{{") {
		for _, pTemplateData = range aTemplateData {
			pEmail, err = template.New("email").Parse(sSubject)
			if nil != err {
				return
			}
			if err = pEmail.Execute(&oBuffer, pTemplateData); nil != err {
				return
			}
			sSubject = oBuffer.String()
			oBuffer.Reset()
		}
	}
	if strings.Contains(sBody, "{{") {
		for _, pTemplateData = range aTemplateData {
			pEmail, err = template.New("email").Parse(sBody)
			if nil != err {
				return
			}
			if err = pEmail.Execute(&oBuffer, pTemplateData); nil != err {
				return
			}
			sBody = oBuffer.String()
			oBuffer.Reset()
		}
	}
	return this.Send(aRecipients, sSubject, sBody, sBcc, oUnsubscribe)
}
