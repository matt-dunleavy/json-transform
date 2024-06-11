// config/config.go
package config

import (
    "errors"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

var Cfg Config

type AIService struct {
    Name   string `yaml:"name"`
    APIKey string `yaml:"apiKey"`
    Model  string `yaml:"model"`
    APIURL string `yaml:"apiURL"`
}

type Config struct {
    AIServices      []AIService `yaml:"aiServices"`
    GoogleProjectID string      `yaml:"googleProjectID"`
}

func LoadConfig(filePath string) error {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }

    // Debug log to show file content
    log.Printf("Config file content:\n%s\n", data)

    err = yaml.Unmarshal(data, &Cfg)
    if err != nil {
        return err
    }

    // Debug log to show parsed configuration
    log.Printf("Parsed config: %+v\n", Cfg)
    for _, service := range Cfg.AIServices {
        log.Printf("Service: %+v\n", service)
    }
    log.Printf("Google Project ID: %s\n", Cfg.GoogleProjectID)

    return nil
}

func GetAIServiceByName(name string) (*AIService, error) {
    for _, service := range Cfg.AIServices {
        if service.Name == name {
            return &service, nil
        }
    }
    return nil, errors.New("AI service not found")
}
