#!/bin/sh

## Utility script to release project with a tag
## Usage:
##   ./release.sh <tag-name>

MODULE=semita

if [ "$1" == "" ]; then
	echo "Usage: $0 tag-name"
	exit -1
fi

echo "$MODULE/$1"
git commit -m "$MODULE/$1"
git tag -f -a "$MODULE/$1" -m "$MODULE/$1"
git push origin "$MODULE/$1" -f
git push

