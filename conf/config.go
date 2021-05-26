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
    API *struct {
        Host  string
        Port  int
        WeCom *struct {
            AppId     string
            AppSecret string
        }
    }

    Services *struct {
        Location *struct {
            Key string
        }

        MiniPro *struct {
            AppId     string
            AppSecret string
        }
    }

    Logging *log.Config
    ORM     *orm.Config
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
