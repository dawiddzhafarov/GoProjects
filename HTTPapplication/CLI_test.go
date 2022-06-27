package poker_test

import (
	poker "HTTPapplication"
	"bytes"
	"io"
	"strings"

	"testing"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}
var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartedWith    int
	StartedCalled  bool
	FinishedCalled bool
	FinishedWith   string
	BlindAlert     []byte
}

func (g *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartedCalled = true
	alertsDestination.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishedCalled = true
}

func TestCLI(t *testing.T) {
	t.Run("start a game with 3 playerd and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCli(in, stdout, game)

		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCli(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCli(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputMsg)
	})
	t.Run("it prints an error when wrong finish phrase is used", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("5", "Floyd is a killer")

		cli := poker.NewCli(in, stdout, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishedCalled {
		t.Errorf("game not finished but should be")
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartedCalled {
		t.Errorf("game has started but shouldnt")
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, want string) {
	t.Helper()
	if game.FinishedWith != want {
		t.Errorf("game finished with %q, wanted %q", game.FinishedWith, want)
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, want int) {
	t.Helper()
	if game.StartedWith != want {
		t.Errorf("Game started with %d, wanted %d", game.StartedWith, want)
	}
}

func userSends(msgs ...string) io.Reader {
	return strings.NewReader(strings.Join(msgs, "\n"))
}
func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	if got.Amount != want.Amount {
		t.Errorf("got %d amount of chips, but wanted %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("alert scheduled at %v, but wanted at %v", got.At, want.At)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got  %q, expected %+v", got, messages)
	}
}
