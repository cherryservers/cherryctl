## cherryctl backup create

Create a backup storage.

### Synopsis

Create a backup storage for specified server.

```
cherryctl backup create {--server-id <id> | --server-hostname <hostname>} --plan <backup_plan_slug> --region <region_slug> [-p <project_id>] [--ssh-key <plain_ssh_key>] [flags]
```

### Examples

```
  # Create backup storage with 100GB space in EU-Nord-1 location for server with hostname "delicate-zebra":
  cherryctl backup create --server-hostname delicate-zebra --plan backup_100 --region eu_nord_1 --project-id 123
```

### Options

```
  -h, --help                     help for create
      --plan string              Backup storage plan slug.
  -p, --project-id int           The project's ID.
      --region string            Slug of the region.
      --server-hostname string   The Hostname of a server.
  -s, --server-id int            The server's ID.
      --ssh-key string           Plain SSH key will be stored in backup service.
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

