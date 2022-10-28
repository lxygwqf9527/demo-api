# Host 服务模块

## IMPL
    Host Service的具体实现,上册业务就基于Service进行编程,而像接口

    ```
    http
     |
    Host Service(interface impl)
     |
    impl(基于MysSQL实现)
    ```

    HostService定义 并把实现变成编写完成，使用方式有多种用途：
    + 用于内部模块调用，基于他实现更高一层的业务逻辑，比如发布服务
    + Host Service对外暴露：http协议(暴露给用户)
    + Host Service对外暴露：Grpc(暴露给内部服务)

## 