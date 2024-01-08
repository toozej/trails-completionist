package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

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
	Run:              rootCmdRun,
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	config.GetEnvVars()

	// gather raw trails
	inputFilename := viper.GetString("INPUT_FILENAME")
	fmt.Printf("Parsing filename: %s\n", inputFilename)
	rawTrails, err := parser.ParseTrailsFromRawInputFile(inputFilename)
	if err != nil {
		fmt.Println("Error parsing trails from raw input file: ", err)
		os.Exit(1)
	}
	if viper.GetBool("debug") {
		fmt.Printf("Parsed trails from raw input:\n %v\n", rawTrails)
	}

	// generate checklist
	checklistFilename := viper.GetString("CHECKLIST_FILENAME")
	fmt.Printf("Parsing filename: %s\n", checklistFilename)
	if err = generator.GenerateChecklist(checklistFilename, rawTrails); err != nil {
		fmt.Println("Error generating checklist: ", err)
		os.Exit(1)
	}

	// parse trails from checklist
	trails, err := parser.ParseTrailsFromChecklist(checklistFilename)
	if err != nil {
		fmt.Println("Error parsing trails from checklist: ", err)
		os.Exit(1)
	}

	// generate HTML table from checklist
	htmlFilename := viper.GetString("HTML_FILENAME")
	if err = generator.GenerateHTMLOutput("./out/html/"+htmlFilename, trails); err != nil {
		fmt.Println("Error generating HTML output file: ", err)
		os.Exit(1)
	} else {
		http.Handle("/", http.FileServer(http.Dir("./out/html")))
		server := &http.Server{
			Addr:         ":3000",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Error serving generated HTML file: ", err)
			os.Exit(1)
		}
	}
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
