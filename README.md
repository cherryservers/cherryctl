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

Usage examples:

List objects
------------

```
$ cherryctl.darwin list teams
```

Output:

```
API Endpoint: https://api.cherryservers.com/v1/teams

-------		---------		---------------		------------	-------
Team ID		Team name		Promo remaining		Promo usage	    Pricing
-------		---------		---------------		------------	-------
28519		Team  team	    6675.05	     		608.29	    	0.5468
28990		Super team		0			0		0
-------		---------		---------------		------------	-------
```