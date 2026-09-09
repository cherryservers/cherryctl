package outputs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
	"gopkg.in/yaml.v2"
)

type Format string

const (
	FormatTable Format = "table"
	FormatJSON  Format = "json"
	FormatYAML  Format = "yaml"
)

type Outputer interface {
	Output(interface{}, []string, *[][]string) error
	SetFormat(Format)
}

type Standard struct {
	Format Format
}

func outputJSON(in interface{}) error {
	output, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func outputYAML(in interface{}) error {
	output, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func (o *Standard) Output(in interface{}, header []string, data *[][]string) error {
	if o.Format == FormatJSON {
		return outputJSON(in)
	} else if o.Format == FormatYAML {
		return outputYAML(in)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table = table.Configure(func(cfg *tablewriter.Config) {
			cfg.Row.Alignment.Global = tw.AlignLeft
		})
		table.Header(header)
		if err := table.Bulk(*data); err != nil {
			return err
		}
		return table.Render()
	}
}

func (o *Standard) SetFormat(fmt Format) {
	o.Format = fmt
}
