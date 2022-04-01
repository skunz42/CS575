import os
import sys

if len(sys.argv) != 2:
    print("Need two args")
    sys.exit(1)

with open(sys.argv[1]) as f:
    lines = f.read().splitlines()
    for i in range(len(lines)):
        for j in range(i+1, len(lines)):
            print(f"{lines[i]} -> {lines[j]}")
