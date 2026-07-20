package init

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"

	"github.com/cherryservers/cherryctl/internal/cli"

	"github.com/cherryservers/cherrygo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"gopkg.in/yaml.v2"
)

type Client struct {
	Servicer    Servicer
	UserService cherrygo.UsersService
}

func NewClient(s Servicer) *Client {
	return &Client{
		Servicer: s,
	}
}

type configFormat struct {
	APIKey    string `yaml:"api-key,omitempty"`
	ProjectID int    `yaml:"project-id,omitempty"`
	TeamID    int    `yaml:"team-id,omitempty"`
}

func (c *Client) NewCommand() *cobra.Command {
	// initCmd represents a command that, when run, generates a
	// set of initironment variables, for use in shell initironments
	initCmd := &cobra.Command{
		Use:   `init`,
		Short: "Configuration file initialization.",
		Long: `Init will prompt for account settings and store the values as defaults in a configuration file.
This file is stored in the default user configuration directory (platform dependent), unless otherwise specified by the --config flag.
The --context flag can be used to change the default config file name.
Any Cherry CLI command line argument can be specified in the config file.
Be careful not to define options that you do not intend to use as defaults.
The configuration directory path can be changed with the CHERRY_CONFIG environment variable or --config option.`,

		Example: `  # Example config:
  --
  api-key: foo
  team-id: 123
  project-id: 123`,

		DisableFlagsInUseLine: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				return err
			}

			// If config path is empty, check env var and if that is not set, use OS default.
			if configPath == "" {
				configPath = os.Getenv(cli.EnvPrefix + "_CONFIG")

				if configPath == "" {
					configDir, err := os.UserConfigDir()
					if err != nil {
						return err
					}
					configPath = filepath.Join(configDir, cli.DefaultConfigDirName)
				}

			}

			err = c.checkAndCreateConfigDir(configPath)
			if err != nil {
				return err
			}

			apiKey, err := c.readAPIKey()
			if err != nil {
				return err
			}
			if err = c.validateAPIKey(cmd); err != nil {
				return err
			}

			cherryClient := c.Servicer.API(cmd)

			userTeam, err := c.readUserTeam(cmd.Context(), cherryClient)
			if err != nil {
				return err
			}

			userProj, err := c.readUserProject(cmd.Context(), cherryClient, userTeam)
			if err != nil {
				return err
			}

			b, err := formatConfig(userProj, userTeam, apiKey)
			if err != nil {
				return err
			}

			context, _ := cmd.Flags().GetString("context")
			if !strings.HasSuffix(context, ".yaml") {
				context += ".yaml"
			}

			return writeConfig(configPath, context, b)
		},
	}

	return initCmd
}

func (c *Client) readAPIKey() (string, error) {
	fmt.Print("Cherry Servers API Keys can be obtained through the portal at https://portal.cherryservers.com/settings/api-keys.\n\n")
	fmt.Print("API key (hidden): ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return string(b), err
	}
	fmt.Println()
	apiKey := string(b)
	c.Servicer.SetAPIKey(apiKey)

	return apiKey, nil
}

func (c *Client) validateAPIKey(cmd *cobra.Command) error {
	cherryClient := c.Servicer.API(cmd)
	c.UserService = cherryClient.Users
	_, _, err := c.UserService.CurrentUser(cmd.Context(), nil)
	if err != nil {
		return errors.Wrap(err, "Invalid API key")
	}
	return nil
}

func (c *Client) readUserTeam(ctx context.Context, cherryClient *cherrygo.Client) (string, error) {
	teams, _, err := cherryClient.Teams.List(ctx, &cherrygo.GetOptions{})
	if err != nil {
		return "", errors.Wrap(err, "Failed to get team list")
	}
	fmt.Printf("Choose Team ID from list below:")
	tw := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(tw, "\n-------\t---------\t---------------\n")
	fmt.Fprintf(tw, "Team ID\tTeam name\tRemaining credit\n")
	fmt.Fprintf(tw, "-------\t---------\t---------------\n")

	for _, t := range teams {
		fmt.Fprintf(tw, "%v\t%v\t%v\n",
			t.ID, t.Name, (t.Credit.Account.Remaining + t.Credit.Promo.Remaining))
	}
	tw.Flush()
	fmt.Printf("\n")
	fmt.Printf("Team ID: ")
	userTeam := ""
	fmt.Scanln(&userTeam)

	return userTeam, nil
}

func (c *Client) readUserProject(ctx context.Context, cherryClient *cherrygo.Client, userTeam string) (string, error) {
	teamID, _ := strconv.Atoi(userTeam)
	projects, _, err := cherryClient.Projects.List(ctx, teamID, &cherrygo.GetOptions{})
	if err != nil {
		return "", errors.Wrap(err, "Failed to get project list")
	}
	fmt.Printf("Choose Project ID from list below:")
	tw := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
	fmt.Fprintf(tw, "\n-------\t---------\n")
	fmt.Fprintf(tw, "Project ID\tProject name\n")
	fmt.Fprintf(tw, "-------\t---------\n")

	for _, p := range projects {
		fmt.Fprintf(tw, "%v\t%v\n",
			p.ID, p.Name)
	}
	tw.Flush()
	fmt.Printf("\n")
	fmt.Printf("Project ID: ")
	userProj := ""
	fmt.Scanln(&userProj)
	return userProj, nil
}

func formatConfig(userProj, userTeam, apiKey string) ([]byte, error) {
	userProjInt, _ := strconv.Atoi(userProj)
	userTeamInt, _ := strconv.Atoi(userTeam)
	f := configFormat{ProjectID: userProjInt, TeamID: userTeamInt, APIKey: apiKey}
	b, err := yaml.Marshal(f)

	if err != nil {
		return nil, err
	}
	b = append([]byte("---\n"), b...)
	return b, err
}

func writeConfig(configPath string, context string, b []byte) error {
	configDest := filepath.Join(configPath, context)
	fmt.Fprintf(os.Stderr, "\nWriting configuration to: %s\n", configDest)
	return os.WriteFile(configDest, b, 0o600)
}

func (c *Client) checkAndCreateConfigDir(configDir string) error {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0o700); err != nil {
			return fmt.Errorf("could not create directory %q: %s", configDir, err)
		}
	}
	return nil
}

type Servicer interface {
	API(*cobra.Command) *cherrygo.Client
	SetAPIKey(string)
}

var _ Servicer = (*cli.Client)(nil)
