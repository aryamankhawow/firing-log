# import the 2008 paper log, one row per line
import sys

for line in sys.stdin:
    parts = line.strip().split()
    if len(parts) < 5:
        continue
    print(",".join(parts))
