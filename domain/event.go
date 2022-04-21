package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/guiaanonima/mitnick-event-bot/utils"
)

const (
	EventIdPositions          = 3
	EventIdLength             = EventIdPositions * 2
	EventNameMinLength        = 5
	EventNameMaxLength        = 100
	EventDescriptionMinLength = 10
	EventDescriptionMaxLength = 1000
)

type Event struct {
	ID          string
	Name        string
	Description string
	Date        time.Time
	Finished    bool
}

func NewEvent(id, name, description string, finished bool, date time.Time) (*Event, error) {
	event := &Event{
		ID:          id,
		Name:        name,
		Description: description,
		Date:        date,
		Finished:    finished,
	}

	if err := event.isValid(); err != nil {
		return nil, err
	}

	return event, nil
}

func (event *Event) isValid() error {
	if err := event.validateName(); err != nil {
		return err
	}

	if err := event.validateDescription(); err != nil {
		return err
	}

	return nil
}

func (event *Event) validateName() error {
	if utils.IsEmptyString(event.Name) {
		return errors.New("Event name is required")
	}

	if !utils.HaveMinLength(event.Name, EventNameMinLength) {
		return fmt.Errorf("Event name must have at least %d characters", EventNameMinLength)
	}

	if !utils.HaveMaxLength(event.Name, EventNameMaxLength) {
		return fmt.Errorf("Event name must have at most %d characters", EventNameMaxLength)
	}

	return nil
}

func (event *Event) validateDescription() error {
	if utils.IsEmptyString(event.Description) {
		return errors.New("Event description is required")
	}

	if !utils.HaveMinLength(event.Description, EventDescriptionMinLength) {
		return fmt.Errorf("Event description must have at least %d characters", EventDescriptionMinLength)
	}

	if !utils.HaveMaxLength(event.Description, EventDescriptionMaxLength) {
		return fmt.Errorf("Event description must have at most %d characters", EventDescriptionMaxLength)
	}

	return nil
}
