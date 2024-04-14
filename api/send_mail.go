package api

import "net/smtp"

func SendEmail(to, subject, body string) error {
	from := "matrasser@mail.ru"
	password := "7732178Mm"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	return err
}
