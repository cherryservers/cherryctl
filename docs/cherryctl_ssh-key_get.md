## cherryctl ssh-key get

Retrieves SSH key.

### Synopsis

Retrieves the specified SSH key.

```
cherryctl ssh-key get ID [flags]
```

### Examples

```
  # Get SSH key:
  cherryctl ssh-key get 12345
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --api-key string   API key. Can be created at https://portal.cherryservers.com/settings/api-keys.
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
```

### SEE ALSO

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - SSH key operations.

