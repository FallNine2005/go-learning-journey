#!/bin/bash
for num in 100 200 300
do
./if.sh $num
done

sum=0
for (( i=1; i<10; i++))
do
	sum=$[ $sum + $i ]
	echo "$sum"
done

