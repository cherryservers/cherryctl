## cherryctl ip get

Get an IP address details.

### Synopsis

Get the details of the specified IP address.

```
cherryctl ip get ID [flags]
```

### Examples

```
  # Gets the details of the specified IP address:
  cherryctl ip get 30c15082-a06e-4c43-bfc3-252616b46eba
```

### Options

```
  -h, --help   help for get
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

* [cherryctl ip](cherryctl_ip.md)	 - IP address operations. For more information on IP addresses, check out the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#ip-addressing

