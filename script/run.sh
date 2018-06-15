../bin/manager -c ../config/manager/config.json >> info.log &
sleep 1s
../bin/login -c ../config/login/config.json >> info.log &
sleep 1s
../bin/dbserver -c ../config/dbserver/config.json >> info.log &
sleep 1s
../bin/game -c ../config/game/config.json >> info.log & 
sleep 1s
../bin/gateway -c ../config/gateway/config.json  >> info.log &
