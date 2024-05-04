package mailer

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/ftsog/ecom/config"
	gomail "gopkg.in/mail.v2"
)

var (
	MailConfig *config.Config
)

type Mail struct {
	smtpHost   string
	port       string
	senderMail string
	password   string
}

type msgData struct {
	Username string
	Token    string
}

var ForgetMessage = `
	
	<h2 style="text-align: center"> Ecom </h2><br> 
	
	
	<h2 style="text-align:center">  </h2>
	<p style="text-align:center"> Hey %s, your are about to reset your password,
	click the link below to reset your password </p><br>
	
	<p style="align:center; height:30px; width:200px; background-color:blue; border-radius: 5px; padding:10px;"><a href="http://127.0.0.1:8080/api/auth/password/reset/for/%s/%s" style="text-decoration:none; color:white; text-align:center;"> Confirm email </a></p>
	<br>
	<p style="text-align:center;">if you did not request for this kindly ignore this email </p>
	
`

var VerifyMessage = `

	<h2 style="text-align: center"> Ecom </h2><br>
	
	<h2 style="text-align:center"> Verify Your Email </h2>
	
	<p> Hey %s click the link below to finish verifying your email address </p><br>
	
	<p style="text-align:center; height:30px; width:200px; background-color:blue; border-radius: 5px; padding:10px;"><a href="http://127.0.0.1:8080/api/auth/verify_email/%s/%s" style="text-decoration:none; color:white; text-align:center;"> Confirm email </a></p>
	<br>
	<p style="text-align:center;">didn't create an account? <a href="">click here</a> to remove this email address </p>
	
`

func NewMessage(receiver, subject, message, username, token string) *gomail.Message {

	mail := gomail.NewMessage()
	mail.SetHeader("From", MailConfig.Mailer.Email)
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", subject)

	msg := fmt.Sprintf(message, username, username, token)

	mail.SetBody("text/html", msg)

	return mail

}

func NewMail(receiver, subject, message, username, token string) error {

	mail := NewMessage(receiver, subject, message, username, token)

	port, err := strconv.Atoi(MailConfig.Mailer.Port)
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(MailConfig.Mailer.Host, port, MailConfig.Mailer.Email, MailConfig.Mailer.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = dialer.DialAndSend(mail)
	if err != nil {
		return err
	}

	return nil
}

/*
func SendVerificationEmail(receiver string, username string, token string) error {

	msgD := msgData{
		Username: username,
		Token:    token,
	}

	mail := NewMailStruct()
	m := gomail.NewMessage()
	m.SetHeader("From", mail.senderMail)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "Verify Your Email")

	message := `

	<h2 style="text-align: center"> Ecom </h2><br>

	<p style="text-align:center;"><img src="http://127.0.0.1:7000/images/IMG_20220318_111126_790.jpg"></p>

	<h2 style="text-align:center"> Verify Your Email </h2>
	<p style="text-align:center"> Hey {{.Username}}, thanks for helping us keep your account secure!
	click the button below to finish verifying your email address </p><br>

	<p style="text-align:center; height:30px; width:200px; background-color:blue; border-radius: 5px; padding:10px;"><a href="http://127.0.0.1:8080/api/auth/verify_email/{{.Username}}/{{.Token}}" style="text-decoration:none; color:white; text-align:center;"> Confirm email </a></p>
	<br>
	<p style="text-align:center;">didn't create an account? <a href="">click here</a> to remove this email address </p>

`

	f, err := os.OpenFile("message.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	tmpl := template.New("tmpl")
	tmp, err := tmpl.Parse(message)
	if err != nil {
		return err
	}

	err = tmp.Execute(f, msgD)
	if err != nil {
		return err
	}

	msg, err := ioutil.ReadFile("./message.txt")
	if err != nil {
		return err
	}

	m.SetBody("text/html", string(msg))

	port, _ := strconv.Atoi(mail.port)
	d := gomail.NewDialer(mail.smtpHost, port, mail.senderMail, mail.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = d.DialAndSend(m)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
*/
