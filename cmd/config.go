package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"orion/internal/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Orion configuration",
	Long:  `Get, set, and list configuration values.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration values",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		fmt.Printf("Config file: %s\n\n", config.Path())

		v := reflect.ValueOf(cfg)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := typeOfS.Field(i)
			tag := field.Tag.Get("mapstructure")
			if tag == "" {
				continue
			}

			val := v.Field(i).Interface()

			// Mask secrets
			if strings.Contains(strings.ToLower(tag), "key") {
				strVal := fmt.Sprintf("%v", val)
				if len(strVal) > 4 {
					val = strVal[:4] + "****"
				} else if len(strVal) > 0 {
					val = "****"
				}
			}

			fmt.Printf("%-20s: %v\n", tag, val)
		}
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Example: `  orion config set ai_key my-secret-key
  orion config set ai_provider gemini`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		value := args[1]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		// Use reflection to set the field
		v := reflect.ValueOf(&cfg).Elem()
		typeOfS := v.Type()

		fieldFound := false
		for i := 0; i < v.NumField(); i++ {
			field := typeOfS.Field(i)
			tag := field.Tag.Get("mapstructure")
			if tag == key {
				f := v.Field(i)
				if f.Kind() == reflect.String {
					f.SetString(value)
					fieldFound = true
				} else {
					return fmt.Errorf("field '%s' is not a string, cannot set via CLI yet", key)
				}
				break
			}
		}

		if !fieldFound {
			return fmt.Errorf("unknown config key: %s", key)
		}

		if err := config.Save(cfg); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Printf("✅ Updated %s\n", key)
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		v := reflect.ValueOf(cfg)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := typeOfS.Field(i)
			tag := field.Tag.Get("mapstructure")
			if tag == key {
				fmt.Println(v.Field(i).Interface())
				return nil
			}
		}

		return fmt.Errorf("unknown config key: %s", key)
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	rootCmd.AddCommand(configCmd)
}
