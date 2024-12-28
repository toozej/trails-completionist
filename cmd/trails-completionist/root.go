package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/toozej/trails-completionist/internal/generator"
	"github.com/toozej/trails-completionist/internal/parser"
	"github.com/toozej/trails-completionist/pkg/config"
	"github.com/toozej/trails-completionist/pkg/man"
	"github.com/toozej/trails-completionist/pkg/version"
)

var conf config.Config

var rootCmd = &cobra.Command{
	Use:              "trails-completionist",
	Short:            "tools for tracking completion of trails",
	Long:             `A simple Golang application to parse a list of trails, then display that in a searchable HTML table for ease of tracking completion of trails`,
	PersistentPreRun: rootCmdPreRun,
	Run:              rootCmdRun,
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	if viper.GetBool("debug") {
		fmt.Printf("rootCmdRun: conf Config struct contains: %v\n", conf)
	}

	// gather raw trails
	if viper.GetBool("debug") {
		fmt.Printf("Parsing filename: %s\n", conf.InputFile)
	}
	rawTrails, err := parser.ParseTrailsFromRawInputFile(conf.InputFile)
	if err != nil {
		log.Fatal("Error parsing trails from raw input file: ", err)
	}
	if viper.GetBool("debug") {
		fmt.Printf("Parsed trails from raw input:\n %v\n", rawTrails)
	}

	// generate checklist
	if viper.GetBool("debug") {
		fmt.Printf("Parsing filename: %s\n", conf.ChecklistFile)
	}
	if err = generator.GenerateChecklist(conf.ChecklistFile, rawTrails); err != nil {
		log.Fatal("Error generating checklist: ", err)
	}

	// parse trails from checklist
	trails, err := parser.ParseTrailsFromChecklist(conf.ChecklistFile)
	if err != nil {
		log.Fatal("Error parsing trails from checklist: ", err)
	}

	if viper.GetBool("debug") {
		log.Println(trails)
	}

	// generate HTML table from checklist
	if err = generator.GenerateHTMLOutput(conf.HTMLFile, trails); err != nil {
		log.Fatal("Error generating HTML output file: ", err)
	} else if conf.Serve {
		htmlDir := filepath.Dir(conf.HTMLFile)
		http.Handle("/", http.FileServer(http.Dir(htmlDir)))
		server := &http.Server{
			Addr:         ":3000",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Error serving generated HTML file: ", err)
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
		log.Fatal(err.Error())
	}
}

func init() {
	_, err := maxprocs.Set()
	if err != nil {
		log.Error("Error setting maxprocs: ", err)
	}

	// get configuration from environment variables
	conf = config.GetEnvVars()

	// create rootCmd-level flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug-level logging")

	if conf.InputFile == "" {
		// optional flag for input file if not specified by env var
		rootCmd.Flags().StringVarP(&conf.InputFile, "inputFile", "i", "", "Input file")
	}

	if conf.ChecklistFile == "" {
		// optional flag for checklist file if not specified by env var
		rootCmd.Flags().StringVarP(&conf.ChecklistFile, "checklistFile", "c", "", "Checklist file")
	}

	if conf.HTMLFile == "" {
		// optional flag for html filename if not specified by env var
		rootCmd.Flags().StringVarP(&conf.HTMLFile, "htmlFile", "o", "", "HTML file")
	}

	if !conf.Serve {
		// optional flag to serve the generated HTML file if not specified by env var
		rootCmd.Flags().BoolVarP(&conf.Serve, "serve", "s", false, "Serve the generated HTML file")
	}

	// add sub-commands
	rootCmd.AddCommand(
		man.NewManCmd(),
		version.Command(),
	)
}
