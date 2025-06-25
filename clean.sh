#! /bin/bash

found=
while IFS= read -r line; do
    if [[ $line =~ "# Demo programs" ]]; then found=1; fi
    [[ $found && -x $line ]] && rm -f $line
done < .gitignore

find . -name screenshot.png -delete