## cherryctl storage attach

Attach storage volume to a specified server.

### Synopsis

Attach storage volume to a specified server.

```
cherryctl storage attach -i <storage_id> {--server-id | --server-hostname} [-p <project_id>] [flags]
```

### Examples

```
  # Attach storage to specified server:
  cherryctl storage attach -i 12345 -s 12345
```

### Options

```
  -h, --help                     help for attach
  -p, --project-id int           The project's ID.
      --server-hostname string   The Hostname of a server.
  -s, --server-id int            The server's ID.
  -i, --storage-id int           The storage's ID.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl storage](cherryctl_storage.md)	 - Storage operations. For more information visit https://docs.cherryservers.com/knowledge/elastic-block-storage/.

