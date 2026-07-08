# Slug

A small Go package that converts strings to a "slugified" version. The slugified version refers here to a string that should still represent the original meaning, but remove most special characters. No consideration was given to performance, this is mostly to convert file names to ones more easily parseable/pipeable.

```sh
go get -u github.com/phasesoftware/slug@v1.1.0
```

## Usage

```go
import "github.com/phasesoftware/slug"

slug.Ify("Schwiizerdütsch & svensk omljud ä ö å behandling")
// "Schwiizerdutsch-and-svensk-omljud-a-o-a-behandling"

slug.German.Ify("Schwiizerdütsch & svensk omljud ä ö å behandling")
// "Schwiizerduetsch-and-svensk-omljud-ae-oe-a-behandling"
```

## Command line utility

The included `slugify` command line tool treats each input argument as a file system path and attempts to rename it to a slugified version.

<img src="https://github.com/phasesoftware/slug/raw/master/screenshot.png?raw=true" width="640"/>

### Install

```sh
go install -v github.com/phasesoftware/slug/cmd/slugify@latest
```

### Usage examples:

```sh
slugify "path/to/file"
slugify -de "path/to/german/file" # umlauts as digraphs, e.g. ü→ue
slugify path/with/wildcard*
ls *.mp4 | xargs -d '\n' slugify # GNU xargs
ls *.mp4 | tr \\n \\0 | xargs -0 slugify # macOS
```
