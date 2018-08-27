### 结构

gateway 负责长连接，可分布式部署

logic 负责处理业务逻辑，无状态

### 代码

https://www.processon.com/view/link/5a0998cae4b049e7f4fcdcdc

### 计划

* 基础推送网关
    * [x] 保持用户长链接
    * [x] 连接数量统计
    * [x] 消息收发
    * [ ] 消息存储      hui
    * [ ] 登录逻辑      hui
    * [ ] 群组处理
    * [ ] 分布式
    * [ ] 推送 接口        gao

* basic （8月31日）
    * [ ] 消息格式制定，登录、普通消息、图文、@某人 自定义   jason
    * [ ] 表结构制定     jason

* logic-user （8月31日接口文档 9月7日完成接口）
    * [ ] admin-dashboard  xxxxx
    * [ ] app 应用管理接口  crud app       gao
    * [ ] app 用户信息管理接口             hui
    * [ ] 用户登录接口，gateway rpc调用,jwt  hui

* logic-push （8月31日接口文档 9月7日完成）
    * [ ] push 特定用户    gao
    * [ ] push 所有用户    gao
    * [ ] push 标签用户    gao
    * [ ] push 历史记录接口 gao
    * [ ] 前端 长连接 认证  shao

* logic-group
    * [ ] group-crud

* logic-msg
    * [ ] 聊天 ui
    * [ ] 用户登录
    * [ ] 消息处理
    * [ ] 聊天记录
    * [ ] 群组
    * [ ] 消息存储
    * [ ] rabbitmq
