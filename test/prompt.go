package test

import (
	"bytes"
	"testing"

	"github.com/AlecAivazis/survey/v2/core"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
	"github.com/stretchr/testify/require"
)

// nolint: gochecknoinits
func init() {
	// Disable color output for all prompts to simplify testing.
	core.DisableColor = true
}

// Stdio returns terminal.Stdio from expected.Console.
func Stdio(c *expect.Console) terminal.Stdio {
	return terminal.Stdio{In: c.Tty(), Out: c.Tty(), Err: c.Tty()}
}

// Run runs test with expected.Console.
func Run(t *testing.T, procedure func(*expect.Console), test func(terminal.Stdio)) {
	t.Helper()

	// Multiplex output to a buffer as well for the raw bytes.
	buf := new(bytes.Buffer)
	c, state, err := vt10x.NewVT10XConsole(expect.WithStdout(buf))
	require.Nil(t, err)

	defer c.Close() // nolint: errcheck

	done := make(chan struct{})

	go func() {
		defer close(done)
		procedure(c)
	}()

	test(Stdio(c))

	// Close the slave end of the pty, and read the remaining bytes from the master end.
	err = c.Tty().Close()
	require.NoError(t, err)
	<-done

	t.Logf("Raw output: %q", buf.String())

	// Dump the terminal's screen.
	t.Logf("\n%s", expect.StripTrailingEmptyLines(state.String()))
}
