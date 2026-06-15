## cherryctl docs

Generate a local version of the CLI documentation.

### Synopsis

Generates a local version of the CLI documentation in the specified directory. Each command gets a markdown file.

```
cherryctl docs <destination>
```

### Examples

```
  # Generate documentation in the ./docs directory:
  cherryctl docs ./docs
```

### Options

```
  -h, --help   help for docs
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

* [cherryctl](cherryctl.md)	 - Cherry Servers Command Line Interface (CLI)

