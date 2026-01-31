package lint

import "github.com/spf13/viper"

type Config struct {
	Strict  bool
	Quiet   bool
	Format  string
	NoVoice bool
}

func LoadConfig() Config {
	return Config{
		Strict:  viper.GetBool("strict"),
		Quiet:   viper.GetBool("quiet"),
		Format:  viper.GetString("format"),
		NoVoice: viper.GetBool("no-voice"),
	}
}
