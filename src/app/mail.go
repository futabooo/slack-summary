package app

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

// CreateMailHeader ...
func CreateMailHeader(from, to, title string) map[string]string {
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	return header
}

// SendMail ...
func SendMail(conf ConfToml, mailBody string) {
	const host string = "smtp.gmail.com"
	const addr string = "smtp.gmail.com:587"

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		conf.GoogleAccount.Name,
		conf.GoogleAccount.Pass,
		host,
	)

	from := mail.Address{conf.GoogleAccount.Name, conf.GoogleAccount.Name}
	to := mail.Address{conf.GoogleAccount.Name, conf.GoogleAccount.Name}
	title := Name

	message := ""
	for k, v := range CreateMailHeader(from.String(), to.String(), title) {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(mailBody))

	fmt.Printf(message)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		addr,
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
}
