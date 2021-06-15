---
layout: default
title: Getting Started
prev_page: /
permalink: /get-started/
redirect_from: /getting-started/
---

## Installation

All z builds are checksummed then signed with GPG. You can verify the signature using the [public key](./pubkey.asc).

### Using the install script

This will download the latest z binary into `/usr/local/bin`:
```
curl -s https://raw.githubusercontent.com/serramatutu/z/main/install.sh | sudo sh -
```

You can also specify a version:
```
curl -s https://raw.githubusercontent.com/serramatutu/z/main/install.sh | sudo sh -s - v0.1.0
```

### Manual installation
1. Go to our [releases](https://github.com/serramatutu/z/releases) page.
2. Download the latest release binary.
3. Include it in your `$PATH` by adding `export PATH=$PATH:/path/to/z` to your shell's rc file.

### Building from source
1. Make sure you have [Go](https://golang.org/) installed
2. Clone this repository
3. Run `make build`. 

This will produce a z binary file inside the `bin/` folder.
