# CMD 入口命令

项目提供 5 个独立入口，均通过 `-f` 指定配置文件（默认 `.env`）。

## cmd/api - API 服务

```bash
go run cmd/api/main.go -f .env
```

启动流程：`config → logx → db → redisx → queue → event → httpx engine → middleware → router → run`

## cmd/cron - 定时任务服务

```bash
go run cmd/cron/main.go -f .env
```

启动流程：`config → logx → db → redisx → event → cronx.New() → cron.Init() → cronx.Run()`

`cronx.Run()` 会阻塞直到收到 SIGINT/SIGTERM 信号。

## cmd/queue - 队列消费者服务

```bash
go run cmd/queue/main.go -f .env
```

启动流程：`config → logx → db → redisx → queue.InitServer() → task.Init() → queue.Start()`

## cmd/migrate - 数据库迁移

```bash
go run cmd/migrate/main.go -f .env
```

读取配置后连接数据库，执行 `migration/` 目录下注册的迁移文件。

## cmd/make - 代码生成工具集

```bash
go run ./cmd/make/... make:enum                     # 扫描 const/enum/ 生成枚举方法
go run ./cmd/make/... make:dao -f .env              # 生成所有表的 entity/do/dao
go run ./cmd/make/... make:dao -f .env -t user      # 生成指定表
go run ./cmd/make/... make:migration create_orders  # 生成迁移文件
```

注意：必须使用 `./cmd/make/...` 语法。
