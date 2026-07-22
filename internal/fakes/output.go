package fakes

import "github.com/cherryservers/cherryctl/internal/outputs"

type Outputer struct {
	Calls []OutputRecord
	Err   error
}

func (o *Outputer) Output(in any, th []string, td *[][]string) error {
	o.Calls = append(o.Calls, OutputRecord{in: in, th: th, td: *td})
	return o.Err
}

func (o *Outputer) SetFormat(outputs.Format) {

}
