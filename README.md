<p align="center">
  <img alt="z logo" src="docs/img/logo.png" height="140" />
  <p align="center">pipes made easy.</p>
</p>

[![test](https://github.com/serramatutu/z/actions/workflows/test.yml/badge.svg)](https://github.com/serramatutu/z/actions/workflows/test.yml)

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

## Basic usage

z makes it easy to perform lots of different types of operations on pipes.

```
# replace ":" with "-"
echo -n "replace:me" | z replace : -

# hash a file
z hash md5 < myfile.txt

# get the length of an md5 hash
echo -n "hashme" | z hash md5 _ length

# get the length of a file
z length < myfile.txt

# encode to hex
echo -n "hexme" | z encode hex

# decode from hex 
echo -n 6865786d65 | z decode hex

# get the lengths of each line in a file
z split "\n" _ length < myfile.txt
```

Learn more about z with our [usage guide](https://serramatutu.github.io/z/docs/usage/).


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

## Documentation and command reference

To learn how to use z, check out [the docs](https://serramatutu.github.io/z/docs/).

## Contributing

Check out our [contributing guidelines](https://serramatutu.github.io/z/contribute/).

## NOTICE! Z IS STILL A WORK IN PROGRESS 
z is still under development and many of its features are not implemented yet. Check out development progress [here](./TODO.md).

For this reason, z is still not stable and may change a lot in the upcoming weeks/months. **DO NOT USE IT FOR SERIOUS APPLICATIONS**
