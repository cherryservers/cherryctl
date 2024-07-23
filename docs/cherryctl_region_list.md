## cherryctl region list

Retrieves list of regions.

### Synopsis

Retrieves list of regions.

```
cherryctl region list [flags]
```

### Examples

```
  # Gets list of regions:
  cherryctl region list
```

### Options

```
  -h, --help   help for list
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

* [cherryctl region](cherryctl_region.md)	 - Region operations. For more information on Networking check Product Docs: https://docs.cherryservers.com/knowledge/product-docs#network

