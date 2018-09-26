package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Email string
	Name  string
}

var DataBase = []User{
	{Email: "alexander.davis@example.com", Name: "Alexander Davis"},
	{Email: "alexander.jackson@example.com", Name: "Alexander Jackson"},
	{Email: "avry.williams@example.com", Name: "Avery Williams"},
	{Email: "charlotte.smith@example.com", Name: "Charlotte Smith"},
	{Email: "daniel.miller@example.com", Name: "Daniel Miller"},
	{Email: "ella.smith@example.com", Name: "Ella Smith"},
	{Email: "jacob.white@example.com", Name: "Jacob White"},
	{Email: "james.martinez@example.com", Name: "James Martinez"},
	{Email: "james.miller@example.com", Name: "James Miller"},
	{Email: "jayden.jackson@example.com", Name: "Jayden Jackson"},
	{Email: "liam.robinson@example.com", Name: "Liam Robinson"},
	{Email: "mason.martin@example.com", Name: "Mason Martin"},
	{Email: "matthew.jackson@example.com", Name: "Matthew Jackson"},
	{Email: "mia.smith@example.com", Name: "Mia Smith"},
	{Email: "michael.white@example.com", Name: "Michael White"},
	{Email: "natalie.martin@example.com", Name: "Natalie Martin"},
	{Email: "sofia.garcia@example.com", Name: "Sofia Garcia"},
	{Email: "william.brown@example.com", Name: "William Brown"},
}

type Worker struct {
	users []User
	ch    chan *User
	name  string
}

func NewWorker(users []User, ch chan *User, name string) *Worker {
	return &Worker{users: users, ch: ch, name: name}
}

func (w *Worker) Find(email string) {
	// Since this function is using a channel, it no longer needs a return.
	for i := range w.users {
		user := w.users[i]
		if strings.Contains(user.Email, email) {
			log.Println(w.name)
			w.ch <- &user
		}
	}
}

func main() {
	// Allows for command line arguments
	email := os.Args[1]
	ch := make(chan *User)

	log.Printf("Looking for %s", email)
	go NewWorker(DataBase[:6], ch, "Ch #1").Find(email)
	go NewWorker(DataBase[6:12], ch, "Ch #2").Find(email)
	go NewWorker(DataBase[12:], ch, "Ch #3").Find(email)

	// The channel returns but is never closed with the line below.
	// user := <-ch

	for {

		select {
		case user := <-ch:
			log.Printf("The email %s is owned by %s", user.Email, user.Name)
		case <-time.After(1 * time.Second):
			// log.Printf("The email %s was not found", email)
			return
		}
	}
}
