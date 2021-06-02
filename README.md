# z
[![test](https://github.com/serramatutu/z/actions/workflows/test.yml/badge.svg)](https://github.com/serramatutu/z/actions/workflows/test.yml)

pipes and streams made easy

## What is z?
z is a unix stream/pipe processor written in [Go](https://golang.org/) that aims to be easy and intuitive to work with.

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

_* All builds are checksummed then signed with GPG. You can verify the signature using the [public key](./pubkey.asc)_

## Usage
z is built to have an easy and intuitive interface without giving up on functionality.

### Basic usage
z commands are executed on input given through the standard input stream (aka STDIN), and their outputs are written to the standard output stream (aka STDOUT)
```
# prints the length of STDIN contents
z length
```

For this reason, you can stream file contents or pipe program outputs into z:
```
# hashes the contents of "hashme.txt" and prints it
z hash md5 < hashme.txt

# writes "bye world!" to "byeworld.txt"
echo -n "hello world!" | z replace hello bye > byeworld.txt
```

### Command chaining
It is possible to chain z commands without having to use OS pipes:
```
# with pipes
echo -n "hashme" | z hash md5 | z length

# with z command chaining
echo -n "hashme" | z hash md5 _ length
```

Both approaches have the exact same behavior. However, there are some advantages to chaining commands instead of piping:
1. `split` and `join` commands only work with command chaining
2. Information stays inside z without having to travel to the OS and back. This should only make a difference for very large files or performance-sensitive applications though.
3. Only having to type `_` instead of `| z`

### Splits and joins
z's input is always interpreted as a byte array or a string. However, there may be the need for splitting it into an array of strings and mapping operations onto the array elements. The z way of approaching this is via `split` and `join`:
1. `split` the inputs by a delimiter
2. map every split element using normal z commands such as `length`, `hash` or `replace`
3. implicitly concatenate them back into a string or `join` them with a delimiter

Here are some examples:
```
# getting the length of every line in infile.txt and writing that to outfile.txt's lines
# (split's default delimiter is "\n")
z split _ length _ join \n < infile.txt > outfile.txt

# print the md5 hashes of "a", "b" and "c", separated by ","
echo -n "a:b:c" | z split : _ hash md5 _ join ,

# print the implicitly concatenated lengths of "one", "two" and "three"
echo -n "one,two,three" | z split , _ length

# print the explicitly concatenated lengths of "one", "two" and "three"
echo -n "one,two,three" | z split , _ length _ join ""
```

Without splits and joins, the same operations would have very different results:
```
# getting the length of infile.txt's content and writing that to outfile.txt
z length < infile.txt > outfile.txt

# print the md5 hash of "a:b:c"
echo -n "a:b:c" | z hash md5

# print the length of "one,two,three"
echo -n "one,two,three" | z length
```

The `match` command also returns an array of strings. Joining is done in exactly the same fashion as `split`:

```
# finding all occurrences of "findme" in file.txt and printing them, separated by commas
z match findme _ join , < file.txt
```

To better understand how `split`, `match` and `join` work, refer to our [help files](./help/) or run `z help`.

### Consuming from ever growing, endless streams
By default, z reads from its input until it reaches the end (EOF). However, there are some use cases where there's no expected end, such as tailing rotating log files. z approaches this by providing a `stream <delimiter>` command, which makes it consume in chunks separated by `<delimiter>`.

Here's an example:
```
# follows the tail of a mylogfile.log while printing all occurrences of 
# pattern "findme=[A-z]+ " joined by ","
# (the default stream delimiter is "\n", so we can omit the argument)
tail -f mylogfile.log | z stream _ match "/findme=[A-z]+ /" _ join ,
```

### Command reference

Check out our [help files](./help/).

If you already have z installed, avoid referring to this repo by running `z help`.

## Design principles
z was designed with the following principles in mind
1. **SIMPLE INTERFACE**. All z commands must have obvious names and perform clear, well-defined operations. Any user should be able to understand what their command chain does without referring to any documentation.
2. **EASY INSTALLATION**. All z releases must export a single lightweight binary. Installing it should be as simple as downloading the binary and including it in the `$PATH`. Want to uninstall? Just delete it.
3. **NO EXTERNAL DEPENDENCIES**. z must only depend on the Go core library functionality. This avoids the dependency hell and potential security vulnerabilities.


## NOTICE! Z IS STILL A WORK IN PROGRESS 
z is still under development and many of its features are not implemented yet. Check out development progress [here](./TODO.md).

For this reason, z is still not stable and may change a lot in the upcoming weeks/months. **DO NOT USE IT FOR SERIOUS APPLICATIONS**
