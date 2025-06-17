## cherryctl init

Configuration file initialization.

### Synopsis

Init will prompt for account settings and store the values as defaults in a configuration file.
This file is stored in the default user configuration directory (platform dependent), unless otherwise specified by the --config flag.
The --context flag can be used to change the default config file name.
Any Cherry CLI command line argument can be specified in the config file.
Be careful not to define options that you do not intend to use as defaults.
The configuration directory path can be changed with the CHERRY_CONFIG environment variable or --config option.

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
      --config string    Path to configuration file directory. The CHERRY_CONFIG environment variable can be used as well.
      --context string   Specify a custom context name (default "default")
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl](cherryctl.md)	 - Cherry Servers Command Line Interface (CLI)

