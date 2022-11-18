package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+`).MatchString
)

func ValidateString(value string, minLenght int, maxLength int) error {
	n := len(value)

	if n < minLenght || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLenght, maxLength)
	}

	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(username) {
		return fmt.Errorf("must contain digits, lowercase letters or underscores")
	}

	return nil
}

func ValidatePassword(password string) error {
	return ValidateString(password, 6, 100)
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email address")
	}

	return nil
}

func ValidateFullname(fullname string) error {
	if err := ValidateString(fullname, 3, 100); err != nil {
		return err
	}

	if !isValidFullname(fullname) {
		return fmt.Errorf("must contain only letters or spaces")
	}

	return nil
}
