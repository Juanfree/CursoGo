#!/usr/bin/env bash

PROJECT_PATH=src/github.com/example

TARGET_PATH="$HOME/go/$PROJECT_PATH"

mkdir -p $TARGET_PATH

cp -R CursoGO $TARGET_PATH

echo "Project is located in $TARGET_PATH"