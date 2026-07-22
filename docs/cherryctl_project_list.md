## cherryctl project list

Retrieves a list of projects details.

### Synopsis

Retrieves the details of projects.

```
cherryctl project list -t <team_id> [flags]
```

### Examples

```
  # List available projects:
  cherryctl project list -t 12345
```

### Options

```
  -h, --help          help for list
  -t, --team-id int   The team's ID.
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

* [cherryctl project](cherryctl_project.md)	 - Project operations.

