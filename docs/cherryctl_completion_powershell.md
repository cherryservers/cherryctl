## cherryctl completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	cherryctl completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
cherryctl completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
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

