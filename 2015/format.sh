#!/usr/bin/env bash

for file in ./*; do
  case "$file" in
  *.txt)
    echo "Skipping: $file"
    ;;
  *.c)
    echo "Formatting: $file"
    clang-format -i "$file"
    ;;
  *)
    echo "Skipping $file"
    ;;
  esac
done
