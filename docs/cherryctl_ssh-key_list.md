## cherryctl ssh-key list

Retrieves project members ssh-keys details.

### Synopsis

Retrieves project members ssh-keys details.

```
cherryctl ssh-key list [-p <project_id>] [flags]
```

### Examples

```
  # List of project ssh-keys:
  cherryctl ssh-key list -i 12345
```

### Options

```
  -h, --help             help for list
  -p, --project-id int   The project's ID.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - Ssh-key operations.

