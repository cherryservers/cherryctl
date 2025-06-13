## cherryctl backup delete

Delete a backup storage.

### Synopsis

Deletes the specified backup storage with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl backup delete ID [-f] [flags]
```

### Examples

```
  # Deletes the specified backup storage:
  cherryctl backup delete 12345
  >
  ✔ Are you sure you want to delete backup storage 12345: y
  		
  # Deletes a storage, skipping confirmation:
  cherryctl backup delete 12345 -f
```

### Options

```
  -f, --force   Skips confirmation for the backup storage deletion.
  -h, --help    help for delete
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

