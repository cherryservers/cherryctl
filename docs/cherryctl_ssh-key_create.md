## cherryctl ssh-key create

Adds an SSH key for the current user's account.

### Synopsis

Adds an SSH key for the current user's account.

```
cherryctl ssh-key create --key <public_key> --label <label> [flags]
```

### Examples

```
  # Adds a key labled "example-key" to the current user account.
  cherryctl ssh-key create --key ssh-rsa AAAAB3N...user@domain.com --label example-key
```

### Options

```
  -h, --help           help for create
      --key string     Public SSH key string.
      --label string   Label of the SSH key.
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

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - Ssh-key operations.

