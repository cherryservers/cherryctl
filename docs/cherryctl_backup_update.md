## cherryctl backup update

Update a backup storage.

### Synopsis

Update the backup user password or SSH key. Passwords are used in the FTP and SMB protocols, while SSH keys are used in BORG.

```
cherryctl backup update <backup_ID> [--password <plain_text>] [--ssh-key <plain_ssh_key>] [flags]
```

### Examples

```
  # Update backup storage password and SSH key:
  cherryctl backup update 12345 --password somePassword --ssh-key  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6ec8eT..."
  
  # Update backup storage user password:
  cherryctl backup update 12345 --password somePassword
```

### Options

```
  -h, --help              help for update
      --password string   Backup storage user password.
      --ssh-key string    Plain SSH key which will be stored in backup service.
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

