package vo

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	MinLength = 5
	MaxLength = 30
)

var (
	ErrUsernameMinLength = fmt.Errorf("username must be at least %v characters", MinLength)
	ErrUsernameMaxLength = fmt.Errorf("username must be at least %v characters", MinLength)
	ErrUsernameValidate  = fmt.Errorf("username can only contain letters, numbers and underscores")

	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

type Username struct {
	value string
}

func NewUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)

	if len(username) < MinLength {
		return Username{}, ErrUsernameMinLength
	}
	if len(username) > MaxLength {
		return Username{}, ErrUsernameMaxLength
	}

	if !usernameRegex.MatchString(username) {
		return Username{}, ErrUsernameValidate
	}

	return Username{value: username}, nil
}

func (u Username) Value() string {
	return u.value
}

func (u Username) Equals(other Username) bool {
	return u.value == other.value
}
