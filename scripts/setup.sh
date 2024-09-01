#!/bin/bash

cd ./front/
nvm use v20.9.0
npm install

cd ../back
go mod download
