## cherryctl backup update-method

Update a backup storage access method.

### Synopsis

Enable or disable the selected backup access method or set a list of available IP addresses allowed to use this method.

```
cherryctl backup update-method <backup_ID> --method-name <string> [--enable] [--disable] [--whitelist <ip_addresses>] [flags]
```

### Examples

```
  # Enable FTP protocol for your backup storage:
  cherryctl backup update-method 12345 --method-name FTP --enable
```

### Options

```
  -d, --disable              Disable method.
  -e, --enable               Enable method.
  -h, --help                 help for update-method
  -n, --method-name string   Backup access method name.
      --whitelist strings    A comma separated list of IP addresses to be whitelisted for access via this backup method. (default [0])
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

