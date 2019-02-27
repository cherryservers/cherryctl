Cherry Cloud CLI
================

Introduction
------------

**cherryctl** is an command line tool to manage Cherry Servers services, order new servers, manage floating ips, ssh keys, etc.

* __cherryctl.darwin__: MacOS version binary of cli
* __cherryctl.linux__: Linux version binary of cli
* __cherryctl.windows__: Windows version binary of cli

Installation
------------

You may put the binary whereever you want and just launch it as an executable. Don't forget to set exec flags on a binary file:

```
$ chmod +x cherryctl.linux
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
$ cherryctl.darwin -h

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
$ cherryctl.darwin list -h

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
$ cherryctl.darwin list teams
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
$ cherryctl.darwin list projects -t 28519

API Endpoint: https://api.cherryservers.com/v1/teams/28519/projects

----------	------------		----
Project ID	Project name		Href
----------	------------		----
79811		My Project		/projects/79811
79813		For NL ACL testing	/projects/79813
----------	------------		----
```