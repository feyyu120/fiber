package main

import "fmt"

type Notification interface {
	Send(s string) string
}

type Email struct{}

func (e Email) Send(s string) string {
	return s
}

type SMS struct{}

func (sm SMS) Send(s string) string {
	return s
}

func NewNotification(s string) Notification {
	switch s {
	case "sms":
		return SMS{}
	case "email":
		return Email{}
	default:
		return nil
	}
}

func main() {
	notif := NewNotification("sms")
	a := notif.Send("welcome")
	fmt.Println(a)
}
