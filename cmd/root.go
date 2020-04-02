package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/adyb-fj-cns/aws-login/cmd/mfa"
	"github.com/adyb-fj-cns/aws-login/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "aws-login",
	Short: "AWS Login tool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws-login version: 0.1.0 by Ady Buxton")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	config.InitConfigFromFlags(RootCmd, config.GlobalConfig)

	RootCmd.AddCommand(mfa.MFACmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Following is the order of priority
	//1. explicit call to Set
	//2. flag
	//3. env
	//4. config
	//5. key/value store, ie consul
	//6. default

	//initConfigFromEnv()
	initConfigFromFile()
}

func initConfigFromEnv() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv() // read in environment variables that match
}

func initConfigFromFile() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigType("toml")
	if runtime.GOOS == "windows" {
		viper.SetConfigFile(fmt.Sprintf("%s/.aws-login", home))
	} else {
		viper.SetConfigFile(fmt.Sprintf("$HOME/.aws-login"))
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
