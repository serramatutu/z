---
title: length
permalink: /docs/commands/length/
description: Measure the length of a string
arguments:
- name: mode
  optional: true
  default: bytes
  type: enum
  enum:
  - bytes
  - unicode
  description: Whether to measure the length in bytes or in unicode characters
examples: |
  # get the length in bytes of "maçã", which equals 6
  echo -n "maçã" | z length

  # get the length in unicode characters of "maçã", which equals 4
  echo -n "maçã" | z length unicode
---
