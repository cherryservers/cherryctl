## cherryctl plan list

Retrieves a list of server plans.

### Synopsis

Retrieves a list of server plans with their corresponding hourly rates and stock volumes.

```
cherryctl plan list [-t <team_id>] [--region-id <region_slug>] [--type <type>] [flags]
```

### Examples

```
  # List available plans:
  cherryctl plans list
```

### Options

```
  -h, --help            help for list
  -r, --region string   The Slug or ID of region.
  -t, --team-id int     The team's ID. Return plans prices based on team billing details.
      --type strings    Comma separated list of available plan types (baremetal,virtual,vps)
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

* [cherryctl plan](cherryctl_plan.md)	 - Plan operations.

