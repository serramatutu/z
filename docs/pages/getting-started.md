---
layout: default
title: Getting Started
permalink: /get-started/
---

Welcome to z!

## What is z?
z is a pipe processor written in [Go](https://golang.org/) that aims to be easy and intuitive to work with.

## Why z?
Whenever I worked with unix streams/pipes, I noticed I kept having to read extensive documentation about many different programs with unintuitive names and interfaces. Sometimes I also found myself reading through Stack Overflow posts that explain how to perform the most simple of tasks.

Here are just some examples about how convoluted some of these are:

```
# replacing ":" by "\n"
echo -n "split:me" | sed 's/:/\n/g'
echo -n "split:me" | tr ':' $'\n'

# hashing 
echo -n "hashme" | md5sum
echo -n "hashme" | sha1sum
echo -n "hashme" | sha256sum

# encode to hex
echo -n "hexme" | od -A n -t x1

# decode from hex 
echo -n 6865786d65 | xxd -r -p

# get length of string
echo -n "lengthme" | wc -c
```

I asked myself: what if we had an intuitive and consistent way of doing all that?

That's how z was born.

## Installation
z is pretty easy to install:
1. Go to our [releases](https://github.com/serramatutu/z/releases) page.
2. Download the latest release binary.
3. Include it in your `$PATH` by adding `export PATH=$PATH:/path/to/z` to your shell's rc file.

That's it! You're all set!

You can also compile z from source by cloning this repository and running `make build`. This will produce a z binary file inside the `bin/` folder. If you're doing this, make sure you have [Go](https://golang.org/) installed.

_* All builds are checksummed then signed with GPG. You can verify the signature using the [public key](https://github.com/serramatutu/z/blob/main/pubkey.asc)_
