package main

import "fmt"

const (
	ErrNotFound = ErrDicionary("could not find the word you are looking for")
	ErrWordExist = ErrDicionary("cannot add the word as it already exists")
	ErrWordNonexistent = ErrDicionary("could not update the word as it does not exist")
)

type Dicionary map[string]string

type ErrDicionary string

func (e ErrDicionary) Error() string {
	return string(e)
}

func Search(dicionary map[string]string, word string) string {
	if valor, ok := dicionary[word]; ok {
		return valor
	}
	return ""
}

func (d Dicionary) Search(word string) (string, error) {
	definition, existe := d[word]
	if !existe {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dicionary) Adds(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err

	}

	return nil
}

func (d Dicionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNonexistent
	case nil:
		d[word] = definition
	default:
		return err

	}

	return nil
}

func (d Dicionary) Delete(word string) {
	delete(d, word)
}

func main() {
	fmt.Println("Test maps")
}
