package poker_test

import (
	poker "HTTPapplication"
	"bytes"

	"strings"
	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(dummyBlindAlerter, playerStore)
		cli := poker.NewCli(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(dummyBlindAlerter, playerStore)
		cli := poker.NewCli(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
	/*t.Run("it schedules printing of blind values", func(t *testing.T) {
			in := strings.NewReader("5\n")
			playerStore := &poker.StubPlayerStore{}
			blindAlerter := &SpyBlindAlerter{}
			game := poker.NewGame(blindAlerter, playerStore)
			cli := poker.NewCli(in, dummyStdOut, game)
			cli.PlayPoker()

			cases := []scheduledAlert{
				{0 * time.Minute, 100},
				{10 * time.Minute, 200},
				{20 * time.Minute, 300},
				{30 * time.Minute, 400},
				{40 * time.Minute, 500},
				{50 * time.Minute, 600},
				{60 * time.Minute, 800},
				{70 * time.Minute, 1000},
				{80 * time.Minute, 2000},
				{90 * time.Minute, 4000},
				{100 * time.Minute, 8000},
			}

			for i, want := range cases {
				t.Run(fmt.Sprint(want), func(t *testing.T) {

					if len(blindAlerter.alerts) <= i {
						t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
					}

					got := blindAlerter.alerts[i]
					assertScheduledAlert(t, got, want)

				})
			}
		})
		t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
			stdout := &bytes.Buffer{}
			in := strings.NewReader("7\n")
			blindAlerter := &SpyBlindAlerter{}
			game := poker.NewGame(blindAlerter, dummyPlayerStore)
			cli := poker.NewCli(in, stdout, game)
			cli.PlayPoker()

			got := stdout.String()
			want := poker.PlayerPrompt

			if got != want {
				t.Errorf("got %q, but wanted %q", got, want)
			}

			cases := []scheduledAlert{
				{0 * time.Second, 100},
				{12 * time.Minute, 200},
				{24 * time.Minute, 300},
				{36 * time.Minute, 400},
			}

			for i, want := range cases {
				t.Run(fmt.Sprint(want), func(t *testing.T) {

					if len(blindAlerter.alerts) <= i {
						t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
					}

					got := blindAlerter.alerts[i]
					assertScheduledAlert(t, got, want)
				})
			}


		})

	}

	*/
}
func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	if got.Amount != want.Amount {
		t.Errorf("got %d amount of chips, but wanted %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("alert scheduled at %v, but wanted at %v", got.At, want.At)
	}
}

