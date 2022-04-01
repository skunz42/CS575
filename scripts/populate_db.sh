#!/bin/bash

declare -a cities

while IFS='' read -r line || [[ -n "$line" ]]; do
    cities+=("$line")
done < $1

for i in ${cities[@]}; do
    echo "$i"
done
