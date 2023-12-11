#!/bin/bash

echo Creating template for day $1: $2

DIR=$1"-"$2
mkdir $DIR

cd $DIR
go mod init $2
touch input.txt
cat ../template.go >> $2.go

cd ..
go work use $DIR
cd $DIR