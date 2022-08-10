## cherryctl storage detach

Detach storage volume from a server.

### Synopsis

Detach storage volume from a server.

```
cherryctl storage detach -i <storage_id> [flags]
```

### Examples

```
  # Detach storage:
  cherryctl storage detach -i 12345
```

### Options

```
  -h, --help             help for detach
  -i, --storage-id int   The storage's ID.
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

