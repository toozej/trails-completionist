package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	trailscompletionist "github.com/toozej/trails-completionist/internal/trails-completionist"
	"github.com/toozej/trails-completionist/pkg/config"
)

var FullCmd = &cobra.Command{
	Use:   "full",
	Short: "Run the full trails-completionist workflow",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.ConfigFromViper()
		debug := viper.GetBool("debug")
		if err := trailscompletionist.RunTrailsCompletionist(conf, debug); err != nil {
			log.Fatal(err)
		}
	},
}
