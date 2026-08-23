package main

import (
	"errors"
	"slices"
	"testing"
)

type SpyUserStore struct {
	user map[string]User
}

func NewSpyUserStore() *SpyUserStore {
	return &SpyUserStore{user: make(map[string]User)}
}

func (s *SpyUserStore) GetUser(email string) *User {
	user := s.user[email]
	return &user
}

func (s *SpyUserStore) Save(email string) error {
	if _, ok := s.user[email]; ok {
		return ErrExistingUser
	}

	s.user[email] = User{Email: email}
	return nil
}

type SpyMailer struct {
	sentTo []string
}

func NewSpyMailer() *SpyMailer {
	return &SpyMailer{}
}

func (m *SpyMailer) SendWelcome(email string) error {
	m.sentTo = append(m.sentTo, email)
	return nil
}

func TestRegister(t *testing.T) {
	t.Run("register new user", func(t *testing.T) {
		email := "test@email.com"
		store := NewSpyUserStore()
		mailer := NewSpyMailer()

		svc := NewRegistrationService(store, mailer)

		err := svc.Register(email)

		assertError(t, err, nil)
		assertEqual(t, store, email)
		assertSentTo(t, mailer.sentTo, email)
	})

	t.Run("register existing user", func(t *testing.T) {
		email := "test@email.com"
		store := NewSpyUserStore()
		mailer := NewSpyMailer()

		svc := NewRegistrationService(store, mailer)

		store.Save(email)
		err := svc.Register(email)

		assertError(t, err, ErrExistingUser)

		if len(mailer.sentTo) != 0 {
			t.Errorf("no expected to send email")
		}
	})

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if want != nil && got == nil {
		t.Fatal("expected an error but got none")
	}

	if !errors.Is(got, want) {
		t.Errorf("expected got error %v, but want %v", got, want)
	}
}

func assertEqual(t testing.TB, store UserStore, want string) {
	t.Helper()

	got := store.GetUser(want).Email

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertSentTo(t testing.TB, sentTo []string, want string) {
	t.Helper()

	if !slices.Contains(sentTo, want) {
		t.Errorf("expected welcome mail for %q, but got none", want)
	}
}
