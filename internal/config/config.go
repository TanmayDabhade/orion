package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	AIProvider    string          `mapstructure:"ai_provider"`
	AIEndpoint    string          `mapstructure:"ai_endpoint"`
	AIModel       string          `mapstructure:"ai_model"`
	AIKey         string          `mapstructure:"ai_key"`
	SearchEngine  string          `mapstructure:"search_engine"`
	RiskThreshold string          `mapstructure:"risk_threshold"`
	Features      map[string]bool `mapstructure:"features"`
}

func Default() Config {
	return Config{
		AIProvider:    "ollama",
		AIEndpoint:    "http://localhost:11434",
		AIModel:       "llama3.1",
		SearchEngine:  "https://google.com/search?q=%s",
		RiskThreshold: "medium",
		Features:      map[string]bool{},
	}
}

func Load() (Config, error) {
	cfg := Default()

	v := viper.New()
	v.SetConfigFile(Path())
	v.SetConfigType("yaml")
	v.SetDefault("ai_provider", cfg.AIProvider)
	v.SetDefault("ai_endpoint", cfg.AIEndpoint)
	v.SetDefault("ai_model", cfg.AIModel)
	v.SetDefault("ai_key", os.Getenv("ORION_AI_KEY"))
	v.SetDefault("search_engine", cfg.SearchEngine)
	v.SetDefault("risk_threshold", cfg.RiskThreshold)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return cfg, nil
		}
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	if cfg.SearchEngine == "" {
		cfg.SearchEngine = Default().SearchEngine
	}
	if cfg.AIEndpoint == "" {
		cfg.AIEndpoint = Default().AIEndpoint
	}
	if cfg.AIModel == "" {
		cfg.AIModel = Default().AIModel
	}
	if cfg.RiskThreshold == "" {
		cfg.RiskThreshold = Default().RiskThreshold
	}

	return cfg, nil
}

func Dir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(os.TempDir(), "orion")
	}
	return filepath.Join(home, ".config", "orion")
}

func Path() string {
	return filepath.Join(Dir(), "config.yaml")
}

func ShortcutsPath() string {
	return filepath.Join(Dir(), "shortcuts.yaml")
}

func HistoryPath() string {
	return filepath.Join(Dir(), "history.db")
}
