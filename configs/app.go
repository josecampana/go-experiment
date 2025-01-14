package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"io"

	"golang.org/x/mod/modfile"
)

type ProvidersConfig struct {
	Range string `json:"range"`
}

type Config struct {
	AppName   string          `json:"appName"`
	Port      string          `json:"port"`
	BasePath  string          `json:"basePath"`
	LogLevel  string          `json:"logLevel"`
	Timeout   time.Duration   `json:"timeout"`
	Providers ProvidersConfig `json:"providers"`
}

// singleton management
var (
	instance *Config
	once     sync.Once
)

func getAppName() (string, error) {
	workDir, err := os.Getwd()
	path := filepath.Join(workDir, "../../", "go.mod")
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return "", err
	}

	return modFile.Module.Mod.Path, nil
}

func loadFromFile(filename string, config *Config) error {
	workDir, err := os.Getwd()

	path := filepath.Join(workDir, "../../", filename)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(config)
}

func combineConfigs(baseConfig, envConfig *Config) {
	baseVal := reflect.ValueOf(baseConfig).Elem()
	envVal := reflect.ValueOf(envConfig).Elem()

	for i := 0; i < baseVal.NumField(); i++ {
		baseField := baseVal.Field(i)
		envField := envVal.Field(i)

		if !envField.IsZero() {
			baseField.Set(envField)
		}
	}
}

func createConfig() *Config {
	env := os.Getenv("ENV")
	if env == "" || env == "development" {
		env = "default"
	}

	// Cargar la configuración base (default)
	baseConfig := &Config{}
	err := loadFromFile("env/config.default.json", baseConfig)
	if err != nil {
		return &Config{}
	}

	if env != "default" {
		// Cargar la configuración específica del entorno
		envConfig := &Config{}
		err := loadFromFile(fmt.Sprintf("env/config.%s.json", env), envConfig)

		if err != nil {
			panic(fmt.Sprintf("error loading config file for ENV '%s' => %v", env, err))
		}

		combineConfigs(baseConfig, envConfig)

		// Override values from environment variable
		port := os.Getenv("PORT")
		logLevel := os.Getenv("LOG_LEVEL")

		if port != "" {
			baseConfig.Port = port
		}
		if logLevel != "" {
			baseConfig.LogLevel = logLevel
		}
	}

	appName, err := getAppName()
	if err == nil {
		baseConfig.AppName = path.Base(appName)
	}

	return baseConfig
}

func Get() *Config {
	once.Do(func() {
		instance = createConfig()
	})

	return instance
}
