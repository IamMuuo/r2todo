package models

import (
	"strconv"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Completed   bool
	DateCreated time.Time
	DueDate     time.Time
}

func serializeTime(t time.Time) string {
	return t.Format(time.RFC3339) + "Z"
}

func deserializeTime(s string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	t, err := time.Parse(layout, s[:len(s)-1])
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// Serializes a todo into an  slice of strings
// suitable for storing in a CSV
func (t *Todo) Serialize() []string {
	return []string{
		strconv.Itoa(t.ID),
		t.Description,
		strconv.FormatBool(t.Completed),
		serializeTime(t.DateCreated),
		serializeTime(t.DueDate),
	}
}

// Deserialize a todo from its string representation
// the [data] must contain 5 items
func Deserialize(data []string) (*Todo, error) {
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, err
	}

	complete, err := strconv.ParseBool(data[2])
	if err != nil {
		return nil, err
	}

	dateCreated, err := deserializeTime(data[3])
	if err != nil {
		return nil, err
	}
	dueDate, err := deserializeTime(data[4])
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:          id,
		Description: data[1],
		Completed:   complete,
		DateCreated: dateCreated,
		DueDate:     dueDate,
	}, nil
}
