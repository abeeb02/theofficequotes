#!/bin/bash
for i in `seq 1 9`;
do
	curl https://the-office-api.herokuapp.com/season/$i/format/quotes/ > season$i.json
	echo "Season $i quotes downloaded"
done
