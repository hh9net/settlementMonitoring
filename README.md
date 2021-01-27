# settlementMonitoring
结算数据监控平台
# settlementMonitoring
# conf 静态配置文件
# config 配置文件初始化
# db  数据库初始化文件
# docs  生成接口文档
# dto request和reponse的数据结构
# router 路由文件
# utils 常用工具

- 关于迁移新平台时出现的topic、redis等相关的问题的解析与解决方案的文档记录的生成;

1.
        腾讯云服务器部署结算监控平台时，发现topic有问题，核心原因是自己的topic没有在kafka那边注册。
        Kafka中间件对与topic进行了严格的管理
        解决方案：
        删除无关的topic，把所用到的topic全部在通行宝那边完成注册。

2.
        redis 的问题。所有的值都需要经过Json处理，这是java那边共用redis的一系列的key的值，
        所以需要对一些值进行初始化才可以。
        解决办法：
        触发初始化redis的值的接口：http://***.*.*.*:8089/settlementMonitoring/api/v1/sw/setredis；
        注意替换ip

