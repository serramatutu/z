---
title: hash
permalink: /docs/commands/hash/
description: Hash a string according to an algorithm.
arguments:
- name: algorithm
  optional: false
  type: enum
  enum:
  - md5
  - sha1
  - sha256
  description: Which algorithm to use for hashing.
examples: |
  # hash a string with md5
  z hash md5

  # hash a string with sha1
  z hash sha1

  # hash a string with sha256
  z hash sha256
---
