#!/bin/bash

if [[ $# -gt 0 ]]; then
  printf "Regular : %s\n" "$1"
  printf "Bold    : \033[1m%s\033[0m\n" "$1"
  printf "Italic  : \033[3m%s\033[0m\n" "$1"
else
  printf "This is normal text. \033[1mThis is bold text!\033[0m \033[3mThis is italic text!\033[0m\n"
fi
