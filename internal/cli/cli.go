package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	EnvPrefix            = "CHERRY"
	DefaultContext       = "default"
	DefaultConfigDirName = "cherryctl"
	OldDefaultContext    = "cherry"
	OldConfigPathSuffix  = "/.config/cherry"
)

type Client struct {
	apiClient *cherrygo.Client

	fields       *[]string
	queryParams  map[string]string
	configPath   string
	context      string
	outputFormat string
	cherryToken  string
	apiURL       string
	Version      string
	rootCmd      *cobra.Command
	viper        *viper.Viper
}

func (c *Client) SetToken(token string) {
	c.cherryToken = token
}

func NewClient(cherryToken, apiURL, Version string) *Client {
	return &Client{
		cherryToken: cherryToken,
		apiURL:      apiURL,
		Version:     Version,
	}
}

func (c *Client) API(cmd *cobra.Command) *cherrygo.Client {
	if c.cherryToken == "" {
		log.Fatal("Cherry Servers API authentication token not provided. Please set the 'CHERRY_AUTH_TOKEN' environment variable or create a configuration file using 'cherryctl init'.")
	}

	if c.apiClient == nil {
		args := []cherrygo.ClientOpt{cherrygo.WithAuthToken(c.cherryToken), cherrygo.WithUserAgent("cherry-cli/" + c.Version)}

		if c.apiURL != "" {
			args = append(args, cherrygo.WithURL(c.apiURL))
		}

		client, err := cherrygo.NewClient(args...)
		if err != nil {
			return nil
		}

		c.apiClient = client
	}

	return c.apiClient
}

func (c *Client) Format() outputs.Format {
	format := outputs.FormatTable

	switch f := outputs.Format(c.outputFormat); f {
	case "":
		break
	case outputs.FormatTable,
		outputs.FormatJSON,
		outputs.FormatYAML:
		format = f
	default:
		log.Printf("unknown format: %q. Using default.", f)
	}
	return format
}

func (c *Client) NewCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:               "cherryctl",
		Short:             "Cherry Servers Command Line Interface (CLI)",
		Long:              `cherryctl is a command line interface (CLI) for Cherry Servers API`,
		DisableAutoGenTag: true,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			c.Config(cmd)
		},
	}
	rootCmd.PersistentFlags().String("token", "", "API Token (CHERRY_AUTH_TOKEN)")
	rootCmd.PersistentFlags().String("auth-token", "", "API Token (Alias)")
	authtoken := rootCmd.PersistentFlags().Lookup("auth-token")
	authtoken.Hidden = true

	rootCmd.PersistentFlags().StringVar(&c.configPath, "config", "", "Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.")
	rootCmd.PersistentFlags().StringVar(&c.context, "context", DefaultContext, "Specify a custom context name")
	rootCmd.PersistentFlags().StringVar(&c.apiURL, "api-url", c.apiURL, "Override default API endpoint")
	rootCmd.PersistentFlags().StringVarP(&c.outputFormat, "output", "o", "", "Output format (*table, json, yaml)")
	c.fields = rootCmd.PersistentFlags().StringSlice("fields", nil, "Comma separated object field names to output in result. Fields can be used for list and get actions.")

	rootCmd.Version = c.Version
	c.rootCmd = rootCmd
	return c.rootCmd
}

func (c *Client) Config(cmd *cobra.Command) *viper.Viper {
	if cmd.Name() != "init" {
		if c.viper == nil {
			v := viper.New()
			v.SetEnvPrefix(EnvPrefix)
			v.AutomaticEnv()

			replacer := strings.NewReplacer("-", "_", ".", "_")
			v.SetEnvKeyReplacer(replacer)

			v.SetConfigName(c.context)

			// Viper looks for config files in the specified paths in the order they were added.
			// This can be leveraged to prioritize the most directly defined paths.
			// Look for config in the --config flag specified path first.
			v.AddConfigPath(c.configPath)

			// Look for config in the env variable specified path.
			v.AddConfigPath(os.Getenv(EnvPrefix + "_CONFIG"))

			// If no config is found in the user specified path, look in the standard default path.
			defaultConfigPath, err := getDefaultConfigPath()
			if err != nil {
				log.Fatalln(err)
			}
			v.AddConfigPath(defaultConfigPath)

			// If no config is found in the standard default path, check the old default path.
			defaultConfigPath = filepath.Join(userHomeDir(), OldConfigPathSuffix)
			v.AddConfigPath(defaultConfigPath)

			if err = v.ReadInConfig(); err != nil {
				// For backward compatability (the default context was renamed from `cherry` to `default`).
				if c.context == DefaultContext {
					v.SetConfigName(OldDefaultContext)
					if err = v.ReadInConfig(); err != nil {
						log.Fatalln(fmt.Errorf("could not read config: %s. Initiate new configuration with `cherryctl init`", err))
					}
				} else {
					log.Fatalln(fmt.Errorf("could not read config: %s. Initiate new configuration with `cherryctl init`", err))
				}

			}

			c.viper = v
			bindFlags(cmd, v)
		}

		flagToken := cmd.Flag("token").Value.String()
		envToken := cmd.Flag("auth-token").Value.String()
		c.cherryToken = flagToken
		if envToken != "" {
			c.cherryToken = envToken
		}

		return c.viper
	}
	return nil
}

// Credit to https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		// if strings.Contains(f.Name, "-") {
		// 	envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
		// 	_ = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", EnvPrefix, envVarSuffix))
		// }

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func getDefaultConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return configDir, err
	}

	return path.Join(configDir, DefaultConfigDirName), nil
}

// Deprecated.
// Used only for checking the old default config directory.
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func (c *Client) GetOptions() *cherrygo.GetOptions {
	getOptions := &cherrygo.GetOptions{}

	if c.rootCmd.Flags().Changed("fields") {
		getOptions.Fields = *c.fields
	}

	if c.rootCmd.Flags().Changed("query-params") {
		getOptions.QueryParams = *&c.queryParams
	}

	return getOptions
}
