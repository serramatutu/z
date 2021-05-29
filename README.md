
# z
streams made easy

## What is z?
z is a Linux stream processor written in Go that aims to be easy and intuitive to work with.

## Why z?
Whenever I worked with linux streams (especially string manipulation), I noticed I kept having to read extensive documentation about many different programs with unintuitive names and interfaces. Sometimes I also found myself reading through Stack Overflow posts that explain how to perform the most simple of tasks.

Here are just some examples about how convoluted some of these are:

```
# splitting a string by ":"
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

_* Want to uninstall? Just delete the binary and remove it from your `$PATH`._

## Using z
z was invented to be intuitive. Here are some usage examples:
```
# splitting a string by ":"
echo -n "split:me" | z replace : \n

# hashing 
echo -n "hashme" | z hash md5
echo -n "hashme" | z hash sha1
echo -n "hashme" | z hash sha256

# encode to hex
echo -n "hexme" | z encode hex

# decode from hex
echo -n "hexme" | z decode hex

# get length of string
echo -n "lengthme" | z length
```

Need to pipe multiple z's? There's a shorter way of doing it:
```
# get the length of an md5 hash
echo -n "hashme" | z hash md5 _ length
```

## Design principles
z was designed with the following principles in mind
1. **SIMPLE INTERFACE**. All z commands must have obvious names and perform clear, well-defined operations. Any user should be able to understand what their command chain does without referring to any documentation.
2. **EASY INSTALLATION**. All z releases must export a single lightweight binary. Installing it should be as simple as downloading the binary and including it in the `$PATH`. Want to uninstall? Just delete it.

## Command reference

Check out our [help files](./help/).

If you already have z installed, avoid referring to this repo by running `z help`.
