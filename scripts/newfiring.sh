#!/bin/sh
# append one row from the console
echo "$(date -u +%Y-%m-%d),$1,$2,$3,$4,$USER,$5" >> ../log/firings.csv
