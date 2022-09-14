## cherryctl server list

Retrieves server list.

### Synopsis

Retrieves a list of servers in the project.

```
cherryctl server list [-p <project_id>] [flags]
```

### Examples

```
  # Gets a list of servers in the specified project:
		cherryctl server list -p 12345
```

### Options

```
  -h, --help             help for list
  -p, --project-id int   The project's ID.
      --search string    Search server by Hostname or Public IP phrase.
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on provisioning on Cherry Servers, visit https://docs.cherryservers.com/knowledge/product-docs.

