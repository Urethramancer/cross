# Cross
Cross-platform functions I use in some utilities.

## Functions

### GetConfigName(program, filename)
Returns a path suitable for command-line programs, combining the name of the program and the configuration file.
Typically a .<program> directory on Unix-likes, except on macOS, where it returns a folder from ~/Library/Application Support.

### GetServerConfigName(program, filename)
Linux and other Unix-likes: Returns something like /etc/program/filename.
macOS: Simply returns output from GetConfigName().

### Exists(path)
Check for the existence of a path and return true or false.
