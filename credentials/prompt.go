package credentials

import (
	"context"
	"errors"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/bool64/ctxd"
	"github.com/nhatthm/n26api"
	"github.com/nhatthm/surveymock/options"
	"github.com/nhatthm/surveymock/options/cobra"
)

var _ n26api.CredentialsProvider = (*Prompt)(nil)

// Option configures Prompt.
type Option func(p *Prompt)

// Prompt provides the credentials using cli prompt.
type Prompt struct {
	options []survey.AskOpt
	logger  ctxd.Logger

	username string
	password string
}

func (p *Prompt) prompt(message string, response interface{}) error {
	return survey.AskOne(&survey.Password{Message: message}, response, p.options...)
}

// Username provides a username via cli prompt.
func (p *Prompt) Username() string {
	if p.username != "" {
		return p.username
	}

	message := "Enter username (input is hidden) >"

	if err := p.prompt(message, &p.username); err != nil {
		if !errors.Is(err, terminal.InterruptErr) {
			p.logger.Error(context.Background(), "could not read username", "error", err)
		}

		return ""
	}

	return p.username
}

// Password provides a password via cli prompt.
func (p *Prompt) Password() string {
	if p.password != "" {
		return p.password
	}

	message := "Enter password (input is hidden) >"

	if err := p.prompt(message, &p.password); err != nil {
		if !errors.Is(err, terminal.InterruptErr) {
			p.logger.Error(context.Background(), "could not read password", "error", err)
		}

		return ""
	}

	return p.password
}

// New initiates a new CredentialsProvider.
func New(options ...Option) *Prompt {
	p := &Prompt{
		options: []survey.AskOpt{
			survey.WithValidator(survey.Required),
		},
		logger: ctxd.NoOpLogger{},
	}

	for _, o := range options {
		o(p)
	}

	return p
}

// WithStdio sets stdio for the prompts.
func WithStdio(stdio terminal.Stdio) Option {
	return func(p *Prompt) {
		p.options = append(p.options, options.WithStdio(stdio))
	}
}

// WithStdioProvider sets stdio for the prompts.
func WithStdioProvider(stdio cobra.StdioProvider) Option {
	return func(p *Prompt) {
		p.options = append(p.options, cobra.WithStdioProvider(stdio))
	}
}

// WithLogger sets logger for Prompt.
func WithLogger(logger ctxd.Logger) Option {
	return func(p *Prompt) {
		p.logger = logger
	}
}

// WithCredentialsAtLast sets CredentialsProvider as a n26api.CredentialsProvider.
func WithCredentialsAtLast(options ...Option) n26api.Option {
	return n26api.WithCredentialsProviderAtLast(New(options...))
}
