#!/usr/bin/env bash

NUM_NODE=4

# tackle size chronic goose deny inquiry gesture fog front sea twin raise
# acid pulse trial pill stumble toilet annual upgrade gold zone void civil
# antique onion adult slot sad dizzy sure among cement demise submit scare
# lazy cause kite fence gravity regret visa fuel tone clerk motor rent
HARDCODED_MNEMONIC=true

set -e
set -o errexit
set -a
set -m

set -x # activate debugging

source oec.profile
WRAPPEDTX=false
PRERUN=false
while getopts "isn:b:p:c:Smxwk:" opt; do
  case $opt in
  i)
    echo "OKCHAIN_INIT"
    OKCHAIN_INIT=1
    ;;
  w)
    echo "WRAPPEDTX=$OPTARG"
    WRAPPEDTX=true
    ;;
  x)
    echo "PRERUN=$OPTARG"
    PRERUN=true
    ;;
  s)
    echo "OKCHAIN_START"
    OKCHAIN_START=1
    ;;
  k)
    echo "LOG_SERVER"
    LOG_SERVER="--log-server $OPTARG"
    ;;
  c)
    echo "Test_CASE"
    Test_CASE="--consensus-testcase $OPTARG"
    ;;
  n)
    echo "NUM_NODE=$OPTARG"
    NUM_NODE=$OPTARG
    ;;
  b)
    echo "BIN_NAME=$OPTARG"
    BIN_NAME=$OPTARG
    ;;
  S)
    STREAM_ENGINE="analysis&mysql&localhost:3306,notify&redis&localhost:6379,kline&pulsar&localhost:6650"
    echo "$STREAM_ENGINE"
    ;;
  p)
    echo "IP=$OPTARG"
    IP=$OPTARG
    ;;
  m)
    echo "HARDCODED_MNEMONIC"
    HARDCODED_MNEMONIC=true
    ;;
  \?)
    echo "Invalid option: -$OPTARG"
    ;;
  esac
done

echorun() {
  echo "------------------------------------------------------------------------------------------------"
  echo "["$@"]"
  $@
  echo "------------------------------------------------------------------------------------------------"
}

killbyname() {
  NAME=$1
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2", "$8}'
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2}' | sh
  echo "All <$NAME> killed!"
}

init() {
  killbyname ${BIN_NAME}

  (cd ${OKCHAIN_TOP} && make install VenusHeight=1)

  rm -rf cache

  echo "=================================================="
  echo "===== Generate testnet configurations files...===="
  echorun exchaind testnet --v $1 -o cache -l \
    --chain-id ${CHAIN_ID} \
    --starting-ip-address ${IP} \
    --base-port ${BASE_PORT} \
    --keyring-backend test \
    --mnemonic=${HARDCODED_MNEMONIC}
}

run() {

  index=$1
  seed_mode=$2
  p2pport=$3
  rpcport=$4
  restport=$5
  p2p_seed_opt=$6
  p2p_seed_arg=$7


  if [ "$(uname -s)" == "Darwin" ]; then
      sed -i "" 's/"enable_call": false/"enable_call": true/' cache/node${index}/exchaind/config/genesis.json
      sed -i "" 's/"enable_create": false/"enable_create": true/' cache/node${index}/exchaind/config/genesis.json
      sed -i "" 's/"enable_contract_blocked_list": false/"enable_contract_blocked_list": true/' cache/node${index}/exchaind/config/genesis.json
      #sed -i "" 's/node_key_whitelist.*/node_key_whitelist = ["00323d59a5ce412d58dbeb79450775aa78cde477", "850502b4ecd570f1d13fd63e6c6ca96ae8f7f947", "91f41250223bc59203cae5f6a75d6df30653b8cf", "dc50439a015454d53ebb8407393c45c7120f057e"]/' cache/node${index}/exchaind/config/config.toml
  else
      sed -i 's/"enable_call": false/"enable_call": true/' cache/node${index}/exchaind/config/genesis.json
      sed -i 's/"enable_create": false/"enable_create": true/' cache/node${index}/exchaind/config/genesis.json
      sed -i 's/"enable_contract_blocked_list": false/"enable_contract_blocked_list": true/' cache/node${index}/exchaind/config/genesis.json
      #sed -i 's/node_key_whitelist.*/node_key_whitelist = ["00323d59a5ce412d58dbeb79450775aa78cde477", "850502b4ecd570f1d13fd63e6c6ca96ae8f7f947", "91f41250223bc59203cae5f6a75d6df30653b8cf", "dc50439a015454d53ebb8407393c45c7120f057e"]/' cache/node${index}/exchaind/config/config.toml
  fi

  exchaind add-genesis-account 0xbbE4733d85bc2b90682147779DA49caB38C0aA1F 900000000okt --home cache/node${index}/exchaind
  exchaind add-genesis-account 0x4C12e733e58819A1d3520f1E7aDCc614Ca20De64 900000000okt --home cache/node${index}/exchaind
  exchaind add-genesis-account 0x83D83497431C2D3FEab296a9fba4e5FaDD2f7eD0 900000000okt --home cache/node${index}/exchaind
  exchaind add-genesis-account 0x2Bd4AF0C1D0c2930fEE852D07bB9dE87D8C07044 900000000okt --home cache/node${index}/exchaind

  LOG_LEVEL=main:debug,*:error,consensus:error,state:info,ante:info,txdecoder:info

  echorun nohup exchaind start \
    --home cache/node${index}/exchaind \
    --p2p.seed_mode=$seed_mode \
    --p2p.allow_duplicate_ip \
    --enable-dynamic-gp=false \
    --enable-wtx=${WRAPPEDTX} \
    --p2p.pex=false \
    --p2p.addr_book_strict=false \
    $p2p_seed_opt $p2p_seed_arg \
    --p2p.laddr tcp://${IP}:${p2pport} \
    --rpc.laddr tcp://${IP}:${rpcport} \
    --consensus.timeout_commit 600ms \
    --log_level ${LOG_LEVEL} \
    --chain-id ${CHAIN_ID} \
    --upload-delta=false \
    --enable-gid \
    --append-pid=true \
    ${LOG_SERVER} \
    --elapsed DeliverTxs=0,Round=1,CommitRound=1,Produce=1 \
    --rest.laddr tcp://localhost:$restport \
    --enable-preruntx=$PRERUN \
    --consensus-role=v$index \
    ${Test_CASE} \
    --keyring-backend test >cache/val${index}.log 2>&1 &

#     --iavl-enable-async-commit \    --consensus-testcase case12.json \
#     --upload-delta \
#     --enable-preruntx \
}

function start() {
  killbyname ${BIN_NAME}
  index=0

  echo "============================================"
  echo "=========== Startup seed node...============"
  ((restport = REST_PORT)) # for evm tx
  run $index true ${seedp2pport} ${seedrpcport} $restport
  seed=$(exchaind tendermint show-node-id --home cache/node${index}/exchaind)

  echo "============================================"
  echo "======== Startup validator nodes...========="
  for ((index = 1; index < ${1}; index++)); do
    ((p2pport = BASE_PORT_PREFIX + index * 100 + P2P_PORT_SUFFIX))
    ((rpcport = BASE_PORT_PREFIX + index * 100 + RPC_PORT_SUFFIX))  # for exchaincli
    ((restport = index * 100 + REST_PORT)) # for evm tx
    run $index false ${p2pport} ${rpcport} $restport --p2p.seeds ${seed}@${IP}:${seedp2pport}
  done
  echo "start node done"
}

if [ -z ${IP} ]; then
  IP="127.0.0.1"
fi

if [ ! -z "${OKCHAIN_INIT}" ]; then
  init ${NUM_NODE}
fi

if [ ! -z "${OKCHAIN_START}" ]; then
  start ${NUM_NODE}
fi
