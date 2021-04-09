# tf_server

tf_server项目是TravelFriend学习项目的子项目，使用golang实现的服务端。主要涉及旅行主题社交，业务包含用户、团队、动态、相册、推荐等。

##### 主要技术栈包含

```
gin	golang的web框架
gorm	orm框架
go-redis	redis缓存
go-mapper	dto对象映射
jwt-go	jwt验证
aliyun-oss-go-sdk	oss对象存储
...
```

#### 项目结构

- ##### api
  webapi文件目录，用于存放不同版本的api接口。restful接口设计，**只用post**。
- ##### application
  应用层目录，主要用于存放命令（command）与查询（query）职责分离及validator验证等应用层所需功能。
- ##### config
  配置文件，包括ini文件及其映射文件。
- ##### middleware
  中间件，用于存放各种gin框架使用的中间件。
- ##### model
  模型定义，仅用于存放数据库模型定义及数据库相关（db）文件。
- ##### repository
  仓储目录，封装各种业务对应的数据库相关操作，包括缓存、事务等。
- ##### route
  路由
- ##### util
  工具目录，存放各种工具组件的封装。

###### 目录结构调整完善中...

