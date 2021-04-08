package credentials_test

import (
	"testing"

	"github.com/AlecAivazis/survey/v2/terminal"
	expect "github.com/Netflix/go-expect"
	"github.com/bool64/ctxd"
	"github.com/stretchr/testify/assert"

	"github.com/nhatthm/n26prompt/credentials"
	"github.com/nhatthm/n26prompt/test"
)

func TestCredentialsProvider_Username(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		expect         func(c *expect.Console)
		expectedResult string
	}{
		{
			scenario: "username is entered at the first time",
			// nolint: errcheck,gosec
			expect: func(c *expect.Console) {
				c.ExpectString("Enter username (input is hidden) > ")
				c.Send("username")
				c.SendLine("")
				c.ExpectEOF()
			},
			expectedResult: "username",
		},
		{
			scenario: "username is skipped at the first time and then entered",
			// nolint: errcheck,gosec
			expect: func(c *expect.Console) {
				c.ExpectString("Enter username (input is hidden) > ")
				c.SendLine("")
				// Username is required, ask again.
				c.ExpectString("Enter username (input is hidden) > ")
				c.Send("username")
				c.SendLine("")
				c.ExpectEOF()
			},
			expectedResult: "username",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			test.Run(t, tc.expect, func(stdio terminal.Stdio) {
				p := credentials.New(
					credentials.WithStdio(stdio),
					credentials.WithLogger(ctxd.NoOpLogger{}),
				)

				// 1st try: read from input.
				result := p.Username()

				assert.Equal(t, tc.expectedResult, result)

				// 2nd try: read from memory
				result = p.Username()

				assert.Equal(t, tc.expectedResult, result)
			})
		})
	}
}

func TestCredentialsProvider_Password(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		scenario       string
		expect         func(c *expect.Console)
		expectedResult string
	}{
		{
			scenario: "password is entered at the first time",
			// nolint: errcheck,gosec
			expect: func(c *expect.Console) {
				c.ExpectString("Enter password (input is hidden) > ")
				c.Send("password")
				c.SendLine("")
				c.ExpectEOF()
			},
			expectedResult: "password",
		},
		{
			scenario: "password is skipped at the first time and then entered",
			// nolint: errcheck,gosec
			expect: func(c *expect.Console) {
				c.ExpectString("Enter password (input is hidden) > ")
				c.SendLine("")
				// Password is required, ask again.
				c.ExpectString("Enter password (input is hidden) > ")
				c.Send("password")
				c.SendLine("")
				c.ExpectEOF()
			},
			expectedResult: "password",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()

			test.Run(t, tc.expect, func(stdio terminal.Stdio) {
				p := credentials.New(
					credentials.WithStdio(stdio),
					credentials.WithLogger(ctxd.NoOpLogger{}),
				)

				// 1st try: read from input.
				result := p.Password()

				assert.Equal(t, tc.expectedResult, result)

				// 2nd try: read from memory
				result = p.Password()

				assert.Equal(t, tc.expectedResult, result)
			})
		})
	}
}
