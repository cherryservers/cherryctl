## cherryctl completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(cherryctl completion zsh)

To load completions for every new session, execute once:

#### Linux:

	cherryctl completion zsh > "${fpath[1]}/_cherryctl"

#### macOS:

	cherryctl completion zsh > $(brew --prefix)/share/zsh/site-functions/_cherryctl

You will need to start a new shell for this setup to take effect.


```
cherryctl completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl completion](cherryctl_completion.md)	 - Generate the autocompletion script for the specified shell

