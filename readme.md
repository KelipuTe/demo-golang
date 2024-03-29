## 说明（readme）

放 golang 代码，一些聚焦于单个知识点的示例，一些自己实现的框架。

### 目录里面是什么

- "algorithm/"，算法。
- "cache/"，自己实现的缓存框架。
- "demo/"，聚焦于单个知识点的示例。
- "leetcode/"，leetcode 算法题。
- "orm/"，自己实现的 orm 框架。
- "rpc/"，自己实现的 rpc 协议。
- "web/"，自己实现的 web 框架。

自己实现的缓存框架。

- 主要实现：
    - 缓存框架
    - 缓存框架支持本地缓存作为底层
    - 缓存框架支持 redis 作为底层
    - redis 锁
- 次要实现：
    - 内存限制和 lru 淘汰策略
    - read through
- 计划实现：
    - 其他的缓存模式

自己实现的 rpc 协议。

- 主要实现：
    - 自定义 rpc 协议
    - rpc 客户端
    - rpc 服务端
- 次要实现：
    - 自定义 json 协议
    - json 序列化
- 计划实现：
    - 增加 protobuf 序列化
    - 增加 gzip 压缩

自己实现的 web 框架。

- `v40/`，用 net/http 包，实现一个复杂的 web 框架。
    - 主要实现：
        - 路由树（静态、通配符、路径参数、正则表达式）、路由组
        - 全局中间件、可路由的中间件
    - 次要实现：
        - 内存池（请求上下文对象复用）
        - 服务管理（管理多个子服务）
        - 优雅退出、退出时的回调方法
    - 计划实现：
        - 用户认证（中间件实现）
        - 文件操作（上传、下载）
        - 单元测试、集成测试、性能测试

