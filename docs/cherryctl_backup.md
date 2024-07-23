## cherryctl backup

Server backup operations. For more information on backups check Product Docs: https://docs.cherryservers.com/knowledge/backup-storage

### Synopsis

Server backup storage operations: plans, get, list, create, update, methods, update-method, method-whitelist and remove.

### Options

```
  -h, --help   help for backup
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

* [cherryctl](cherryctl.md)	 - Cherry Servers Command Line Interface (CLI)
* [cherryctl backup create](cherryctl_backup_create.md)	 - Create a backup storage.
* [cherryctl backup delete](cherryctl_backup_delete.md)	 - Delete a backup storage.
* [cherryctl backup get](cherryctl_backup_get.md)	 - Retrieves backup storage details.
* [cherryctl backup list](cherryctl_backup_list.md)	 - Retrieves a list of backup storages.
* [cherryctl backup methods](cherryctl_backup_methods.md)	 - Retrieves backup storage access methods.
* [cherryctl backup methods-whitelist](cherryctl_backup_methods-whitelist.md)	 - Retrieves a list of backup storage methods whitelist.
* [cherryctl backup plans](cherryctl_backup_plans.md)	 - Retrieves available backup storage plans.
* [cherryctl backup update](cherryctl_backup_update.md)	 - Update a backup storage.
* [cherryctl backup update-method](cherryctl_backup_update-method.md)	 - Update a backup storage access method.
* [cherryctl backup upgrade](cherryctl_backup_upgrade.md)	 - Upgrade a backup storage plan.

