../../bin/manager -c ../configs/manager/config.json >> info.log &
sleep 1s
../../bin/login -c ../configs/login/config.json >> info.log &
sleep 1s
../../bin/dbserver -c ../configs/dbserver/config.json >> info.log &
sleep 1s
../../bin/game -c ../configs/game/config.json >> info.log & 
sleep 1s
../../bin/gateway -c ../configs/gateway/config.json  >> info.log &
