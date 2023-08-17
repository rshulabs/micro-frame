## 项目说明

- 一个极简的 golang 微服务框架

## 功能实现

- 基于 cobra pflag viper 构建客户端工具
- 基于 Makefile 构建项目
- 内嵌 errorx logx 封装包
- 内嵌 http 响应码生成工具 codegen

## 版本

- go1.20.3 | > 19
- GNU Make 4.3

## 使用

- cd tools/codegen && go build codegen.go && mv codegen "your GOBIN directory" // 安装响应码生成工具

  make codegen // 执行 codegen 工具

  ````
  # 错误码

  !! demo 系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。

  ## 功能说明

  如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：

  ```json
  {
    "code": 100101,
    "message": "Database error"
  }
  ```

  上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

  ## 错误码列表

  支持的错误码列表如下：

  | Identifier | Code | HTTP Code | Description |
  | ---------- | ---- | --------- | ----------- |
  | ErrSuccess | 100001 | 200 | OK |
  | ErrUnknown | 100002 | 500 | Internal server error |
  | ErrValidation | 100003 | 400 | Validation failed |
  | ErrNotFound | 100004 | 404 | Page not found |
  | ErrTokenInvalid | 100005 | 401 | Token invalid |
  | ErrBind | 100006 | 400 | Error occurred while binding the request body to the struct |
  | ErrDatabase | 100101 | 500 | Database error |
  | ErrEncrypt | 100201 | 401 | Error occurred while encrypting the user password |
  | ErrSignatureInvalid | 100202 | 401 | Signature invalid |
  | ErrInvalidAuthHeader | 100203 | 401 | Invalid authorization header |
  | ErrMissingHeader | 100204 | 401 | Missing authorization header |
  | ErrPasswordIncorrect | 100205 | 401 | Password incorrect |
  | ErrPermissionDenied | 100206 | 403 | Permission denied |
  | ErrBlackListCheck | 100207 | 401 | Black list check failed |
  | ErrGuardTokenCheck | 100208 | 401 | Guard token check failed |
  | ErrEncodingFailed | 100301 | 500 | Encoding failed due to an error with the data |
  | ErrDecodingFailed | 100302 | 500 | Decoding failed due to an error with the data |
  | ErrInvalidJSON | 100303 | 500 | Invalid json data |
  | ErrEncodingJSON | 100304 | 500 | JSON data could not be encoded |
  | ErrDecodingJSON | 100305 | 500 | JSON data could not be decoded |
  | ErrInvalidYaml | 100306 | 500 | Invalid yaml data |
  | ErrEncodingYaml | 100307 | 500 | YAML data could not be encoded |
  | ErrDecodingYAML | 100308 | 500 | YAML data could not be decoded |
  ````

- make run.demo // 运行 demo 示例项目

  ```
  2023-08-14 14:09:12.411 INFO    app/app.go:170  WorkingDir: /root/workspace/go/app/micro-frame
  2023-08-14 14:09:12.411 INFO    demo/run.go:9   {"app":{"name":"demo"},"grpc":{"host":"192.168.60.34","port":9791},"http":{"host":"192.168.60.34","port":8791}}
  ```

- make build // 编译所有 main.go 文件
