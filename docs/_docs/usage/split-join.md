---
title: Splits and joins
prev_page: /docs/usage/chaining/
next_page: /docs/usage/stream/
permalink: /docs/usage/split-join/
---

z's input is always interpreted as a byte array or a string. However, there may be the need for splitting it into an array of strings and mapping operations onto the array elements. The z way of approaching this is via `split` and `join`:
1. `split` the inputs by a delimiter
2. map every split element using normal z commands such as `length`, `hash` or `replace`
3. implicitly concatenate them back into a string or `join` them with a delimiter

Here are some examples:
```
# getting the length of every line in infile.txt and writing that to outfile.txt's lines
# (split's default delimiter is "\n")
z split _ length _ join "\n" < infile.txt > outfile.txt

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

## Other split and join commands

There are other commands which either split or join arrays using different algorithms.

`match` splits input into an array of regex pattern matches.

```
# finding all words in file.txt and printing them, separated by commas
z match "[A-z]+" _ join , < file.txt
```

`count` joins an array by counting how many elements are in it.

```
# counting words in file.txt
z match "[A-z]+" _ count < file.txt
```

To better understand how split and join operations work, refer to the [command reference](../../commands/) or run `z help`.
