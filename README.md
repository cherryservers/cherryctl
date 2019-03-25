cherryctl
================

Introduction
------------

**cherryctl** is an command line tool to manage Cherry Servers services, order new servers, manage floating ips, ssh keys, etc. You may download it from below for desired operating system:

* Download [cherryctl](http://downloads.cherryservers.com/other/cherryctl/linux/cherryctl) for Linux.
* Download [cherryctl](http://downloads.cherryservers.com/other/cherryctl/mac/cherryctl) for Mac.
* Download [cherryctl](http://downloads.cherryservers.com/other/cherryctl/windows/cherryctl) for Windows.

Installation
------------

You may put the binary whereever you want and just launch it as an executable. Don't forget to set exec flags on a binary file:

```
$ chmod +x cherryctl
```

Requirements
------------

In order to use this module you will need to export Cherry Servers API token. You can generate and get it from your Client Portal. The easiest way is to export variable like this:

```
$ export CHERRY_AUTH_TOKEN="2b00042f7481c7b056c4b410d28f33cf"
```

Usage
-----

You may get help with __-h__ key passed to the program:

```
$ cherryctl -h

Usage:
  cherryctl [command]

Available Commands:
  add         Ads various objects
  help        Help about any command
  list        Lists various objects
  power       Manages power on servers
  remove      Removes various objects
  update      Updates various objects

Flags:
  -h, --help   help for cherry-cloud-cli

Use "cherryctl [command] --help" for more information about a command.
```

```
$ cherryctl list -h

Usage:
  cherryctl list [command]

Available Commands:
  images       List images
  ip-address   List specific ip address
  ip-addresses List ip addresses
  plans        List plans
  projects     List projects
  server       List specific server
  servers      List servers
  ssh-key      List ssh key
  ssh-keys     List ssh keys
  teams        List teams

Flags:
  -h, --help   help for list

Use "cherryctl list [command] --help" for more information about a command.
```

Examples:
---------

List objects
------------

List teams:
```
$ cherryctl list teams
```

```
API Endpoint: https://api.cherryservers.com/v1/teams

-------		---------		---------------		------------	-------
Team ID		Team name		Promo remaining		Promo usage	    Pricing
-------		---------		---------------		------------	-------
28519		Team  team	    6675.05	     		608.29	    	0.5468
28990		Super team		0			0		0
-------		---------		---------------		------------	-------
```

List projects:

```
$ cherryctl list projects -t 28519

API Endpoint: https://api.cherryservers.com/v1/teams/28519/projects

----------	------------		----
Project ID	Project name		Href
----------	------------		----
79811		My  Project		      /projects/79811
79813		For NL ACL testing	/projects/79813
----------	------------		----
```

List plans:
```
$ cherryctl list plans -t 28519

API Endpoint: https://api.cherryservers.com/v1/teams/28519/plans

-------		-------------	----------	-------------	---		-------		---
Plan ID		Plan name	    Plan price	CPU		        RAM		Regions		Qty
-------		-------------	----------	-------------	---		-------		---
59		    Smart8		    0.1089		  E3-1240		    8		  EU-East-1	34
93		    SSD Smart8	  0.1089		  X5650		      8		  EU-East-1	4
94		    SSD Smart16	  0.121		    X5650		      16		EU-East-1	26
92		    Smart16		    0.121		    E3-1240		    16		EU-East-1	36
90		    Quad Smart	  0.1331		  56xx		      16		EU-East-1	30
165		    Hexa Smart	  0.1573		  E5645		      24		EU-East-1	11
86		    E3-1240v3	    0.242		    E3-1240v3	    16		EU-East-1	43
113		    E3-1240v5	    0.2662		  E3-1240v5	    32		EU-East-1	47
109		    E5-1660v3	    0.3993		  E5-1660v3	    64		EU-East-1	1
126		    2x E5-2620v4	0.6897		  2x E5-2620v4	128		EU-East-1	11
96		    2x E5-2620v2	0.7139		  2x E5-2620v2	128		EU-East-1	1
144		    2x E5-2630Lv4	0.7381		  2x E5-2630Lv4	128		EU-East-1	8
107		    2x E5-2630v3	0.7986		  2x E5-2630v3	128		EU-East-1	5
98		    2x E5-2650v2	0.8712		  2x E5-2650v2	128		EU-East-1	1
122		    2x E5-2650v4	0.9317		  2x E5-2650v4	128		EU-East-1	2
147		    2x E5-2697v2	1.0285		  2x E5-2697v2	384		EU-East-1	13
117		    2x E5-2670v3	1.0527		  2x E5-2670v3	128		EU-East-1	4
108		    2x E5-2680v4	1.1132		  2x E5-2680v4	128		EU-East-1	1
-------		-------------	----------	-------------	---		-------		---
```

Add objects
-----------

### Add Project

Flags:
```
Flags:
  -h, --help                  help for project
  -p, --project-name string   Provide project-name
  -t, --team-id int           Provide team-id
```

Add project:
```
$ cherryctl add project \
  -p "Superb Project" \
  -t 28519
```

### Add SSH keys

Flags:
```
Flags:
  -h, --help               help for ssh-key
  -l, --key-label string   Provide ssh key label (default "ssh-key-label")
  -f, --key-path string    Provide path to ssh key
  -k, --raw-key string     Provide ssh raw key
```

Add key
```
$ cherryctl add ssh-key \
    -l linas \
    -f $KEY_PATH/linas.key
```

### Add servers
Flags:
```
Flags:
  -h, --help                   help for server
  -s, --hostname string        Provide hostname (default "server-name.examples.com")
  -i, --image string           Provide image
  -d, --ip-addresses strings   Provide image
  -l, --plan-id string         Provide plan-id
  -p, --project-id string      Provide project-id
  -g, --region string          Provide region (default "EU-East-1")
  -k, --ssh-keys strings       Provide ssh-keys
```

Add server
```
$ cherryctl add server \
    -s server02.aleja.lt \
    -i 'Ubuntu 16.04 64bit' \
    -l 161 \
    -p 79813 \
    -k 95 \
    -g EU-East-1
```

### Add floating IP addresses

Update objects
--------------

### Update project
Flags:
```
Flags:
  -h, --help                  help for project
  -i, --project-id string     Provide project-id
  -p, --project-name string   Provide project-name
```

Update project name:
```
$ cherryctl update project \
  -p "Superb Project Prod" \
  -i 83445
```

### Update SSH keys
Flags:
```
Flags:
  -h, --help               help for ssh-key
  -k, --key-id string      Provide ssh key id for update
  -l, --key-label string   Provide new label for key
```

Update SSH key label
```
$ cherryctl ssh-key \
    -k 95 \
    -l marius
```

### Update Floating IP addresses
Flags:

```
Flags:
  -a, --a-record string              Provide a-record (default "a-record.example.com")
  -i, --floating-id string           Provide floating ip id for update
  -f, --floating-ip string           Provide floating ip for update
  -h, --help                         help for ip-address
  -p, --project-id string            Provide project-id
  -r, --ptr-record string            Provide ptr-record (default "ptr-record.examples.com")
  -t, --routed-to string             Provide ipaddress_id to route to
  -n, --routed-to-hostname string    Provide hostname of the server to route to
  -d, --routed-to-server-id string   Provide id of the server to route to
  -s, --routed-to-server-ip string   Provide primary ip of the server to route to
```

Route floating IP address (by floating IP ID) to server`s hostname
```
$ cherryctl update ip-address \
    -p 79813 \
    -i 6a17b7ef-5617-2e85-e34e-986bb80fe3a9  \
    -a bla3.testas.lt. \
    -r bla3.testas.lt. \
    -n server02.aleja.lt
```

Route floating IP address (by floating IP ID) to server`s ID
```
$ cherryctl update ip-address \
    -p 79813 \
    -i 6a17b7ef-5617-2e85-e34e-986bb80fe3a9  \
    -a bla3.testas.lt. \
    -r bla3.testas.lt. \
    -d 175821
```

Route floating IP address (by floating IP ID) to server`s IP
```
$ cherryctl update ip-address \
    -p 79813 \
    -i 6a17b7ef-5617-2e85-e34e-986bb80fe3a9  \
    -a bla3.testas.lt. \
    -r bla3.testas.lt. \
    -s 188.214.132.41
```

Route floating IP address to server`s IP
```
$ cherryctl update ip-address \
    -p 79813
    -f 5.199.171.104 \
    -a bla3.testas.lt. \
    -r bla3.testas.lt. \
    -s 188.214.132.41
```


Remove objects
--------------

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations.