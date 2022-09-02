## cherryctl server delete

Delete a server.

### Synopsis

Deletes the specified server with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl server delete ID [-f] [flags]
```

### Examples

```
  # Deletes the specified server:
  cherryctl server delete 12345
  >
  ✔ Are you sure you want to delete server 12345: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl server delete 12345 -f
```

### Options

```
  -f, --force   Skips confirmation for the server deletion.
  -h, --help    help for delete
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on provisioning on Cherry Servers, visit https://docs.cherryservers.com/knowledge/product-docs.

