#!/bin/bash


if [ ! -f ".env" ]; then
  cp .env.example .env
fi


npm install

# chmod -R 775 /home/node/app/dist


npm run start:dev