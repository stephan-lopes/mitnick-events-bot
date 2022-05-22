package domain

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/guiaanonima/mitnick-event-bot/utils"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestNewEvent(t *testing.T) {

	eventID, _ := utils.GenerateToken(EventIdPositions)
	validName := randName(EventNameMinLength)
	createdDate := time.Now()

	type args struct {
		id          string
		name        string
		description string
		finished    bool
		date        time.Time
	}

	tests := []struct {
		name string
		args args
		want *Event
	}{
		{
			name: "Create event with valid parameters",
			args: args{
				id:          eventID,
				name:        validName,
				description: "This is a test event",
				finished:    false,
				date:        createdDate,
			},
			want: &Event{
				ID:          eventID,
				Name:        validName,
				Description: "This is a test event",
				Date:        createdDate,
				Finished:    false,
			},
		},
		{
			name: "Create event withouth name",
			args: args{
				id:          eventID,
				name:        "",
				description: "This is a test event",
				finished:    false,
				date:        createdDate,
			},
			want: nil,
		},
		{
			name: "Create event with name less than min characters",
			args: args{
				id:          eventID,
				name:        randName(EventNameMinLength - 1),
				description: "This is a test event",
				finished:    false,
				date:        createdDate,
			},
			want: nil,
		},
		{
			name: "Create event with name more than max characters",
			args: args{
				id:          eventID,
				name:        randName(EventNameMaxLength + 1),
				description: "This is a test event",
				finished:    false,
				date:        createdDate,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewEvent(tt.args.id, tt.args.name, tt.args.description, tt.args.finished, tt.args.date)

			if !cmp.Equal(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func randName(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = ALPHABET[rand.Intn(len(ALPHABET))]
	}
	return string(b)
}
