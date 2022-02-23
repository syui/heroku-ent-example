#!/bin/zsh

t=`grep -R heroku-ent-example .`
n=`echo $t|wc -l`
echo $n
for ((i=1;i<=$n;i++))
do
	tt=`echo $t|awk "NR==$i"|cut -d : -f 1`
	sed -i "" 's/heroku-ent-example/heroku-ent-example/g' $tt
done

