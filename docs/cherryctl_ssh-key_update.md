## cherryctl ssh-key update

Updates an SSH key.

### Synopsis

Updates an SSH key with either a new public key, a new label, or both.

```
cherryctl ssh-key update -i <ssh_key_id> [--label] [--key <public_key>] [flags]
```

### Examples

```
  # Update team to change currency to EUR:
  cherryctl ssh-key update -i 12345 --key AAAAB3N...user@domain.com
```

### Options

```
  -h, --help             help for update
      --key string       Public SSH key string.
      --label string     Label of the SSH key.
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

