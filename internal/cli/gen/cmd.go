package gen

import (
	"github.com/spf13/cobra"
)

func NewCmdGenerate() *cobra.Command {
	var opts struct {
		GoPackage       string
		HL7Version      string
		SchemaDirectory string
		OutputDirectory string
	}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Short string.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Generate(HL7Version(opts.HL7Version), opts.SchemaDirectory, opts.OutputDirectory, &Options{GoPackage: opts.GoPackage})
		},
	}

	cmd.Flags().StringVar(&opts.GoPackage, "package", "", "the name of the go package")
	cmd.Flags().StringVar(&opts.HL7Version, "version", "2.8", "the version of the HL7 standard")
	cmd.Flags().StringVar(&opts.SchemaDirectory, "schema", "", "the directory to download the schema to")
	cmd.Flags().StringVar(&opts.OutputDirectory, "output", "", "the directory to write the generated code to")
	cmd.MarkFlagRequired("schema")
	cmd.MarkFlagRequired("output")

	return cmd
}

func NewCmdDownload() *cobra.Command {
	var opts struct {
		HL7Version      string
		SchemaDirectory string
	}

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Short string.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return DownloadContext(HL7Version(opts.HL7Version), opts.SchemaDirectory)
		},
	}

	cmd.Flags().StringVar(&opts.HL7Version, "version", "2.8", "the version of the HL7 standard")
	cmd.Flags().StringVar(&opts.SchemaDirectory, "output", "", "the directory to download the schema files to")
	cmd.MarkFlagRequired("output")

	return cmd
}
