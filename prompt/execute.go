package prompt

import (
	"github.com/c-bata/go-prompt"
	"github.com/masterzen/dashlane-cli/command"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "uki", Description: "Manage computer registration."},
		{Text: "vault", Description: "Manage the vault."},
		{Text: "version", Description: "Displays dahslane-cli version."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// Execute the prompt
func Execute() {
	p := prompt.New(
		command.Execute,
		completer,
		prompt.OptionTitle("dashlane-cli: command line for dashlane"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
