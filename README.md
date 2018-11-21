# cloner

cloner is a simple tool to clone all of the github repositories in specified organization, inspired by https://gist.github.com/tagomoris/1394916845a1b8020e43

## Use case

When you join to new team/company which uses github, you can clone all of repositories easily.

## Installation

Just go get

```
go get golang.org/x/oauth2
go get github.com/google/go-github/github
go get github.com/hiroakis/cloner
```

or build

```
git clone git@github.com:hiroakis/cloner.git
cd cloner
make
```

## Usage

* basic usage

All of the repository will be cloned to the current directory. Note, the tool requires git command, so you have to install git to your machine.

```
cloner -token YOUR_GITHUB_API_TOKEN -org TARGET_ORGANIZATION_NAME
```

* options

```
-token  The github access token. The token should have "repo" scope. REQUIRED.
-org    The organization name. REQUIRED.
-type   The type of the repository. "private" or "public". OPTIONAL(default: private)
-page   The page num. OPTIONAL(default: 1)
-per    The number of results to include per page. OPTIONAL(default: 100)
```

## License

MIT
# cloner
