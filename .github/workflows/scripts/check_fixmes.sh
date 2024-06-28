#!/bin/bash
set -xu

while IFS= read -r -d '' file; do
    grep -rli FIXME "$file"
    return_code=$?
    if [ $return_code -eq 0 ]; then
        echo "FIXME found in code"
        exit 1
    elif [ $return_code -ne 1 ]; then
        echo "Error while detecting FIXMEs in code"
        exit 1
    fi
done< <(find . -type f -name '*.rs' -print0)
echo "No FIXMEs found in the code"