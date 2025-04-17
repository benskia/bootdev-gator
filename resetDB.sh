#!/usr/bin/env bash

(
  cd ./sql/schema/
  goose -v postgres "postgres://postgres:postgres@localhost:5432/gator" reset
  goose -v postgres "postgres://postgres:postgres@localhost:5432/gator" up
)
