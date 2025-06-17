## cherryctl server create

Create a server.

### Synopsis

Create a server in specified project.

```
cherryctl server create -p <project_id> --plan <plan_slug> --region <region_slug> [--hostname <hostname>] [--image <image_slug>] [--ssh-keys <ssh_key_ids>] [--ip-addresses <ip_addresses_ids>] [--os-partition-size <size>] [--userdata-file <filepath>] [--tags] [--spot-instance] [--storage-id <storage_id>] [flags]
```

### Examples

```
  # Provisions a E5-1620v4 server in the LT-Siauliai location running on Ubuntu 24.04:
  cherryctl server create -p <project_id> --plan e5_1620v4 --hostname staging-server-1 --image ubuntu_24_04 --region LT-Siauliai
```

### Options

```
  -h, --help                    help for create
      --hostname string         Server hostname.
      --image string            Operating system slug for the server.
      --ip-addresses strings    Comma separated list of IP addresses ID's to be embed in the Server.
      --os-partition-size int   OS partition size in GB.
      --plan string             Slug of the plan.
  -p, --project-id int          The project's ID.
      --region string           Slug of the region where server will be provisioned.
      --spot-instance           Provisions the server as a spot instance.
      --ssh-keys strings        Comma separated list of SSH key ID's to be embed in the Server.
      --storage-id int          ID of the storage that will be attached to server.
      --tags strings            Tag or list of tags for the server: --tags="key=value,env=prod".
      --userdata-file string    Path to a userdata file for server initialization.
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

* [cherryctl server](cherryctl_server.md)	 - Server operations. For more information on server types, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#compute

