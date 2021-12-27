package config

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
    Server struct {
        Port string `yaml:"port"`
    } `yaml:"server"`
    Endpoints struct {
        CreateEmployee string `yaml:"new_employee"`
    } `yaml:"endpoints"`
    Database struct {
        User string `yaml:"user"`
        Host string `yaml:"host"`
        Port string `yaml:"port"`
        Db string `yaml:"db"`
    } `yaml:"database"`
    Flags struct {
        Is_StdOut string `yaml:"is_stdout"`
    } `yaml:"flags"`
}

func (c *Config) GetConfig(configFile string) *Config {

    yamlFile, err := ioutil.ReadFile(configFile)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}