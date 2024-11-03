package main

import (
	"regexp"
	"strings"

	"github.com/go-mail/mail"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address."
	}

	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please enter a message"
	}

	return len(msg.Errors) == 0
}

func (msg *Message) DeliverAsync() error {
	username := goDotEnvVariable("MAILTRAP_UN")
	password := goDotEnvVariable("MAILTRAP_PW")

	email := mail.NewMessage()
	email.SetHeader("To", "carlos@csaenz.dev")
	email.SetHeader("From", "server@csaenz.dev")
	email.SetHeader("Reply-To", msg.Email)
	email.SetHeader("Subject", "New message via Contact Form")
	email.SetBody("text/plain", msg.Content)
	email.AddAlternative("text/html",
		`
        <html>
            <body>
                <h1>New Message via Contact Form</h1>
                <p>"
		`+msg.Content+`
		"</p>
                <p>Made with &#128140;<br>Mailtrap</p>
            </body>
        </html>
		`)

	return mail.NewDialer("live.smtp.mailtrap.io", 587, username, password).DialAndSend(email)
}
