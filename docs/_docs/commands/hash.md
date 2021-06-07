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
  - sha224
  - sha256
  description: Which algorithm to use for hashing. md5 and sha1 are not criptographically secure and should not be used for secure applications.
examples: |
  # hash a string with md5
  z hash md5

  # hash a string with sha1
  z hash sha1
	
	# hash a string with sha224
  z hash sha224

  # hash a string with sha256
  z hash sha256
---
