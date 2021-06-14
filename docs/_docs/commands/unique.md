---
title: unique
permalink: /docs/commands/unique/
description: | 
  Get unique elements of an array of strings based on keys extracted from each string. 

  "unique" must come before a splitting such as "split" or "match".	
  
  If it finds repeated values, the first to be found is kept
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
  description: The index of the field that should be used as the key for determining uniqueness. If some string has less fields than the index, "" is assumed as its uniqueness key.
examples: |
  # getting the unique lines of a file
  z split "\n" _ unique _ join "\n" < file.txt

  # getting unique rows of a csv file based on its first column
  z split "\n" _ unique "," 0 _ join "\n" < data.csv

  # getting unique rows of a tsv file based on its third column
  z split "\n" _ unique "\t" 3 _ join "\n" < data.tsv
---
