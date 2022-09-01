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
      --config string    Path to JSON or YAML configuration file
      --fields strings   Comma separated object field names to output in result. Fields can be used for list and get actions.
  -o, --output string    Output format (*table, json, yaml)
      --token string     API Token (CHERRY_AUTH_TOKEN)
```

### SEE ALSO

* [cherryctl user](cherryctl_user.md)	 - User operations.

