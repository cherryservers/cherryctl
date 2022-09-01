## cherryctl project get

Retrieves project details.

### Synopsis

Retrieves the details of the specified project.

```
cherryctl project get ID [-p <project_id>] [flags]
```

### Examples

```
  # Gets the details of the specified project:
  cherryctl project get 12345
```

### Options

```
  -h, --help             help for get
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

