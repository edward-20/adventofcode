#!/bin/sh

while read line
do
	if test -z "${line// }"
	then
		break
	fi
	echo $line
done < input.txt
