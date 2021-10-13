package Entities

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	id          string
	userName    string
	email       string
	dateCreated time.Time
	password    string
	dob 		string
}

func NewCustomer(username, email string, dateCreated time.Time, password, dob string) Customer { 
	c := Customer { uuid.New().String(), username, email, dateCreated, password, dob  }
	return c
}