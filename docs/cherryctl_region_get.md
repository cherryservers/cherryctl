## cherryctl region get

Retrieves region details.

### Synopsis

Retrieves the details of the specified region.

```
cherryctl region get {ID | SLUG} [flags]
```

### Examples

```
  # Gets the details of the specified region:
  cherryctl region get LT-Siauliai
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

* [cherryctl region](cherryctl_region.md)	 - Region operations. For more information on Networking, check the Product Docs: https://docs.cherryservers.com/knowledge/product-docs#network

