package conf

import (
    "bytes"
    "centnet-fzmps/common/database/orm"
    "centnet-fzmps/common/log"
    "fmt"
    "github.com/BurntSushi/toml"
)

var (
    Conf = &Config{}
)

type Config struct {
    Logging  *log.Config
    MiniPro  *MiniProgram
    ORM      *orm.Config
    Services *Services
}

type Auth struct {
    AppId     string
    AppSecret string
}

type API struct {
    Host string
    Port int
}

type MiniProgram struct {
    Auth Auth
    API  API
}

type Location struct {
    Key string
}

type WeCom struct {
    AppId     string
    AppSecret string
}

type Services struct {
    Location *Location
    WeCom    *WeCom
}

func (c *Config) String() string {
    b := &bytes.Buffer{}
    err := toml.NewEncoder(b).Encode(c)
    if err != nil {
        return fmt.Errorf("toml.NewEncoder failed: %v", err).Error()
    }
    return b.String()
}

func Init() error {
    if len(configFile) == 0 {
        return fmt.Errorf("profile not found")
    }

    _, err := toml.DecodeFile(configFile, Conf)
    if err != nil {
        return fmt.Errorf("toml.DecodeFile: %v\n", err)
    }
    return nil
}
