#!/bin/bash

set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
W=`tput sgr0`

./clean.sh

cd ./server

# create version info as hard-coded 
cd ./static-cfg
go run .
cd -

LDFLAGS="-s -w"
TM=`date +%F@%T@%Z`
OUT=wisite-api\($TM\)

######################### X86 #########################

GOARCH=amd64

# For Docker, one build below for linux64 is enough.
OUTPATH_LINUX=./build/linux64/
mkdir -p $OUTPATH_LINUX
CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
mv $OUT $OUTPATH_LINUX
cp -r ./*config.json $OUTPATH_LINUX
# cp -r ./www $OUTPATH_LINUX 
echo "${G}${OUT}(linux64)${W} built"

# OUTPATH_WIN=./build/win64/
# mkdir -p $OUTPATH_WIN
# CGO_ENABLED=0 GOOS="windows" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT.exe
# mv $OUT.exe $OUTPATH_WIN
# cp -r ./*config.json $OUTPATH_WIN
# # cp -r ./www $OUTPATH_WIN
# echo "${G}${OUT}(win64)${W} built"

# OUTPATH_MAC=./build/mac/
# mkdir -p $OUTPATH_MAC
# CGO_ENABLED=0 GOOS="darwin" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH_MAC
# cp -r ./*config.json $OUTPATH_MAC
# # cp -r ./www $OUTPATH_MAC
# echo "${G}${OUT}(mac)${W} built"

######################### ARM #########################

# GOARCH=arm

# OUTPATH_LARM=./build/linuxarm/
# mkdir -p $OUTPATH_LARM
# CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" GOARM=7 go build -ldflags="$LDFLAGS" -o $OUT
# mv $OUT $OUTPATH_LARM
# cp -r ./*config.json $OUTPATH_LARM
# # cp -r ./www $OUTPATH_LARM
# echo "${G}${OUT}(linuxArm)${W} built"

#######################################################################################

if [[ $1 == 'release' || $1 == 'rel' ]]
then

    RELEASE_NAME=wisite-api\($TM\).tar.gz 
    cd ./build
    echo $RELEASE_NAME
    tar -czvf ./$RELEASE_NAME --exclude='./linux64/data' --exclude='./linux64/temp_test.go' --exclude='./win64/data'  ./linux64 # ./win64 ./mac ./linuxarm

fi