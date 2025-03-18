package pointers

import "testing"

func TestWallet(t *testing.T) {

	//wallet := Wallet{}
	//
	//wallet.Deposit(10)
	//
	//got := wallet.Balance()
	//want := Bitcoin(10)
	//
	//if got != want {
	//	t.Errorf("got %s want %s", got, want)
	//}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdrawl(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdrawl(Bitcoin(100))
		want := Bitcoin(20)

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
		//assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
