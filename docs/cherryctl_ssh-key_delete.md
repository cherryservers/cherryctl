## cherryctl ssh-key delete

Deletes an SSH key.

### Synopsis

Deletes an SSH key with a confirmation prompt. To skip the confirmation use --force. Does not remove the SSH key from existing servers.

```
cherryctl ssh-key delete ID [-f] [flags]
```

### Examples

```
  # Deletes an SSH key, with confirmation:
  cherryctl shh-key delete 12345
  >
  ✔ Are you sure you want to delete SSH key 12345: y
  		
  # Deletes an SSH key, skipping confirmation:
  cherryctl shh-key delete 12345 -f
```

### Options

```
  -f, --force   Skip confirmation.
  -h, --help    help for delete
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

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - SSH key operations.

