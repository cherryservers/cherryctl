## cherryctl storage detach

Detach storage volume from a server.

### Synopsis

Detach storage volume from a server.

```
cherryctl storage detach ID [flags]
```

### Examples

```
  # Detach storage:
  cherryctl storage detach 12345
```

### Options

```
  -h, --help   help for detach
```

### Options inherited from parent commands

```
      --api-key string   API key. Can be created at https://portal.cherryservers.com/settings/api-keys.
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
```

### SEE ALSO

* [cherryctl storage](cherryctl_storage.md)	 - Storage operations. For more information on Elastic Block Storage, check the Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/

