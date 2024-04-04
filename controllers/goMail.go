package controllers

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func SendMail() {
	abc := gomail.NewMessage()

	abc.SetHeader("From", "saputromoedwan@gmail.com")
	abc.SetHeader("To", "waluyajuang330@gmail.com")
	abc.SetHeader("Subject", "Test subject")
	abc.SetBody("text/html", "Hello <b>World</b>")

	a := gomail.NewDialer("smtp.gmail.com", 587, "saputromoedwan@gmail.com", "Waluyagantenk")

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}

}
