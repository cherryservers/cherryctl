## cherryctl ssh-key delete

Deletes an SSH key.

### Synopsis

Deletes an SSH key with a confirmation prompt. To skip the confirmation use --force. Does not remove the SSH key from existing servers.

```
cherryctl ssh-key delete -i <ssh_key_id> [flags]
```

### Examples

```
  # Deletes an SSH key, with confirmation:
  cherryctl shh-key delete -i 12345
  >
  ✔ Are you sure you want to delete SSH key 12345: y
  		
  # Deletes a server, skipping confirmation:
  cherryctl shh-key delete -f -i 12345
```

### Options

```
  -f, --force            Skips confirmation for the SSH key deletion.
  -h, --help             help for delete
  -i, --ssh-key-id int   ID of the SSH key.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl ssh-key](cherryctl_ssh-key.md)	 - Ssh-key operations.

