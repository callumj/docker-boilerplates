#!/usr/bin/env bash

# copy our source code into our /app directory (which already has node_modules thanks to the Docker image)
cp -r /app_staging/* /app/
cd /app
# build it
npm start
# extract the bundle
cp -r /app/dist/* /app_staging/dist/