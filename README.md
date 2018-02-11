# Cross
Cross-platform functions I use in some utilities.

## Functions

### SetBasePath()

This is called by init(). It builds the base directory, which in the case of Linux, BSD etc. is just the user's home directory, and on macOS it's `~/Library/Application Support/`.

### ConfigName(filename)

Returns a path suitable for command-line programs, combining the config path and the configuration file.
- On Linux/*BSD this will be the absolute path of `~/.<program>/<filename>`, and the program name will have spaces removed
- On macOS it returns the absolute path of `~/Library/Application Support/<program>/<filename>`, with any spaces intact

### ServerConfigName(program, filename)

- Linux/*BSD returns `/etc/<program>/<filename>`
- macOS simply returns the output from `GetConfigName()`

### Exists(path)

Checks for the existence of a path and return true or false.

### DirExists(path)

Checks for the existence of a directory. A file at the same path would return false.

### FileExists(path)

Checks for the existence of a file. A directory at the same path would return false.
