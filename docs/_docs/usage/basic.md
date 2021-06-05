---
title: Usage basics
permalink: /docs/usage/basic/
redirect_from: 
- /docs/
- /docs/usage/
---

z is built to have an easy and intuitive interface without giving up on functionality.

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
