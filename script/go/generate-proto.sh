#!/bin/bash

if [[ -z $SRC_DIR ]]; then
  echo "Please, specify SRC_DIR"
  exit 1
fi

if [[ -z $OUT_DIR ]]; then
  echo "Please, specify OUT_DIR"
  exit 1
fi

SRC_DIR=${SRC_DIR:-}
OUT_DIR=${OUT_DIR:-}

SOURCES=$(find ${SRC_DIR} -iname "*.proto")

protoc \
  --proto_path=${SRC_DIR} \
  \
  --go_out=${OUT_DIR} \
  --go-grpc_out=${OUT_DIR} \
  \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  ${SOURCES}
