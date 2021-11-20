package creditcard

import "errors"

type card struct {
	number string
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, errors.New("empty number was passed")
	}

	return card{number}, nil
}

func (c card) GetNumber() string {
	return c.number
}
