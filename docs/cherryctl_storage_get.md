## cherryctl storage get

Retrieves storage details.

### Synopsis

Retrieves the details of the specified storage.

```
cherryctl storage get ID [flags]
```

### Examples

```
  # Gets the details of the specified storage:
  cherryctl storage get 12345
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --api-url string   Override default API endpoint (default "https://api.cherryservers.com/v1/")
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl storage](cherryctl_storage.md)	 - Storage operations. For more information on Elastic Block Storage check Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/

