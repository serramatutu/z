---
layout: default
title: z, pipes made easy
permalink: /
redirect_from: 
- /home
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
