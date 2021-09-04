#!/usr/bin/env sh
mkdir meilidata || exit 0
mkdir backend || exit 0
mkdir frontend || exit 0
cd database/ || exit 1
mkdir "postgres" || exit 1