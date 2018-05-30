package helpers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// CreateConfig creates a config for the given configName and marshals a struct into it.
func CreateConfig(configName string, conf interface{}) {
	c, err := yaml.Marshal(conf)
	if err != nil {
		panic(err)
	}
	if ioutil.WriteFile(configName+".yml", c, 0644) != nil {
		panic(err)
	}

}

// GetConfig unmarshals a .yml file of the given configName
func GetConfig(configName string, c interface{}) {
	f, err := ioutil.ReadFile(configName + ".yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, c)
	if err != nil {
		panic(err)
	}
}
