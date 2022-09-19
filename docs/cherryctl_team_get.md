## cherryctl team get

Retrieves team details.

### Synopsis

Retrieves the details of the specified team.

```
cherryctl team get ID [flags]
```

### Examples

```
  # Gets the details of the specified team:
  cherryctl team get 12345
```

### Options

```
  -h, --help          help for get
  -i, --team-id int   The team's ID.
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

