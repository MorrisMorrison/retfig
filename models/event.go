package models

type Event struct {
	Name      string
	Owner     User
	Recipient User
	Members   []User
}
