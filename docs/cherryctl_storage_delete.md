## cherryctl storage delete

Delete a storage.

### Synopsis

Deletes the specified storage with a confirmation prompt. To skip the confirmation use --force.

```
cherryctl storage delete ID [f] [flags]
```

### Examples

```
  # Deletes the specified storage:
  cherryctl storage delete 12345
  >
  ✔ Are you sure you want to delete storage 12345: y
  		
  # Deletes a storage, skipping confirmation:
  cherryctl storage delete 12345 -f
```

### Options

```
  -f, --force   Skips confirmation for the storage deletion.
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

* [cherryctl storage](cherryctl_storage.md)	 - Storage operations. For more information on Elastic Block Storage, check the Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/

