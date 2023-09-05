#!/bin/bash

if [[ -z $SRC_FILE ]]; then
  echo "Please, specify SRC_FILE"
  exit 1
fi

if [[ -z $OUT_DIR ]]; then
  echo "Please, specify OUT_DIR"
  exit 1
fi

if [[ -z $ROUTER ]]; then
  echo "Please, specify ROUTER"
  exit 1
fi

SRC_FILE=${SRC_FILE:-}
OUT_DIR=${OUT_DIR:-}

ROUTER=${ROUTER:-}

if [ ${ROUTER} = "echo" ]; then
  ROUTER_OPT="server"
  PACKAGE="echo"
fi
if [ ${ROUTER} = "chi" ]; then
  ROUTER_OPT="chi-server"
  PACKAGE="chi"
fi
if [ ${ROUTER} = "gin" ]; then
  ROUTER_OPT=${ROUTER}
  PACKAGE="gin"
fi
if [ ${ROUTER} = "gorilla" ]; then
  ROUTER_OPT=${ROUTER}
  PACKAGE="mux"
fi
if [ ${ROUTER} = "fiber" ]; then
  ROUTER_OPT=${ROUTER}
  PACKAGE="fiber"
fi

OUT_FILE="${OUT_DIR}/${ROUTER}/server.go"

oapi-codegen -package "i${PACKAGE}" -generate "types,${ROUTER_OPT},spec" ${SRC_FILE} >${OUT_FILE}
