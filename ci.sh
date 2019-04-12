#!/bin/bash
#build Linux on mac

#echo "update vendor..."
#export GO111MODULE=on
#go mod tidy
#go mod vendor

echo "building...."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gateway/msgGateway gateway/cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/msgLogic/msgLogic msgLogic/cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/userLogic/userLogic userLogic/cmd/main.go

echo "scp is starting..."
scp -o stricthostkeychecking=no -r ./build/* ubuntu@kuipmake.com:/usr/local/src/kuipmake/msg/release
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "rm -rf /usr/local/src/kuipmake/msg/gateway /usr/local/src/kuipmake/msg/msgLogic /usr/local/src/kuipmake/msg/userLogic && mv /usr/local/src/kuipmake/msg/release/* /usr/local/src/kuipmake/msg"

#由于每个ssh不会自动退出，所以下面每条需要单独运行
echo "running..."
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "pkill msgGateway"
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "cd /usr/local/src/kuipmake/msg/gateway && nohup ./msgGateway 1>msgGateway.log 2>msgGateway.err &"
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "pkill msgLogic"
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "cd /usr/local/src/kuipmake/msg/msgLogic && nohup ./msgLogic 1>msgLogic.log 2>msgLogic.err &"
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "pkill userLogic"
ssh -o stricthostkeychecking=no ubuntu@kuipmake.com "cd /usr/local/src/kuipmake/msg/userLogic && nohup ./userLogic 1>userLogic.log 2>userLogic.err &"

echo "ci ok"