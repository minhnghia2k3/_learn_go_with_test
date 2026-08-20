package main

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(10)
		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("widthdraw insufficient amount", func(t *testing.T) {
		initBalance := Bitcoin(20)

		wallet := Wallet{initBalance}
		want := initBalance
		err := wallet.Withdraw(100)

		assertError(t, err, ErrInsufficientAmount)
		assertBalance(t, wallet, want)
	})
}

func TestBitcoinString(t *testing.T) {
	t.Run("Bitcoin string", func(t *testing.T) {
		btc := Bitcoin(10)

		got := btc.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("not wanted an error")
	}
}
