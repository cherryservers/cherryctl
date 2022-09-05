package init

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"text/tabwriter"

	"github.com/cherryservers/cherrygo/v3"
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
	Token     string `yaml:"token,omitempty"`
	ProjectID int    `yaml:"project-id,omitempty"`
	TeamID    int    `yaml:"team-id,omitempty"`
}

func (c *Client) NewCommand() *cobra.Command {
	// initCmd represents a command that, when run, generates a
	// set of initironment variables, for use in shell initironments
	// v := c.tokener.Config()
	// projectId := v.GetString("project-id")
	initCmd := &cobra.Command{
		Use:   `init`,
		Short: "Create a configuration file.",
		Long:  "Init will prompt for account settings and store the values as defaults in a configuration file that may be shared with other Cherry Servers tools. This file is typically stored in $HOME/.config/cherry/config.yaml. Any Cherry CLI command line argument can be specified in the config file. Be careful not to define options that you do not intend to use as defaults. The configuration file written to can be changed with CHERRY_CONFIG and --config.",

		Example: `  # Example config:
  --
  token: foo
  team-id: 123
  project-id: 123`,

		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			config, _ := cmd.Flags().GetString("config")
			if config == "" {
				config = c.Servicer.DefaultConfig(true)
			}

			fmt.Print("Cherry Servers API Tokens can be obtained through the portal at http://portal.cherryservers.com/.\n\n")
			fmt.Print("Token (hidden): ")
			b, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			fmt.Println()
			token := string(b)
			c.Servicer.SetToken(token)
			cherryclient := c.Servicer.API(cmd)
			c.UserService = cherryclient.Users
			_, _, err = c.UserService.CurrentUser(nil)
			if err != nil {
				return errors.Wrap(err, "Invalid authentication token")
			}

			teams, _, err := cherryclient.Teams.List(&cherrygo.GetOptions{})
			if err != nil {
				return errors.Wrap(err, "Failed to get team list")
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

			teamID, _ := strconv.Atoi(userTeam)
			projects, _, err := cherryclient.Projects.List(teamID, &cherrygo.GetOptions{})
			if err != nil {
				return errors.Wrap(err, "Failed to get project list")
			}
			fmt.Printf("Choose Project ID from list below:")
			tw = tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
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

			b, err = formatConfig(userProj, userTeam, token)
			if err != nil {
				return err
			}
			return writeConfig(config, b)
		},
	}

	return initCmd
}

func formatConfig(userProj, userTeam, token string) ([]byte, error) {
	userProjInt, _ := strconv.Atoi(userProj)
	userTeamInt, _ := strconv.Atoi(userTeam)
	f := configFormat{ProjectID: userProjInt, TeamID: userTeamInt, Token: token}
	b, err := yaml.Marshal(f)

	if err != nil {
		return nil, err
	}
	b = append([]byte("---\n"), b...)
	return b, err
}

func writeConfig(config string, b []byte) error {
	fmt.Fprintf(os.Stderr, "\nWriting %s\n", config)
	dir := filepath.Dir(config)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("could not make directory %q: %s", dir, err)
	}
	return ioutil.WriteFile(config, b, 0o600)
}

type Servicer interface {
	API(*cobra.Command) *cherrygo.Client
	SetToken(string)
	DefaultConfig(bool) string
}
