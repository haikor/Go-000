# 学习笔记

## 作业
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

- 创建两个http服务 8081、8082
- 定义exit chan来监听退出事件
- 监听/shutdown 路由，触发退出事件
- 请求浏览器 http://localhost:8081/shutdown 可关闭所有web服务
- 控制台按Ctrl+C 关闭所有web服务