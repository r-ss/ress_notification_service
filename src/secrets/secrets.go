package secrets

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type YAMLFile struct {
    Secrets Secrets `yaml:"secrets"`
}
type Secrets struct {
    Test string `yaml:"SECRETTEST"`
    Token string `yaml:"TELEGRAM_BOT_TOKEN"`
}

func ReadSecrets() (*Secrets, error) {

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path + "/secrets.yaml")

    secrets := &YAMLFile{}
    cfgFile, err := ioutil.ReadFile(path + "/secrets.yaml")
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(cfgFile, secrets)
	// fmt.Println(secrets.Secrets.Test)
    return &secrets.Secrets, err
}


// secrets, err := readSecrets()
// if err != nil {
// 	panic(err)
// }