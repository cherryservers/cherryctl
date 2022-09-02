## cherryctl project update

Update a project.

### Synopsis

Update a project.

```
cherryctl project update ID [-p <project_id>] [--name <project_name>] [--bgp <bool>] [flags]
```

### Examples

```
  # Update project to enable BGP:
  cherryctl project update 12345 --bgp true
```

### Options

```
  -b, --bgp              True to enable BGP in a project.
  -h, --help             help for update
      --name string      Project name.
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

* [cherryctl project](cherryctl_project.md)	 - Project operations.

