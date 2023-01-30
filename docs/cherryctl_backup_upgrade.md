## cherryctl backup upgrade

Upgrade a backup storage plan.

### Synopsis

Upgrade a backup storage plan to increase it's storage size. ATTENTION! Upgrade can be done once per backup plan.

```
cherryctl backup upgrade <backup_ID> --plan <backup_plan_slug> [flags]
```

### Examples

```
  # Upgrade backup storage size to 1000 gigabytes:
  cherryctl backup upgrade 12345 --plan backup_1000
```

### Options

```
  -h, --help          help for upgrade
      --plan string   Backup storage plan slug.
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl backup](cherryctl_backup.md)	 - Server backup operations. For more information on backups check Product Docs: https://docs.cherryservers.com/knowledge/backup-storage

