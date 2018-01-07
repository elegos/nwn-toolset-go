# Neverwinter Nights Toolset (written in go)

## Goal of the project
To create an alternative Toolset for Neverwinter Nights, cross-platform
(mainly addressed to Linux). After having replicated the funcionalities,
it will then have the second goal to make the user experience better, introducing
new ways of accessing resources and ease of navigation of the toolset; it will also
support, among other features, plugins and eventually a custom NWScript format
based on ECMA script 2015 (if technically possible).

## License
The software is under the MIT license: this allows everyone use the code for pretty
anything (even proprietary software). That said, I'd love to see opensource forks
and merge requests instead.

```
Copyright (c) 2017 Giacomo Furlan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Project requirements
- `go` command ([website](https://golang.org/))
- `ninja` command ([website](https://ninja-build.org/))

## How to build
On Linux and OSX, just type `ninja` in the terminal / command line. It will parse
the build.ninja file and compile the executable in the `bin` folder. If on Windows,
run `.\windows.cmd`: for more options run `.\windows.cmd /?`.

Please note that `ninja` won't compile the binary again, if it doesn't find any
changes in the main go file. For this reason, you may want to run it in this way:

    ninja -t clean; ninja [&& ./bin/cli src/aurora/file/erf/test/module.mod]

On windows:

    .\windows.cmd /a

Which stands for `/a`ll (clean and compile)

`./bin/cli -h` (`.\bin\cli.exe -h` on Windows) will show the help message.

## IDE / editor of choice
The editor of choice for this project is  [Atom](https://atom.io) with the
[go-plus](https://atom.io/packages/go-plus) plugin. The plugin alone covers
project building, testing, linting, hyperclicking.

After installing Atom, you can simply run `apm install go-plus`, or go under Settings,
Install and search for go-plus. After the installation you'll need to configure the plugin:
Settings -> Packages -> go-plus (cogwheel icon: settings) -> `GOPATH`: (where you cloned the repo).

Alternatively to setting go-plus's GOPATH variable, you can create a `.env` file, installing
the `atom-env-for-project` plugin too, and setting GOPATH in the root of the project (this will
allow you to write multiple go projects with different GOPATHs).

Other suggested packages are:

Packages:
- `project-manager`: to manage different projects
- `file-icons`: to recognise the files in the tree view by their icons, rather than their extension.
- `atom-env-for-project`: use `.env` files to setup per-project environment settings

Themes:
- `atom-material-syntax`
- `atom-material-ui`

## Git configuration
The `.gitignore` file is deliberately ignoring system or ide-wise files, like `.DS_Store`
for OSX, `thumbs.db` for Windows, `.idea` for IntelliJ and so on so forth, in order
to be system-agnostic as any `.gitignore` file should be. In order to ignore those files
user-wide, I suggest you [this reading](https://gist.github.com/subfuzion/db7f57fff2fb6998a16c).

## Contributions
You may either ask me the permission to directly write in the repository, or
create pull requests. Whatever the case, the only strict requirement I insist on
is to follow the code style. The suggested editor covers all of this for you.

As per the tests, please don't use the same package name of the one to test, but
use the following syntax: `package <package_name>_test`; this will prevent test
flooding in the package itself.

Always follow the guide lines, read the code, and continue making criticisms in order
to achieve a better, elegant, performant code.

## But I don't know go!
Go has been developed in order to be as fast as C, but way more programmer-friendly.
Amongst the various key concepts, go offers:

1. integrated testing solution
2. powerful auto-formatter
3. written for performance
4. easiness to write standardised code
5. easy to understand (if you know another programming language)
6. severly limited language, which means, more or less, one way to do _your task here_.
Obviously this also means you can easily master it's syntax in few days (literally!)

That said, I chose `go` also to _learn_ it and use it at work, so what's the best
way to having fun doing your job? :)
