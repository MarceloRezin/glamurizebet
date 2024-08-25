#!/bin/bash

cd ./front/
. ~/.nvm/nvm.sh # Sem isso o bash não identifica o comando nvm
nvm use v20.9.0
npm start