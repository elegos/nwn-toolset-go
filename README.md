# Neverwinter Nights Toolset (written in go)

## Goal of the project
To create an alternative Toolset for Neverwinter Nights, cross-platform
(mainly addressed to Linux). After having replicated the funcionalities,
it will then have the second goal to make the user experience better, introducing
new ways of accessing resources and ease of navigation of the toolset; it will also
support, among other features, plugins and eventually a custom NWScript format
based on ECMA script 2015 (if technically possible).

## Project requirements
- `go` command ([website](https://golang.org/))
- `ninja` command ([website](https://ninja-build.org/))

## How to build
In terminal / command line, in the project's directory, simply run `ninja`. It
will parse the build.ninja file and compile the executable in the `bin` folder.

You can manually build the project running the following commands:

- Linux/OSX: `GOBIN=``pwd``/bin GOPATH=``pwd`` go install [project name]`
- Windows: `set GOBIN=%path%; set GOPATH=%path%; go install [project name]`

Where `[project name]` can be `cli` or any `src` top-level directory (note: the `aurora`
folder has no main package, thus it's not an executable)

## IDE / editor of choice
The editor of choice for this project is  [Atom](https://atom.io) with the
[go-plus](https://atom.io/packages/go-plus) plugin. The plugin alone covers
project building, testing, linting, hyperclicking.

After installing Atom, you can simply run `apm install go-plus`, or go under Settings,
Install and search for go-plus. After the installation you'll need to configure the plugin:
Settings -> Packages -> go-plus (cogwheel icon: settings) -> `GOPATH`: (where you cloned the repo).

Other suggested packages are:

Packages:
- `project-manager`: to create a per-project settings (useful to setup go-plus).
- `file-icons`: to recognise the files in the tree view by their icons, rather than their extension.

Themes:
- `atom-material-syntax`
- `atom-material-ui`

## Git configuration
The `.gitignore` file is deliberately ignoring system or ide-wise files, like `.DS_Store`
for OSX, `thumbs.db` for Windows, `.idea` for IntelliJ and so on so forth, in order
to be system-agnostic as any `.gitignore` file should be. In order to ignore those files
user-wide, I suggest you [this reading](https://gist.github.com/subfuzion/db7f57fff2fb6998a16c).
