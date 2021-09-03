#!/usr/bin/env sh
mkdir meilidata
mkdir backend || exit 1
mkdir frontend || exit 1
cd database/ || exit 1
mkdir "postgres"
