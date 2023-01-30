## cherryctl project delete

Delete a project.

### Synopsis

Deletes the specified project with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl project delete ID [-f] [flags]
```

### Examples

```
  # Deletes the specified project:
  cherryctl project delete 12345
  >
  ✔ Are you sure you want to delete project 12345: y
  		
  # Deletes a project, skipping confirmation:
  cherryctl project delete 12345 -f
```

### Options

```
  -f, --force            Skips confirmation for the project deletion.
  -h, --help             help for delete
  -p, --project-id int   The ID of a project.
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

