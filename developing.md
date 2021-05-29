# Developing z

As z depends heavily on STDIN and it's delve's [design decision](https://github.com/go-delve/delve/issues/1274) to not support integrated STDIN debugging, we have developed a workaround.
1. Run `sh scripts/debug.sh <z-args>` in a separate terminal.
2. Configure your editor to attach to delve on `127.0.0.1:2345`
3. Write inputs to the separate terminal

Example: `sh scripts/debug.sh hash md5 _ length`
