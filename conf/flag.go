package conf

import "flag"

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "./conf/config.toml", "default profile")
}
