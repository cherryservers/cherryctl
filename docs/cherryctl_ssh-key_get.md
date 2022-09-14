## cherryctl ssh-key get

Retrieves ssh-key details.

### Synopsis

Retrieves the details of the specified ssh-key.

```
cherryctl ssh-key get ID [flags]
```

### Examples

```
  # Gets the details of the specified ssh-key:
  cherryctl ssh-key get 12345
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - Ssh-key operations.

