package cli

import (
	"hl7/internal/cli/gen"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func Execute() error {
	cmd := &cobra.Command{
		Use:   "hl7",
		Short: "A brief description of your application",
		Long:  "",
	}

	cmd.AddCommand(
		gen.NewCmdGenerate(),
		gen.NewCmdDownload(),
	)

	if err := cmd.Execute(); err != nil {
		if err == promptui.ErrInterrupt {
			err = nil
		}
		return err
	}

	return nil
}
