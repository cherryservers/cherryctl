## cherryctl backup plans

Retrieves available backup storage plans.

### Synopsis

Retrieves the details of available backup storage plans.

```
cherryctl backup plans [flags]
```

### Examples

```
  # Gets the list of available backup storage plans:
  cherryctl backup plans
```

### Options

```
  -h, --help   help for plans
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

* [cherryctl backup](cherryctl_backup.md)	 - Server backup operations. For more information on backups, check the Product Docs: https://docs.cherryservers.com/knowledge/backup-storage

