## cherryctl team delete

Delete a team.

### Synopsis

Deletes the specified team with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl team delete ID -t <team_id> [flags]
```

### Examples

```
  # Deletes the specified team:
  cherryctl team delete 12345
  >
  ✔ Are you sure you want to delete team 12345: y
  		
  # Deletes a team, skipping confirmation:
  cherryctl team delete -f -t 12345
```

### Options

```
  -f, --force         Skips confirmation for the tean deletion.
  -h, --help          help for delete
  -t, --team-id int   The ID of a team.
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

* [cherryctl team](cherryctl_team.md)	 - Team operations. For more information on Teams check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#teams

