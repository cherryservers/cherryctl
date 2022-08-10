## cherryctl project delete

Delete a project.

### Synopsis

Deletes the specified project with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl project delete -p <project_id> [flags]
```

### Examples

```
  # Deletes the specified project:
  cherryctl project delete -p 12345
  >
  ✔ Are you sure you want to delete project 12345: y
  		
  # Deletes a project, skipping confirmation:
  cherryctl project delete -f -p 12345
```

### Options

```
  -f, --force            Skips confirmation for the project deletion.
  -h, --help             help for delete
  -p, --project-id int   The ID of a project.
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

