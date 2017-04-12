#!/bin/bash
echo "building ..."
rm -rf ./dbweb
mkdir ./dbweb
go build -tags sqlite3 -v -o ./dbweb/dbweb
cp -r ./public ./dbweb/
cp -r ./templates ./dbweb/
cp -r ./options ./dbweb/
cp ./*.pem ./dbweb/
echo "done."
