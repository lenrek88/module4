package model

import (
	"encoding/json"
	"fmt"
	"lenrek88/internal/testutil"
	"testing"
)

func validFUserFixture() User {
	data := testutil.LoadFixture("model", "userFixture_valide.json")

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		panic(fmt.Sprintf("failed to parse fixture: %v", err))
	}
	return user
}

func emptyNameUserFixture() User {
	f := validFUserFixture()
	f.Name = ""
	return f
}

func parseEmailUserFixture() User {
	f := validFUserFixture()
	f.Email = "mail.ru"
	return f
}

func TestUserValidate(t *testing.T) {

	t.Run("Valid user", func(t *testing.T) {
		user := validFUserFixture()
		got := user.Validate()
		if got != nil {
			t.Fatalf("unexpected error, got %q", got)
		}
	})

	t.Run("Name Empty", func(t *testing.T) {
		user := emptyTitleFilmFixture()
		got := user.Validate()
		if got == nil {
			t.Fatalf("got %q want error Name do not Empty", got)
		}
	})

	t.Run("Parse Email", func(t *testing.T) {
		user := parseEmailUserFixture()
		got := user.Validate()
		if got == nil {
			t.Fatalf("got %q want error for parse email", got)
		}
	})

}
