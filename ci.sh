#!/bin/bash
#build Linux on mac

#echo "update vendor..."
#dep ensure -v

echo "building...."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gateway/msgGateway gateway/cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/msgLogic/msgLogic msgLogic/cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/userLogic/userLogic userLogic/cmd/main.go

echo "scp is starting..."
export SSHPASS='ipar$1000'
sshpass -e scp -o stricthostkeychecking=no -P 1607 -r ./build/* drone@ideapar.com:~/msg/release
sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "rm -rf ~/msg/gateway ~/msg/msgLogic ~/msg/userLogic && mv ~/msg/release/* ~/msg"

echo "running..."
sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "pkill msgGateway"
sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "cd ~/msg/gateway && nohup ./msgGateway 1>msgGateway.log 2>msgGateway.err &"

sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "pkill msgLogic"
sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "cd ~/msg/msgLogic && nohup ./msgLogic 1>msgLogic.log 2>msgLogic.err &"

sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "pkill userLogic"
sshpass -e ssh -o stricthostkeychecking=no -p 1607 drone@ideapar.com "cd ~/msg/userLogic && nohup ./userLogic 1>userLogic.log 2>userLogic.err &"

echo "ci ok"