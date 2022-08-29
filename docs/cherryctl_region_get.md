## cherryctl region get

Retrieves region details.

### Synopsis

Retrieves the details of the specified region.

```
cherryctl region get [-i <region_slug>] [flags]
```

### Examples

```
  # Gets the details of the specified region:
  cherryctl region get -i eu_nord_1
```

### Options

```
  -h, --help               help for get
  -i, --region-id string   The Slug or ID of region.
```

### Options inherited from parent commands

```
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl region](cherryctl_region.md)	 - Region operations.
