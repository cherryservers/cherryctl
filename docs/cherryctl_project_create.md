## cherryctl project create

Create a project.

### Synopsis

Create a new project in a specified team.

```
cherryctl project create -t <team_id> --name <project_name> [--bgp] [--no-bgp] [flags]
```

### Examples

```
  # To create a new project with BGP enabled:
  cherryctl project create -t 12345 --name "Project with BGP" --bgp
```

### Options

```
  -b, --bgp           Enable BGP in a project.
  -h, --help          help for create
      --name string   Project name.
      --no-bgp        Disable BGP in a project. (default true)
  -t, --team-id int   The teams's ID.
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

* [cherryctl project](cherryctl_project.md)	 - Project operations.

