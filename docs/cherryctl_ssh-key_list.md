## cherryctl ssh-key list

Retrieves ssh-keys.

### Synopsis

Retrieves ssh-keys. If the project ID is specified, will return all SSH keys assigned to a specific project.

```
cherryctl ssh-key list [-p <project_id>] [flags]
```

### Examples

```
  # List of ssh-keys:
  cherryctl ssh-key list
```

### Options

```
  -h, --help             help for list
  -p, --project-id int   The project's ID.
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

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - Ssh-key operations.

