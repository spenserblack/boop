# `boop`

`touch` is officially used to update timestamps for files. But it's often used to create
empty files. `boop` was inspired by *that* usage, and adds a bit more functionality for
creating new files. It is not intended to completely replace `touch`, but help reduce
series of commands that *include* `touch` into just one command.

- Create both files and *directories.* If the filepath ends in a trailing slash, `boop`
  will create a directory.
- Set execute permissions when creating a file (does nothing on Windows).
- Create any parent directories that the filepath needs.
