## cherryctl storage attach

Attach storage volume to a specified server.

### Synopsis

Attach storage volume to a specified server.

```
cherryctl storage attach ID {--server-id <id> | --server-hostname <hostname>} [-p <project_id>] [flags]
```

### Examples

```
  # Attach storage to specified server:
  cherryctl storage attach 12345 --server-id 12345
```

### Options

```
  -h, --help                     help for attach
  -p, --project-id int           The project's ID.
      --server-hostname string   The Hostname of a server.
  -s, --server-id int            The server's ID.
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

* [cherryctl storage](cherryctl_storage.md)	 - Storage operations. For more information on Elastic Block Storage check Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/

