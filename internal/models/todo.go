package models

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
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

// / Creates a new todo item
func NewTodoItem(id int, descr string) Todo {
	return Todo{
		ID:          id,
		Description: descr,
		Completed:   false,
		DateCreated: time.Now(),
	}
}

// Serializes a todo into an  slice of strings
// suitable for storing in a CSV
func (t *Todo) SerializeTodo() []string {
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
func DeserializeTodo(data []string) (*Todo, error) {
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

// / Displays a todo item in a neatly formatted manner
func DisplayTodos(todos *[]Todo) {
	w := tabwriter.NewWriter(os.Stdout,
		0, 0,
		1,
		' ',
		tabwriter.DiscardEmptyColumns|tabwriter.StripEscape)

	fmt.Fprintln(w, "ID\tDescription\tComplete\tDue Date\tDate Created")

	for _, v := range *todos {
		fmt.Fprintf(w, "%d\t%s\t%v\t%s\t%s\n",
			v.ID, v.Description, v.Completed,
			timediff.TimeDiff(v.DueDate),
			timediff.TimeDiff(v.DateCreated))
	}

	w.Flush()
}
