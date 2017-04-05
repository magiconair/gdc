# Github Download Count

GDC is a simple Go port of the cool https://github.com/mmilidoni/github-downloads-count project.

Kudos to @mmilidoni for writing it.

## Installation

```
go get -u github.com/magiconair/gdc
```

## Usage

This script shows the downloads count of the GitHub repositories.

To avoid issues with rate limits you can set a personal access token in the `GITHUB_TOKEN` environment variable.

You must put the archive file in a GitHub release, an example is on [GitHub Blog - Release Your Software](https://github.com/blog/1547-release-your-software).

In order to get the downloads count of all GitHub repositories of a specific user you need to execute:

```
./gdc github-username
```

You can specify a single project:

```
./gdc github-username github-projectname
```

Example:

```
./gdc rethinkdb
837	elasticsearch-river-rethinkdb-1.0.1.zip
93	elasticsearch-river-rethinkdb-1.0.0.zip
930	Total downloads
```


## License

MIT licensed
