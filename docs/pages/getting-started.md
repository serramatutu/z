---
layout: default
title: Getting Started
prev_page: /
permalink: /get-started/
redirect_from: /getting-started/
---

## Installation
z is pretty easy to install:
1. Go to our [releases](https://github.com/serramatutu/z/releases) page.
2. Download the latest release binary.
3. Include it in your `$PATH` by adding `export PATH=$PATH:/path/to/z` to your shell's rc file.

That's it! You're all set!

You can also compile z from source by cloning this repository and running `make build`. This will produce a z binary file inside the `bin/` folder. If you're doing this, make sure you have [Go](https://golang.org/) installed.

_* All builds are checksummed then signed with GPG. You can verify the signature using the [public key](https://github.com/serramatutu/z/blob/main/pubkey.asc)_
