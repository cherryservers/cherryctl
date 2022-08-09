## cherryctl docs

Generate command markdown documentation.

### Synopsis

Generates command markdown documentation in the specified directory. Each command gets a markdown file.

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
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl](cherryctl.md)	 - Command line interface for Cherry Servers

