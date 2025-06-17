## cherryctl server update

Update server.

### Synopsis

Update server.

```
cherryctl server update ID [--name <server_name>] [--hostname <hostname>] [--tags <tags>] [--bgp] [flags]
```

### Examples

```
  # Update server to change tags:
  cherryctl server update 12345 --tags="env=stage"
```

### Options

```
  -b, --bgp               True to enable BGP in a server.
  -h, --help              help for update
      --hostname string   Server hostname.
      --name string       Server name.
      --tags strings      Tag or list of tags for the server: --tags="key=value,env=prod".
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

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on server types, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#compute

