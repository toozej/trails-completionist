package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"

	"github.com/toozej/trails-completionist/internal/generator"
	"github.com/toozej/trails-completionist/internal/parser"
	"github.com/toozej/trails-completionist/pkg/config"
	"github.com/toozej/trails-completionist/pkg/man"
	"github.com/toozej/trails-completionist/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:              "trails-completionist",
	Short:            "tools for tracking completion of trails",
	Long:             `A simple Golang application to parse a list of trails, then display that in a HTML table for ease of tracking completion of trails`,
	PersistentPreRun: rootCmdPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		config.GetEnvVars()
		inputFilename := viper.GetString("INPUT_FILENAME")
		fmt.Printf("Parsing filename: %s\n", inputFilename)

		trails, err := parser.ParseTrailsFromFile(inputFilename)
		if err != nil {
			log.Fatal(err)
		}
		if viper.GetBool("debug") {
			fmt.Printf("Parsed trails:\n %v\n", trails)
		}

		if err = generator.GenerateHTMLOutput("./out/html/index.html", trails); err != nil {
			fmt.Println("Error generating HTML output: ", err)
		} else {
			http.Handle("/", http.FileServer(http.Dir("./out/html")))
			http.ListenAndServe(":3000", nil)
		}
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
