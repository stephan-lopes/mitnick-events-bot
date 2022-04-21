package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/guiaanonima/mitnick-event-bot/utils"
)

func TestCreateEvent(t *testing.T) {
	createdDate := time.Now()
	eventID, _ := utils.GenerateToken(EventIdPositions)

	event, err := NewEvent(eventID, "Test Event", "This is a test event", false, createdDate)

	if err != nil {
		t.Errorf("Error creating event: %s", err)
	}
	if event.Name != "Test Event" {
		t.Errorf("Expected event name to be 'Test Event', got '%s'", event.Name)
	}
	if event.Description != "This is a test event" {
		t.Errorf("Expected event description to be 'This is a test event', got '%s'", event.Description)
	}
	if event.ID != eventID {
		t.Errorf("Expected event ID to be '%s', got '%s'", eventID, event.ID)
	}
}

func TestCreateEventWithoutName(t *testing.T) {
	createdDate := time.Now()
	eventID, _ := utils.GenerateToken(EventIdPositions)

	_, err := NewEvent(eventID, "", "This is a test event", false, createdDate)
	if err == nil {
		t.Errorf("Expected error creating event without name, got nil")
	}
}

func TestCreateEventWithNameLessThanMinLength(t *testing.T) {
	createdDate := time.Now()
	eventID, _ := utils.GenerateToken(EventIdPositions)

	_, err := NewEvent(eventID, "Te", "This is a test event", false, createdDate)
	if err == nil {
		t.Errorf("Expected error creating event with name less than %d characters, got nil", EventNameMinLength)
	}
}

func TestCreateEventWithNameMoreThanMaxLength(t *testing.T) {
	createdDate := time.Now()
	eventID, _ := utils.GenerateToken(EventIdPositions)

	var name string

	for i := 0; i < EventNameMaxLength+1; i++ {
		name = fmt.Sprintf("%s%d", name, i)
	}

	_, err := NewEvent(eventID, name, "This is a test event", false, createdDate)
	if err == nil {
		t.Errorf("Expected error creating event with name more than %d characters, got nil", EventNameMaxLength)
	}
}
