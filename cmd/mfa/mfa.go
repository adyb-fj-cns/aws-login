package mfa

import (
	"github.com/adyb-fj-cns/aws-login/config"
	"github.com/adyb-fj-cns/aws-login/service/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// MFACmd returns MFA command
var MFACmd = &cobra.Command{
	Use:   "mfa",
	Short: "MFA Commands",
	Long:  `This subcommand does an mfa login`,
	Run: func(cmd *cobra.Command, args []string) {

		awsService := aws.NewService(
			viper.GetString("aws.profile"),
			viper.GetString("aws.mfaarn"),
			viper.GetString("aws.mfacode"),
			viper.GetString("aws.region"),
		)
		awsService.GenerateTemporaryCredentialsFromSTS()

	},
}

func init() {
	config.InitConfigFromFlags(MFACmd, config.SubCommandConfig("mfa"))
}
