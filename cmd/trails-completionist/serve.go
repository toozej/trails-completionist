package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	trailscompletionist "github.com/toozej/trails-completionist/internal/trails-completionist"
	"github.com/toozej/trails-completionist/pkg/config"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run web server to display generated HTML page and interact with the trails table",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf := config.ConfigFromViper()
		htmlFile := conf.HTMLFile
		if htmlFile == "" {
			return fmt.Errorf("htmlFile must be specified via flag or env var")
		}
		return trailscompletionist.ServeHTMLFile(htmlFile)
	},
}
