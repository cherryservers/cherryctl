## cherryctl user get

Retrieves information about the current user or a specified user.

### Synopsis

Returns either information about the current user or information about a specified user. Specified user information is only available if that user shares a project with the current user.

```
cherryctl user get ID [flags]
```

### Examples

```
  # Gets the current user's information:
		cherryctl user get
		
		# Returns information on user with ID 123:
		cherryctl user get 123
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

* [cherryctl user](cherryctl_user.md)	 - User operations.

