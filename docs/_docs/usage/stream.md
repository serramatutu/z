---
title: Consuming from ever growing, endless streams
permalink: /docs/usage/stream/
---

By default, z reads from its input until it reaches the end (EOF). However, there are some use cases where there's no expected end, such as tailing rotating log files. z approaches this by providing a `stream <delimiter>` command, which makes it consume in chunks separated by `<delimiter>`.

Here's an example:
```
# follows the tail of a mylogfile.log while printing all occurrences of 
# pattern "findme=[A-z]+ " joined by ","
# (the default stream delimiter is "\n", so we can omit the argument)
tail -f mylogfile.log | z stream _ match "/findme=[A-z]+ /" _ join ,
```
