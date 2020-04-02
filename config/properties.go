package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Property is a string property
type Property struct {
	Name         string
	Description  string
	Value        string
	Type         string
	Global       bool
	SubCommand   string
	DefaultValue string
	Flag         string
	ShortFlag    string
}

// Config is a list of properties
var Config = []Property{

	Property{
		Name:         "aws.region",
		Description:  "AWS Region",
		Global:       false,
		SubCommand:   "mfa",
		Type:         "string",
		DefaultValue: "eu-west-1",
		Flag:         "region",
		ShortFlag:    "r",
	},
	Property{
		Name:         "aws.profile",
		Description:  "AWS Profile",
		Global:       false,
		SubCommand:   "mfa",
		Type:         "string",
		DefaultValue: "default",
		Flag:         "profile",
		ShortFlag:    "p",
	},
	Property{
		Name:         "aws.setdefault",
		Description:  "Set as default AWS Profile",
		Global:       false,
		SubCommand:   "mfa",
		Type:         "bool",
		DefaultValue: "",
		Flag:         "set-default",
		ShortFlag:    "d",
	},
	Property{
		Name:         "aws.mfaarn",
		Description:  "AWS MFA ARN",
		Global:       false,
		SubCommand:   "mfa",
		Type:         "string",
		DefaultValue: "",
		Flag:         "mfa-arn",
		ShortFlag:    "",
	},
	Property{
		Name:         "aws.mfacode",
		Description:  "AWS MFA Code",
		Global:       false,
		SubCommand:   "mfa",
		Type:         "string",
		DefaultValue: "",
		Flag:         "mfa-code",
		ShortFlag:    "c",
	},
}

//FilterConfig filters the properties
func FilterConfig(vs []Property, f func(Property) bool) []Property {
	vsf := make([]Property, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//GlobalConfig is the global config
var GlobalConfig = FilterConfig(Config, func(sp Property) bool { return sp.Global })

// SubCommandConfig returns the sub command config
func SubCommandConfig(subcommand string) []Property {
	return FilterConfig(Config, func(sp Property) bool { return sp.SubCommand == subcommand })
}

// InitConfigFromFlags inits from the flags
func InitConfigFromFlags(cmd *cobra.Command, config []Property) {

	//fmt.Println("Load AWS Config from Flags")
	flagValues := make(map[string]interface{})

	for _, p := range config {
		switch p.Type {
		case "bool":
			var temp bool
			cmd.PersistentFlags().BoolVarP(&temp, p.Flag, p.ShortFlag, false, p.Description)
			viper.BindPFlag(p.Name, cmd.PersistentFlags().Lookup(p.Flag))
			flagValues[p.Name] = temp
		default:
			var temp string
			cmd.PersistentFlags().StringVarP(&temp, p.Flag, p.ShortFlag, p.DefaultValue, p.Description)
			viper.BindPFlag(p.Name, cmd.PersistentFlags().Lookup(p.Flag))
			flagValues[p.Name] = temp
		}
	}

}
