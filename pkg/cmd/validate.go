package cmd

import (
	"errors"

	cli "github.com/spf13/cobra"
	"github.com/timsolov/boilr/pkg/util/exit"
	"github.com/timsolov/boilr/pkg/util/validate"
)

var (
	// ErrTemplateInvalid indicates that the template is invalid.
	ErrTemplateInvalid = errors.New("validate: given template is invalid")
)

// Validate contains the cli-command for validating templates.
var Validate = &cli.Command{
	Use:   "validate <template-path>",
	Short: "Validate a local project template",
	Run: func(c *cli.Command, args []string) {
		if len(args) == 0 {
			_ = c.Usage()
			return
		}

		MustValidateArgs(args, []validate.Argument{
			{Name: "template-path", Validate: validate.UnixPath},
		})

		templatePath := args[0]

		MustValidateTemplate(templatePath)

		exit.OK("Template is valid")
	},
}
