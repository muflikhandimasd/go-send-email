package main

import (
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error load env", err)
	}
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), port, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"))
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USERNAME"))
	to := "email@email.com"
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Friends</b>!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
