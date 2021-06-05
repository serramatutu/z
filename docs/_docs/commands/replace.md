---
title: replace
permalink: /docs/commands/replace/
description: Find occurrences of a pattern and replace them with a string.
arguments:
  - name: pattern
    optional: false
    type: pattern
    description: The regular expression pattern to be matched.
  - name: replace-string
    optional: false
    type: string
    description: The string that will replace the matches.
  - name: occurrence-range
    optional: false
    type: range
    default: "0:0"
    description: >
      The half-open match index range where replacement will occur.
      Negative range end means counting in reverse order.
      Zero range end means last match.
examples: |
  # replacing ":" by "\n"
  z replace : \n

  # removing all ":"
  z replace : ""

  # replacing all alphabetical characters with "0"
  z replace [A-z] 0

  # replacing all but first "0" with "zero"
  z replace 0 "zero" 1:

  # replacing all but first and last "1" with "one"
  z replace 1 "one" 1:-1

  # removing the first "."
  z replace . "" 0:1
---
