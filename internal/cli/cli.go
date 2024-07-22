package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	envPrefix                = "CHERRY"
	DefaultContext           = "default"
	DefaultConfigDirName     = "cherryctl"
	DeprecatedDefaultContext = "cherry"
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

	rootCmd.PersistentFlags().StringVar(&c.configPath, "config", "", "Path to JSON or YAML configuration file")
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
			v.AutomaticEnv()

			replacer := strings.NewReplacer("-", "_", ".", "_")
			v.SetEnvKeyReplacer(replacer)

			if c.configPath == "" {
				var err error
				c.configPath, err = defaultConfigPath()
				if err != nil {
					log.Fatalln(err)
				}
			}

			// Backward compatability (cherry was renamed to default)
			if c.context == DefaultContext {
				if _, err := os.Stat(filepath.Join(c.configPath, DefaultContext)); err != nil {
					if _, err := os.Stat(filepath.Join(c.configPath, DeprecatedDefaultContext)); err != nil {
						log.Fatalln(fmt.Errorf("Couldn't find configuration file. To initiate run `cherryctl init` command"))
					} else {
						c.context = DeprecatedDefaultContext
					}
				}
			}

			v.SetConfigName(c.context)
			v.AddConfigPath(c.configPath)

			if err := v.ReadInConfig(); err != nil {
				log.Fatalln(fmt.Errorf("Could not read config: %s", err))
			}

			v.SetEnvPrefix(envPrefix)
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
		// 	_ = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		// }

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func defaultConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return configDir, err
	}

	return path.Join(configDir, DefaultConfigDirName), nil
}

// TODO deprecate
func (c *Client) ConfigFilePath(withExtension bool) string {
	config := path.Join(c.configPath, c.context)
	if withExtension {
		config = config + ".yaml"
	}
	return config
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
