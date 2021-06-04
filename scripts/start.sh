#!/bin/bash

ulimit -n 65535 1>/dev/null 2>&1

BIN=./bin/centnet-fzmps
PROFILE="./conf/config.toml"

PID=$(ps -ef | grep ${BIN} | grep -v grep | tr -s ' ' | cut -d ' ' -f 2)
if [[ -n "${PID}" ]]; then
  echo "${BIN}(pid: $PID) is running."
  exit
fi

chmod 777 ${BIN}
case $1 in
-d)
  nohup ${BIN} -c ${PROFILE} 1>/dev/null 2>&1 &
  ;;
*)
  ${BIN} -c ${PROFILE}
  ;;
esac
