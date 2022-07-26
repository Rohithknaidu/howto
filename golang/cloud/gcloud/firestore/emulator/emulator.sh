#!/bin/sh

set -x # print all executed commands
set -e # exit on error

emulatorPort="8090"


gcloud beta emulators firestore start --host-port=localhost:$emulatorPort &
sleep 10


go get ./...
GCP_PROJECT="ytf-emulator-project" FIRESTORE_EMULATOR_HOST="localhost:$emulatorPort" go run emulator.go

lsof -i:$emulatorPort -t -s TCP:LISTEN

kill -9 $(lsof -i:$emulatorPort -t -s TCP:LISTEN)
sleep 2
