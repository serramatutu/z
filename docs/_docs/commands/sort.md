---
title: sort
permalink: /docs/commands/sort/
description: | 
  Sort an array of strings based on keys extracted from each string

  "sort" must come before a splitting command such as "split" or "match"

  Sort implements an unstable sorting algorithm, which means that elements with the same sorting key are not guaranteed to be kept in the same order

  Keys are compared alphabetically
arguments:
- name: separator
  optional: true
  default: null
  type: pattern
  description: The regular expression pattern to be used as a field delimiter. If not provided, z assumes the whole string is the key.
- name: key-index
  optional: true
  default: 0
  type: number
  description: The index of the field that should be used as the sorting key. If some string has less fields than the index, "" is assumed as its sorting key.
examples: |
  # sorting the lines of a file
  z split "\n" _ sort _ join "\n" < file.txt

  # sorting the rows of a csv file based on its first column
  z split "\n" _ sort "," 0 _ join "\n" < data.csv

  # sorting the rows of a tsv file based on its third column
  z split "\n" _ sort "\t" 3 _ join "\n" < data.tsv
---
