package domain

import (
	"fmt"
	"regexp"
)

var RegexEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func CheckPassword(password string) error {

	var CheckPasswordA2Z = regexp.MustCompile(`[A-Z]{1}`)
	var CheckPassworda2z = regexp.MustCompile(`[a-z]{1}`)
	var CheckPassword0to9 = regexp.MustCompile(`[0-9]{1}`)

	if !CheckPasswordA2Z.MatchString(password) || !CheckPassworda2z.MatchString(password) || !CheckPassword0to9.MatchString(password) {
		return fmt.Errorf("Password rules do not match, the password must contain an English uppercase, an English lowercase, and a number")
	}

	return nil
}
