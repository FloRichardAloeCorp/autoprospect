package cmd

import (
	"fmt"
	"os"

	fromfile "github.com/FloRichardAloeCorp/autoprospect/cmd/from_file"
	"github.com/FloRichardAloeCorp/autoprospect/cmd/search"
	"github.com/spf13/cobra"
)

var (
	category   string
	activity   string
	department string
	from       string

	root = &cobra.Command{
		Use:   "autoprospect",
		Short: "Find your next propects",
	}
)

func Execute() {
	search.Init(root)
	fromfile.Init(root)

	if err := root.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
