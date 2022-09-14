## cherryctl init

Configuration file initialization.

### Synopsis

Init will prompt for account settings and store the values as defaults in a configuration file that may be shared with other Cherry Servers tools. This file is typically stored in $HOME/.config/cherry/config.yaml. Any Cherry CLI command line argument can be specified in the config file. Be careful not to define options that you do not intend to use as defaults. The configuration file path can be changed with the CHERRY_CONFIG environment variable or --config option.

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
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl](cherryctl.md)	 - Cherry Servers Command Line Interface (CLI)

