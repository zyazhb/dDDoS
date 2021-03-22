#!/bin/bash
count=1000

while(($count>0))
do
    curl -s "172.30.78.223" > /dev/null
    let "count--"
done

sleep 3

count=1000

while(($count>0))
do
    curl -s "172.30.78.223" > /dev/null
    let "count--"
done