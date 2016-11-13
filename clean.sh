#!/bin/bash
#
# clean all go directories inside current directory
$ROOT=$(pwd)
go clean
for dir in $(ls -d */)
do
    up=$(pwd)
    echo $dir
    cd $dir
    go clean
    cd $up
done
