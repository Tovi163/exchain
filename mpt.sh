killall exchaind
rm -rf multi_run.log
make mainnet
echo 3 > /proc/sys/vm/drop_caches 
export EXCHAIND_PATH=~/.exchaind

#rm -rf ${EXCHAIND_PATH}/
#exchaind init multi_run --chain-id exchain-66 --home ${EXCHAIND_PATH}
#cp /Users/oker/scf/data/genesis.json ${EXCHAIND_PATH}/config/genesis.json

exchaind start --chain-id exchain-66 --log_level="main:info,evm:info,*:error"  --home ${EXCHAIND_PATH} > multi_run.log 2>&1 & 

