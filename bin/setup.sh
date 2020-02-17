#!/bin/bash
echo "DUMBP SETUP"

echo "Enter your google api key:"
read KEY
echo "GOOGLE_API_KEY=\"$KEY\"" >> test/.env

go-bindata -pkg data -o ./bin/data/data.go ./bin/data/...
