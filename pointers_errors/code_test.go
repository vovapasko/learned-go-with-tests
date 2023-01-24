package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(10)
		want := Bitcoin(90)

		assertNotError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "amount is greater than balance")
		assertBalance(t, wallet, startingBalance)
	})
}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error, but didn't want one")
	}
}
