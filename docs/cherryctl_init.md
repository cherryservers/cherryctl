## cherryctl init

Create a configuration file.

### Synopsis

Init will prompt for account settings and store the values as defaults in a configuration file that may be shared with other Cherry Servers tools. This file is typically stored in $HOME/.config/cherry/config.yaml. Any Cherry CLI command line argument can be specified in the config file. Be careful not to define options that you do not intend to use as defaults. The configuration file written to can be changed with CHERRY_CONFIG and --config.

```
cherryctl init
```

### Examples

```
  # Example config:
  --
  token: foo
  team-id: 123
  project-id: 123
```

### Options

```
  -h, --help   help for init
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

* [cherryctl](cherryctl.md)	 - Command line interface for Cherry Servers

