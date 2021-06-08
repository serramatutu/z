---
title: match
permalink: /docs/commands/match/
description: Split a string into an array of regex matches.
arguments:
- name: pattern
  optional: false
  type: pattern
  description: The regular expression pattern to be matched.
examples: |
  # matching all lowercase words
  z match "[a-z]+"

  # matching all whitespace characters
  z match \s

  # getting the lengths of each word
  z match "[A-z]+" _ length

  # counting words
  z match "[A-z]+" _ count
---
