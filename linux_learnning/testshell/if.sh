#!/bin/bash
if [ "$1" -gt 100 ]
then
	echo "ge"
elif [ "$1" -eq 100 ]
then
	echo "eq"	
elif [ "$1" -lt 100 ] 
then
	echo "lt"
fi

