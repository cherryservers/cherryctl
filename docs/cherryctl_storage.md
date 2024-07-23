## cherryctl storage

Storage operations. For more information on Elastic Block Storage check Product Docs: https://docs.cherryservers.com/knowledge/elastic-block-storage/

### Synopsis

Storage operations: create, get, list, delete, attach and detach.

### Options

```
  -h, --help   help for storage
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
* [cherryctl storage attach](cherryctl_storage_attach.md)	 - Attach storage volume to a specified server.
* [cherryctl storage create](cherryctl_storage_create.md)	 - Create storage.
* [cherryctl storage delete](cherryctl_storage_delete.md)	 - Delete a storage.
* [cherryctl storage detach](cherryctl_storage_detach.md)	 - Detach storage volume from a server.
* [cherryctl storage get](cherryctl_storage_get.md)	 - Retrieves storage details.
* [cherryctl storage list](cherryctl_storage_list.md)	 - Retrieves storage list.
* [cherryctl storage update](cherryctl_storage_update.md)	 - Update storage volume.

