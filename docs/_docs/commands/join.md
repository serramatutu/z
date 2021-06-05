---
title: join
permalink: /docs/commands/join/
description: |
  Join an array of strings into a string.

  "join" commands are used anywhere after a "split" command.
    1. transform input contents into array with "split"
    2. map the array elements using other commands such as "hash", "replace" or "length"
    3. join the array back into a string by concatenating its entries implicitly or by using "join"

  Split arrays with no closing join have their elements implicitly concatenated.
arguments:
  - name: delimiter
    optional: true
    type: string
    default: ""
    description: The string to be used as delimiter.
examples: |
  # get the lengths of all lines in a file and concatenate them (implicit join)
  z split _ length

  # get the lengths of all lines in a file and join them with ","
  z split _ length _ join ,
---
