## cherryctl storage create

Create storage.

### Synopsis

Create storage in speficied project.

```
cherryctl storage create -p <project_id> --size <gigabytes> --region <region_slug> [--description <text>] [flags]
```

### Examples

```
  # Create storage volume with 500GB of space in the LT-Siauliai location:
  cherryctl storage create -p 12345 --size 500 --region LT-Siauliai
```

### Options

```
      --description string   Storage description.
  -h, --help                 help for create
  -p, --project-id int       The project's ID.
      --region string        Slug of the region.
      --size int             Storage volume size in gigabytes.
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

* [cherryctl storage](cherryctl_storage.md)	 - Storage operations. For more information on Elastic Block Storage, check the Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/

