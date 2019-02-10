package main

import (
	"github.com/gofrs/uuid"
)

type Event struct{
	AccId string
	Type string
}

type CreateEvent struct{
	Event
	AccName string
}


// Helper to create events

func newCreateAccountEvent(name string) CreateEvent{
	uid,_ := uuid.NewV4()

	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccId = uid.String()
	event.AccName = name
	return *event
}