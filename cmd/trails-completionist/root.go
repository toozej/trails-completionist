package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/toozej/trails-completionist/pkg/config"
	"github.com/toozej/trails-completionist/pkg/man"
	"github.com/toozej/trails-completionist/pkg/version"
)

var conf config.Config

var rootCmd = &cobra.Command{
	Use:              "trails-completionist",
	Short:            "tools for tracking completion of trails",
	Long:             `A Golang application to parse a list of trails from a directory of track files, then display the found trails in a searchable HTML table for ease of tracking completion.`,
	PersistentPreRun: rootCmdPreRun,
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

	if conf.TrackFiles == "" {
		// optional flag for track files if not specified by env var
		rootCmd.PersistentFlags().StringVarP(&conf.TrackFiles, "trackFiles", "t", "", "Track files")
	}

	if conf.OSMRegionFile == "" {
		// optional flag for OSM region file if not specified by env var
		rootCmd.PersistentFlags().StringVarP(&conf.OSMRegionFile, "osmRegionFile", "r", "", "OSM region file")
	}

	if conf.InputFile == "" {
		// optional flag for input file if not specified by env var
		rootCmd.PersistentFlags().StringVarP(&conf.InputFile, "inputFile", "i", "", "Input file")
	}

	if conf.ChecklistFile == "" {
		// optional flag for checklist file if not specified by env var
		rootCmd.PersistentFlags().StringVarP(&conf.ChecklistFile, "checklistFile", "c", "", "Checklist file")
	}

	if conf.HTMLFile == "" {
		// optional flag for html filename if not specified by env var
		rootCmd.PersistentFlags().StringVarP(&conf.HTMLFile, "htmlFile", "o", "", "HTML file")
	}

	if !conf.Serve {
		// optional flag to serve the generated HTML file if not specified by env var
		rootCmd.PersistentFlags().BoolVarP(&conf.Serve, "serve", "s", false, "Serve the generated HTML file")
	}

	// add sub-commands from separate files
	rootCmd.AddCommand(
		man.NewManCmd(),
		version.Command(),
		ConvertCmd,
		OsmExportCmd,
		ParseGPXCmd,
		GenerateChecklistCmd,
		GenerateHTMLCmd,
		ServeCmd,
		FullCmd,
	)
}
