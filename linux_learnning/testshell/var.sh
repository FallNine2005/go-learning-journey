#!/bin/bash
#设置变量A
A=100
echo $A
readonly B=200
echo $B
C=`date`
D=$(date)
echo $C
echo "D=$D"
#
echo "0=$0,1=$1,2=$2"
#
S=$[$1+$2]
echo "结果=$S"
