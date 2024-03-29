#!/bin/bash

set -e

rm -rf ./data

rm -rf ./server/__debug_bin
rm -rf ./server/server
rm -rf ./server/tmp*
rm -rf ./server/data

if [[ $1 == 'all' ]] 
then

    rm -rf ./server/build/linux64/data
    rm -rf ./server/build/linux64/tmp-locker
    rm -rf ./server/build/linux64/server*
    rm -rf ./server/build/linux64/*.txt

    rm -rf ./server/build/linuxarm/data
    rm -rf ./server/build/linuxarm/tmp-locker
    rm -rf ./server/build/linuxarm/server*
    rm -rf ./server/build/linuxarm/*.txt

    rm -rf ./server/build/win64

else

    rm -rf ./server/build/linux64/tmp*
    rm -rf ./server/build/linux64/server*
    rm -rf ./server/build/linux64/data/id-event
    rm -rf ./server/build/linux64/data/id-subs
    rm -rf ./server/build/linux64/data/span-ids
    rm -rf ./server/build/linux64/data/owner-ids
    rm -rf ./server/build/linux64/*.txt

    rm -rf ./server/build/linuxarm/tmp*
    rm -rf ./server/build/linuxarm/server*
    rm -rf ./server/build/linuxarm/data/id-event
    rm -rf ./server/build/linuxarm/data/id-subs
    rm -rf ./server/build/linuxarm/data/span-ids
    rm -rf ./server/build/linuxarm/data/owner-ids
    rm -rf ./server/build/linuxarm/*.txt

    rm -rf ./server/build/win64/tmp*
    rm -rf ./server/build/win64/server*
    
fi

rm -rf ./server/build/*.gz
