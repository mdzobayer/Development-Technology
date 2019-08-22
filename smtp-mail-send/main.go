package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hello there")
}

func send(body string) {
	from := "someone@somewhere.com"
	pass := "password"
	to := "someone@somewhere.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello Apu\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:25",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Print("smtp error: %s", err)
		return
	}
}
