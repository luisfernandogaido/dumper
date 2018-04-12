package conf

import (
	"io/ioutil"
	"encoding/json"
)

type Conf struct {
	Token string
}

var defaultConf Conf

func Load(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &defaultConf)
}

func Token() string {
	return defaultConf.Token
}
