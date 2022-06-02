#!/bin/sh

set -x # print all executed commands
set -e # exit on error

emulatorPort="8090"

go get ./...

go fmt ./...

rm -f firestore
go build -o firestore

if [ $(lsof -i:$emulatorPort -t) ]
then
    kill -9 $(lsof -i:$emulatorPort -t);
fi

gcloud beta emulators firestore start --host-port=localhost:$emulatorPort &
sleep 10
PORT="8080" GCP_PROJECT="patch-emulator-project" FIRESTORE_EMULATOR_HOST="localhost:$emulatorPort" bash -c './firestore'