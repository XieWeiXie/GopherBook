#!/usr/bin/env bash

numbers=(0 1 2 3 4 5 6)
for VAR in ${numbers[*]} ; do
    echo ${VAR}
    page=$[36-${VAR}]
    echo ${page}
    echo `curl 'https://www.fina-gwangju2019.com/pg/sportEntriesData.do' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' --data 'sn=${VAR}'`

done
