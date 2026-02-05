package users

import (
	"fmt"
	"strings"
)

func ValidateCreateUser(dto CreateUserDTO) error {
	if !strings.Contains(dto.Email, "@") {
		return fmt.Errorf("invalid email")
	}
	if len(dto.Password) < 6 {
		return fmt.Errorf("password too short")
	}
	return nil
}
