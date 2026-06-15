## cherryctl project create

Create a project.

### Synopsis

Create a new project for the specified team.

```
cherryctl project create -t <team_id> --name <project_name> [--bgp <bool>] [flags]
```

### Examples

```
  # To create a new project with BGP support enabled:
  cherryctl project create -t 12345 --name "Project with BGP" --bgp
```

### Options

```
  -b, --bgp           Enable BGP support.
  -h, --help          help for create
      --name string   Project name.
  -t, --team-id int   The teams's ID.
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

