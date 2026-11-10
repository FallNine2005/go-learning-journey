#!/bin/bash
function SUM()
{
	testF=$[ $n1 + $n2 ]
	echo "n1+n2=$testF"
}

read -p "n1=" n1
read -p "n2=" n2
SUM
