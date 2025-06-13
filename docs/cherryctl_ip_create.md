## cherryctl ip create

Create floating IP address.

### Synopsis

Create floating IP address in specified project.

```
cherryctl ip create -p <project_id> --region <region_slug> [--target-hostname <hostname> | --target-id <server_id> | --target-ip-id <ip_uuid>] [--ptr-record <ptr>] [--a-record <a>] [--tags <tags>] [flags]
```

### Examples

```
  # Create a floating IP address in the LT-Siauliai location:
  cherryctl ip create -p <project_id> --region LT-Siauliai
```

### Options

```
      --a-record string          Relative DNS name for the IP address. Resulting FQDN will be '<relative-dns-name>.cloud.cherryservers.net' and must be globally unique.
  -h, --help                     help for create
  -p, --project-id int           The project's ID.
      --ptr-record string        Reverse DNS name for the IP address.
      --region string            Slug of the region from where IP address will requested.
      --tags strings             Tag or list of tags for the server: --tags="key=value,env=prod".
      --target-hostname string   The hostname of the server to assign the created IP to.
      --target-id int            The ID of the server to assign the created IP to.
      --target-ip-id string      Subnet or primary-ip type IP ID to route the created IP to.
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

* [cherryctl ip](cherryctl_ip.md)	 - IP address operations. For more information on IP addresses, check out the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#ip-addressing

