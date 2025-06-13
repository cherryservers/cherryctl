## cherryctl ip update

Update IP address.

### Synopsis

Update tags, ptr record, a record or target server of an IP address.

```
cherryctl ip update ID [--ptr-record <ptr>] [--a-record <a>] [--tags <tags>] [flags]
```

### Examples

```
  # Updates a record and tags:
  cherryctl ip update 30c15082-a06e-4c43-bfc3-252616b46eba --a-record stage --tags="env=stage"
```

### Options

```
      --a-record string     Relative DNS name for the IP address. Resulting FQDN will be '<relative-dns-name>.cloud.cherryservers.net' and must be globally unique.
  -h, --help                help for update
      --ptr-record string   Reverse DNS name for the IP address.
      --tags strings        Tag or list of tags for the server: --tags="key=value,env=prod".
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

