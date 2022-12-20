package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

/*
Exercise 3. Rather than returning the first error found, return back a single error that contains all errors
discovered during validation. Update the code in +main()+ to properly handle multiple errors.
*/
var ErrInvalidID = errors.New("invalid ID")

type EmptyFieldError struct {
	FieldName string
}

func (fe EmptyFieldError) Error() string {
	return fe.FieldName
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)
		if err != nil {
			// use a type switch to detect an error that can be unwrapped into multiple errors
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				// if we have one use a loop to walk through allErrors and build a single string
				allErrors := err.Unwrap()
				var messages []string
				for _, e := range allErrors {
					messages = append(messages, processError(e, emp))
				}
				message = message + " allErrors: " + strings.Join(messages, ", ")
			default:
				message = message + " error: " + processError(err, emp)
			}
		}
		fmt.Println(message)
	}
}

/*
processError is a method for converting an error into a message. It has special handling for ErrInvalidID and
EmptyFieldError.
*/
func processError(err error, emp Employee) string {
	var fieldErr EmptyFieldError
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("invalid ID: %s", emp.ID)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	// use allErrors to gather the errors
	var allErrors []error
	if len(e.ID) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "ID"})
	} else if !validID.MatchString(e.ID) {
		// append on to allErrors
		allErrors = append(allErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "FirstName"})
	}
	if len(e.LastName) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "LastName"})
	}
	if len(e.Title) == 0 {
		// append on to allErrors
		allErrors = append(allErrors, EmptyFieldError{FieldName: "Title"})
	}
	// depending on the number of errors found, return nil, the single error, or use
	// errors.Join to return all errors within a single error
	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}
