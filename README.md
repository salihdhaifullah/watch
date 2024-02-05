# File Watcher Multiplexer
## Overview
The File Watcher Multiplexer is a program designed to monitor file system changes in the current root directory and execute specified commands based on those changes. It takes its configuration from a JSON file named `watch.config.json` in the root of your project.

## `watch.config.json` Example
```json
{
    "ignore_regex": [".*_templ\\.go$"],
    "ignore": ["static/fonts", "vendor", ".git", ".bin", "main", "static/index.css"],
    "delay": 1500,
    "setup_commands": ["templ generate"],
    "parallel_commands": ["gotip build -mod=vendor -o main -p 4", "npx tailwindcss -i ./tailwind.css -o ./static/index.css"],
    "background_commands": ["./main"]
}
```

## Configuration Options
- **ignore_regex**: An array of regular expressions. Files matching any of these regular expressions will be ignored.

- **ignore**: An array of file or directory names to be ignored.

- **delay**: The delay in milliseconds between file system events to avoid triggering commands too frequently.

- **setup_commands**: An array of setup commands. These commands act as prerequisites for the subsequent commands.

- **parallel_commands**: An array of commands to be executed in parallel. These commands can run independently, and the execution of other programs does not depend on them.

- **background_commands**: An array of commands that run in the background. These commands typically include long-running processes like a web server that does not terminate by itself.

## Usage
Create a `watch.config.json` file in the root of your project with the desired configuration.

#### build the program by.

```bash
git clone https://github.com/salihdhaifullah/watch.git
cd watch
go build .
```
and then you can add it to your path or use it by relative path ag (../../watch/watch)

## Important Notes
- The program watches for file changes in the current root directory.
- It ignores specified files and directories based on the provided configuration.
- Commands specified in the configuration are executed based on their dependencies and parallelism.
- The delay between file system events can be adjusted to control the frequency of command execution.
- Ensure that the required tools and dependencies for the specified commands are installed on your system.
