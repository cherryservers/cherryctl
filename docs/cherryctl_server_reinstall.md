## cherryctl server reinstall

Reinstall a server.

### Synopsis

Reinstall the specified server.

```
cherryctl server reinstall -i <server_id> --hostname --image <image_slug> --password <password> [--ssh-keys <ssh_key_ids>] [--os-partition-size <size>] [flags]
```

### Examples

```
  # Reinstall the specified server:
  cherryctl server reinstall -i 12345 -h staging-server-1 --image ubuntu_20_04 --password G1h2e_39Q9oT
```

### Options

```
  -h, --help                    help for reinstall
      --hostname string         Hostname.
      --image string            Operating system slug for the server.
      --os-partition-size int   OS partition size in GB.
      --password string         Server password.
  -i, --server-id int           The ID of a server.
      --ssh-keys strings        Comma separated list of SSH key IDs to be embed in the Server.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on provisioning on Cherry Servers, visit https://docs.cherryservers.com/knowledge/product-docs.

