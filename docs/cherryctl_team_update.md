## cherryctl team update

Update a team.

### Synopsis

Update a team.

```
cherryctl team update ID [-t <team_id>] [--name <team_name>] [--currency <currency_code>] [--type <team_type>] [flags]
```

### Examples

```
  # Update a team to change currency to EUR:
  cherryctl team update 12345 --currency EUR
```

### Options

```
      --currency string   Team currency, available otions: EUR, USD.
  -h, --help              help for update
      --name string       Team name.
  -t, --team-id int       The team's ID.
      --type string       Team type, available options: personal, business.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl team](cherryctl_team.md)	 - Team operations. For more information on Teams check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#teams

