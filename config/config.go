package config

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Config struct {
    APIKey          string `yaml:"API_KEY"`
    APIService      string `yaml:"API_SERVICE"`
    Model           string `yaml:"MODEL"`
    GeminiAPIURL    string `yaml:"GEMINI_API_URL"`
    GoogleProjectID string `yaml:"GOOGLE_PROJECT_ID"`
}

var Cfg Config

func LoadConfig(configFile string) {
    data, err := ioutil.ReadFile(configFile)
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    err = yaml.Unmarshal(data, &Cfg)
    if err != nil {
        log.Fatalf("Error parsing config file: %v", err)
    }

    fmt.Printf("Loaded config: %+v\n", Cfg)
}
