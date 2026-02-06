package shortcuts

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

func Load(path string) (map[string]string, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return map[string]string{}, nil
		}
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		return nil, err
	}

	m := v.GetStringMapString("shortcuts")
	if m == nil {
		m = map[string]string{}
	}
	return m, nil
}

func Save(path string, shortcuts map[string]string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.Set("shortcuts", shortcuts)
	return v.WriteConfigAs(path)
}

func Normalize(input string) string {
	return strings.ToLower(strings.Join(strings.Fields(input), " "))
}

func Resolve(shortcuts map[string]string, input string) (string, bool) {
	normalized := Normalize(input)
	for key, cmd := range shortcuts {
		if Normalize(key) == normalized {
			return cmd, true
		}
	}
	return "", false
}

func SortedKeys(shortcuts map[string]string) []string {
	keys := make([]string, 0, len(shortcuts))
	for key := range shortcuts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
