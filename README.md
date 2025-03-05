# 介绍

快速创建模板项目的脚本命令行程序，根据 git 代码仓库，快速将仓库拉下来，并根据项目类型，完成对应的初始化操作。

# 安装

脚本已同时发布到 npm，在安装有 nodejs 的环境下，可以通过 npx 使用

```bash
npx pqc -t react test
```

或者全局安装

```bash
npm install -g pqc
yarn global add pqc@latest
```

目前 npm 上支持的平台仅为 macos，其他平台请自行编译。

# 编译

```bash
go build -o pcq main.go
```

# 执行

```bash
pcq -t go github.com/ihezbien/test
#pcq -t vite test
```

```bash
hezebin@ ~ go run main.go -t go github.com/ihezbien/test

Project name: test, Mod name: github.com/ihezbien/test

Generating project Success!

Organizing project files...
[Success]  test/.gitignore
[Success]  test/README.md
[Success]  test/application/test.go
[Success]  test/cmd/root.go
[Success]  test/component/cache/memory.go
[Success]  test/component/cache/redis.go
[Success]  test/component/constant/commom.go
[Success]  test/component/doc/doc.go
[Success]  test/component/doc/swagger.json
[Success]  test/component/email/email.go
[Success]  test/component/pubsub/pulsar.go
[Success]  test/component/sms/sms.go
[Success]  test/component/storage/mongo.go
[Success]  test/component/storage/mysql.go
[Success]  test/config/config.go
[Success]  test/config/config.json
[Success]  test/config/config.toml
[Success]  test/domain/entity/test.go
[Success]  test/domain/repository/impl/mongo/test.go
[Success]  test/domain/repository/impl/redis/test.go
[Success]  test/domain/repository/test.go
[Success]  test/domain/service/test.go
[Success]  test/go.mod
[Success]  test/main.go
[Success]  test/script/test.js
[Success]  test/script/test.py
[Success]  test/server/dto/test/test.go
[Success]  test/server/handler/test.go
[Success]  test/server/middleware/cors.go
[Success]  test/server/server.go
[Success]  test/static/img.png
[Success]  test/worker/timer.go

Init project success!

Now: cd test

```

## 项目类型

不同项目类型会做不同的初始化处理，如 Go 项目会根据新项目名称重命名模板中的 module 名称，React 项目会根据新项目名称重命名模板中的 package.json 中的 name 字段等。

项目类型可以通过 `-t` 或 `-template` 参数来指定，该参数内置支持的项目类型。详情可通过`help`命令查看参数描述。

## 模板项目资源

本脚本应用构建原理依赖已有 git 模板代码仓库，若不想使用默认的代码仓库，可以通过`origin` 或 `-o` 来指定。
