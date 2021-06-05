---
title: Command chaining
prev_page: /docs/usage/basic/
next_page: /docs/usage/split-join/
permalink: /docs/usage/chaining/
---

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
