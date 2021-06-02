# Developing z
Here are some instructions before you start contributing to z!

## Git hooks setup
Before you make any changes, make sure to set up your Git Hooks. Do this by running `make githooks`


## Editor debugging setup
As z depends heavily on STDIN and it's delve's [design decision](https://github.com/go-delve/delve/issues/1274) not to support integrated STDIN debugging, we have developed a workaround.
1. Run `sh scripts/debug.sh <z-args>` in a separate terminal.
2. Configure your editor to attach to delve on `127.0.0.1:2345`
3. Write inputs to the separate terminal

Example: `sh scripts/debug.sh hash md5 _ length`

## Setting up `act`
To run Github Actions locally, make sure you have [act](https://github.com/nektos/act) installed.

This is useful for testing Actions and simulating them before submitting a Pull Request
