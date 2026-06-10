	package domain

	import (
		"fmt"
		"regexp"

		core_errors "github.com/g3nd4ch/todo-app/internal/core/errors"
	)

	type User struct {
		ID          int
		Version     int
		FullName    string
		Phonenumber *string
	}

	func NewUser(id, version int, fullName string, phoneNumber *string) User {
		return User{
			ID:          id,
			Version:     version,
			FullName:    fullName,
			Phonenumber: phoneNumber,
		}
	}
	func NewUserUninitialized(fullName string, phoneNumber *string) User {
		return NewUser(UninitializedID, UninitializedVersion, fullName, phoneNumber)
	}

	func (u *User) Validate() error {
		fullNameLength := len([]rune(u.FullName))
		if fullNameLength < 3 || fullNameLength > 100 {
			return fmt.Errorf(
				"invalid `FullName` len: %d: %w",
				fullNameLength,
				core_errors.ErrInvalidArgument,
			)
		}

		if u.Phonenumber != nil {
			phoneNumberLen := len([]rune(*u.Phonenumber))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf(
					"invalid phoneNumber len: %d :%w",
					phoneNumberLen,
					core_errors.ErrInvalidArgument,
				)
			}

			re := regexp.MustCompile(`^\+[0-9]+$`)

			if !re.MatchString(*u.Phonenumber) {
				return fmt.Errorf("invalid PhoneNumber format: %w",
					core_errors.ErrInvalidArgument,
				)
			}
		}
		return nil
	}
