package fakes

type Prompter struct {
	OK          bool
	Err         error
	GotMessages []string
}

func (p *Prompter) PromptConfirmation(msg string) (bool, error) {
	p.GotMessages = append(p.GotMessages, msg)
	return p.OK, p.Err
}
