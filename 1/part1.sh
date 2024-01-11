#!/bin/bash

sum=0
while read -r line
do
	len=${#line}

	firstPos=-1
	for i in $(seq 0 $((len - 1)))
	do
		if [ ${line:i:1} -ge 0 ] 2> /dev/null
		then
			firstPos=$i
			break
		fi
	done
	firstDigit=${line:firstPos:1}

	lastPos=-1
	for i in $(seq 0 $((len - 1)))
	do
		if [ ${line:i:1} -ge 0 ] 2> /dev/null
		then
			lastPos=$i
		fi
	done
	lastDigit=${line:lastPos:1}

	number=$firstDigit$lastDigit
	sum=$((sum + number))
done < input.txt

echo $sum
