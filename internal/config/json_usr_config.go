package config

import (
    "os"
    "encoding/json"
    "fmt"
)

var configFile string = ".gatorconfig.json"

type Config struct {
    DbURL           string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
    filePath, err := getConfigFilePath()
    if err != nil {
        return Config{}, err
    }
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        return Config{}, err
    }
    
    var userConfig Config
    err = json.Unmarshal(data, &userConfig)
    if err != nil {
        return Config{}, err
    }
    return userConfig, nil
}

func getConfigFilePath() (string, error) {
    homePath, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%s/%s", homePath, configFile), nil
}

func SetUser(userName string) error {
    config, err := Read()
    if err != nil {
        return err
    }
    config.CurrentUserName = userName
    err = write(config)
    if err != nil {
        return err
    }
    return nil
}

func write(config Config) error {
    filePath, err := getConfigFilePath()
    if err != nil {
        return nil
    }
    
    data, err := json.Marshal(config)
    if err != nil {
        return err
    }

    err = os.WriteFile(filePath, data, 0666)
    if err != nil {
        return nil
    }

    return nil
}
