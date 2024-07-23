## cherryctl backup methods

Retrieves backup storage access methods.

### Synopsis

Retrieves a list of available backup access methods.

```
cherryctl backup methods <backup_ID> [flags]
```

### Examples

```
  # Retrieves a list of backup storage methods:
		cherryctl backup methods 123
```

### Options

```
  -h, --help   help for methods
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

* [cherryctl backup](cherryctl_backup.md)	 - Server backup operations. For more information on backups check Product Docs: https://docs.cherryservers.com/knowledge/backup-storage

