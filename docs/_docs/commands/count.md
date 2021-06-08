---
title: count
permalink: /docs/commands/count/
description: Join a split array by counting how many elements are in it.
examples: |
  # check how many elements the split by ":" produced
  z split : _ count

  # count words in input
  z match [A-z]+ _ count
---
