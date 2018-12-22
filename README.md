### 结构

gateway 负责长连接，可分布式部署，提供websocket接口

msgLogic 负责处理消息业务，数据库存储，无状态，提供rpc接口

userLogic 负责用户、组织、app业务，数据库存储，提供http接口

https://www.processon.com/view/link/5a0998cae4b049e7f4fcdcdc

### todo

* 前端: 消息类型
* userLogic: 登录注册api
* msgLogic: 多人群聊的会话
* gateway: 推送push
* wss
* 第三方: 好友关系，在线列表
* 聊天机器人

### APP聊天的流程

1. kuip server: getLinkTokenByToken
2. websocket connect by linkToken
3. getLinkByUserId
4. listConv
5. listMsgHistory
6. sendMsg
    1. msgType: text image video audio duration location others
7. createConv
    1. kuip server: getLinkByUserId
    2. createConv
 
### websocket connect

* ws://kuipmake.com:12315/ws?token=001

### chat room api

* joinChatConv
    * ws://kuipmake.com:12315/ws
```
request:
{
    action: "chat-conv-join,
    linker: {},
    param: {convId}
}
```
* sendChatMsg
    * ws://kuipmake.com:12315/ws
```
request:
{
    action: "chat-msg-im
    linker: {name, key, avt},
    param: {msgKey, convId},
    data: {msgType, content}
}
```

### link api

* kuip server: getLinkTokenByToken

* kuip server: getLinkByUserId

### msg api

* msg-im
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: msg-im,
    linkKey: userId,
    param: {msgKey, convId, msgType, msgContent}
} 
```

* msg-read
    * ws://kuipmake.com:12315/ws

```
request:
{

    action: msg-read,
    linkKey: userId,
    param: {convId}
}
```

* msg-listHistory
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: msg-listHistory,
    linkKey: userId,
    param: {convId,limit,offset}
}
```
### conv api

* conv-create
    * ws://kuipmake.com:12315/ws 

```
request:
{
    action: conv-create,
    linkKey: userId,
    param: {convType: "single or multi", convKey: “random", convAvt, convName: “convName", linkKeys: []}
}
```

* conv-list
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: conv-list,
    linkKey: userId,
}
```

* conv-delete
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: conv-delete,
    linkKey,
    param: {convId}
}
```

* conv-join
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: conv-join,
    linkKey,
    param: {convId}
}
```

* conv-leave
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: conv-leave,
    linkKey,
    param: {convId}
}
```

* conv-inviteLinks
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: conv-inviteLinks,
    linkKey,
    param: {convId: xxx, linkKeys: []}
}
```

* conv-removeLinks
    * ws://kuipmake.com:12315/ws

```
request:
{
    action: conv-removeLinks,
    linkKey,
    param: {convId: xxx, linkKeys: []}
}
```

### user api

* reg
* login
* createOrg
* listOrg
* getOrg
* createApp
* listApp
* getApp

### local startup
* git clone [url]
* cd msg
* 编辑配置文件
* export GO111MODULE=on
* go mod tidy (终端需科学上网)
* go mod vendor
* run

### 部署到kuipmake.com的端口
* gateway
    * ws: 12315
* msgLogic
    * http: 12320
    * rpc: 12321
* userLogic
    * http: 12330
    * rpc: 12331
* mac本地部署，./ci.sh
