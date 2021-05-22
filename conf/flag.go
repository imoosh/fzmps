package conf

import "flag"

var (
    configFile string
)

func init() {
    flag.StringVar(&configFile, "c", "./config.toml", "default profile")
}