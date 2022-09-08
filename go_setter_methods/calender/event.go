package calender

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string  // Field won't be exported
	Date   // Embeded/ Anonymous struct
}

// getter for title field
func (e *Event) Title() string {
	return e.title
}

// setter for title field
func (e *Event) SetTitle(t string) error {
	if utf8.RuneCountInString(t) > 15 {
		return errors.New("title should 30 characters or less")
	}
	e.title = t

	return nil
}