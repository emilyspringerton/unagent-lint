package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"unagent-lint/lint"
)

var rootCmd = &cobra.Command{
	Use:   "unagent-lint [path]",
	Short: "Lint content for UNAGENT accessibility and Emily Voice compliance",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := lint.LoadConfig()
		result := lint.Run(args[0], cfg)

		lint.Print(result, cfg)

		if result.Failures > 0 {
			os.Exit(1)
		}
		if result.Warnings > 0 {
			os.Exit(2)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().Bool("strict", false, "Treat WARN as FAIL")
	rootCmd.Flags().Bool("quiet", false, "Suppress PASS output")
	rootCmd.Flags().String("format", "text", "Output format (text|json)")
	rootCmd.Flags().Bool("no-voice", false, "Disable Emily Voice rules")

	_ = viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {
	viper.SetConfigName("defaults")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}
