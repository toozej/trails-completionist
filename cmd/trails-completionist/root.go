package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"

	"github.com/toozej/trails-completionist/internal/parser"
	"github.com/toozej/trails-completionist/pkg/config"
	"github.com/toozej/trails-completionist/pkg/man"
	"github.com/toozej/trails-completionist/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:              "trails-completionist",
	Short:            "golang starter examples",
	Long:             `Examples of using math library, cobra and viper modules in golang`,
	PersistentPreRun: rootCmdPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		config.GetEnvVars()
		url := viper.GetString("URL")
		fmt.Printf("Parsing URL %s\n", url)

		trails, err := parser.ParseTrailsFromHTML(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(trails)

	},
}

func rootCmdPreRun(cmd *cobra.Command, args []string) {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	// create rootCmd-level flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug-level logging")

	// add sub-commands
	rootCmd.AddCommand(
		man.NewManCmd(),
		version.Command(),
	)
}
