## 项目说明

- 一个极简的golang微服务框架



## 功能实现

- 基于cobra pflag viper 构建客户端工具
- 基于Makefile 构建项目
- 内嵌errorx logx封装包
- 内嵌 http 响应码生成工具 codegen



## 版本

- go1.18.1
- GNU Make 4.3



## 使用

- cd tools/codegen && go build codegen.go && mv codegen "your GOBIN directory" // 安装响应码生成工具

- make run.demo // 运行demo示例项目

  ```
  2023-08-14 14:09:12.411 INFO    app/app.go:170  WorkingDir: /root/workspace/go/app/micro-frame
  2023-08-14 14:09:12.411 INFO    demo/run.go:9   {"app":{"name":"demo"},"grpc":{"host":"192.168.60.34","port":9791},"http":{"host":"192.168.60.34","port":8791}}
  ```

  



