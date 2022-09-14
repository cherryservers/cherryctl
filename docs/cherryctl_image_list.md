## cherryctl image list

Retrieves list of images available for the given plan.

### Synopsis

Retrieves list of images available for the given plan.

```
cherryctl image list --plan <plan_slug> [flags]
```

### Examples

```
  # Lists the operating system images available for E5-1620v4 plan :
  cherryctl images list --plan e5_1620v4
```

### Options

```
  -h, --help          help for list
      --plan string   The Slug or ID of a plan.
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

* [cherryctl image](cherryctl_image.md)	 - Image operations.

