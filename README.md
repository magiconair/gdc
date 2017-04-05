# Github Download Count

GDC is a simple Go port of the cool https://github.com/mmilidoni/github-downloads-count project.

Kudos to @mmilidoni for writing it.

## Installation

```
go get -u github.com/magiconair/gdc
```

## Usage

This script shows the downloads count of the GitHub repositories.

You must put the archive file in a GitHub release, an example is on GitHub Blog - Release Your Software.

In order to get the downloads count of all GitHub repositories of a specific user, you need to execute:

```
gdc github-username
```

You can specify a single project:

```
gdc github-username github-projectname
```

Example:

```
gdc rethinkdb
```

output is:

```
rethinkdb-1.7.0.dmg: 454 downloads
rethinkdb-1.7.0.i686.rpm: 449 downloads
rethinkdb-1.7.0.x86_64.rpm: 450 downloads
rethinkdb-1.6.1.dmg: 455 downloads
```

## License

MIT licensed
