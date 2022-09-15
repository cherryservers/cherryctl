## cherryctl completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cherryctl completion bash)

To load completions for every new session, execute once:

#### Linux:

	cherryctl completion bash > /etc/bash_completion.d/cherryctl

#### macOS:

	cherryctl completion bash > $(brew --prefix)/etc/bash_completion.d/cherryctl

You will need to start a new shell for this setup to take effect.


```
cherryctl completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl completion](cherryctl_completion.md)	 - Generate the autocompletion script for the specified shell

