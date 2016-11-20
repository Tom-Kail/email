package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Tom-Kail/email"
)

const (
	Addr     = "smtp.126.com"
	Host     = Addr + ":25"
	AuthName = "cccbackup@126.com"
	AuthPwd  = "ndasec123"
)

func main() {
	// gen attach
	fileName := "暴走大事件"
	f, err := os.Create(fileName)
	defer os.Remove(f.Name())
	defer f.Close()
	if err != nil {
		panic(err)
	}

	f.WriteString("实践出真知")
	m := email.NewMessage()
	m.Attach(&fileName)
	//	m.Attach(&fileName)
	auth := smtp.PlainAuth(m.Sender, AuthName, AuthPwd, Addr)
	// send email
	count := 4
	finish := make(chan bool)
	//	mail, _ := os.Create("mail.txt")
	//	mail.WriteString(msg)
	for i := 0; i < count; i++ {
		go func() {
			defer func() { finish <- true }()
			//send the email
			if err := email.SendMail(&auth, m); err != nil {
				fmt.Println(err)
				return
			}
		}()
	}

	for i := 0; i < count; i++ {
		<-finish
	}
}
