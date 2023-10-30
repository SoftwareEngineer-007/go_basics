package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

// get-метод
func (e *Event) Title() string {
	return e.title
}

// set-метод
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 { // если поле title содержит более 30 символов,
		return errors.New("invalid title") // возвращается ошибка
	}
	e.title = title
	return nil
}
