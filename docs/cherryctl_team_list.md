## cherryctl team list

Retrieves list of teams details.

### Synopsis

Retrieves the details of teams.

```
cherryctl team list [flags]
```

### Examples

```
  # List available teams:
  cherryctl team list
```

### Options

```
  -h, --help   help for list
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl team](cherryctl_team.md)	 - Team operations. For more information on Teams, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#teams

