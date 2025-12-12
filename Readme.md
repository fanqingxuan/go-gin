# go-gin

用gin框架配合golang方面比较优秀的库，搭建的一个项目结构，方便快速开发项目。出结果
用最少的依赖实现80%项目可以完成的需求

# 功能特性
- 使用主流轻量的路由框架gin，实现路由
- 引入`github.com/go-playground/validator`实现常见的验证，最重要的是引入了中文的提示，以及可以自定义字段名字
- 引入主流的`gorm`库作为数据库层的操作
- 引入`github.com/redis/go-redis/v9`作为缓存层操作
- 引入`github.com/google/uuid`生成traceid，traceid贯穿于各种日志，以及在响应中返回，并且支持自定义traceid的字段名字
- 引入`github.com/labstack/gommon`实现调试模式下日志打印到console，并且不同的日志级别用不用的颜色进行区分
- 引入`github.com/robfig/cron`实现定时任务，定时任务也引入了traceid
- 使用轻量的日志库`github.com/rs/zerolog`进行记录日志
- 引入`gopkg.in/yaml.v3`解析yaml配置文件到golang变量
- 引入`github.com/go-resty/resty/v2`发起http请求，方便的请求第三方接口
- 引入`github.com/hibiken/asynq`实现异步队列
- 引入`github.com/hibiken/asynqmon` 队列使用情况的漂亮web ui界面
- 引入`gorm.io/gorm`操作数据库
# 亮点，很好的优雅度
- 封装了db error，redis error, db层减少了很多if error，else error的判断
- 封装了gin框架，减少了控制器层error相关if else的判断，使代码更美观，清晰，简洁
- 封装了gin的参数绑定，提供了shouldbindhandle、shouldbindqueryhandle等方法，使controller代码更优雅
- crontab、task、event、数据库迁移migration可以优雅的使用
- 提供了状态相关的枚举结构
- 日志记录traceid,方便排查问题
### 依赖库如下
```shell
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.14.0
	github.com/go-resty/resty/v2 v2.13.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang-module/carbon/v2 v2.3.12
	github.com/google/uuid v1.6.0
	github.com/hibiken/asynq v0.24.1
	github.com/hibiken/asynqmon v0.7.2
	github.com/labstack/gommon v0.4.2
	github.com/redis/go-redis/v9 v9.7.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/zerolog v1.33.0
	golang.org/x/text v0.16.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.10
```

### 目录结构

- cmd/ - web服务、cron的主入口目录
- config/ -配置文件目录
- const/ -常量目录
  - errcode - 错误结构
  - enum - 枚举常量结构
- controller/ - 控制器目录
- internal/ -内部功能目录,里面方法不建议修改
- cron/ - 定时任务目录
- middleware/ -中间件目录
- model/ -数据层目录
  - entity/ - 表结构映射（强类型，GORM models）
  - do/ - Data Object（any 类型，用于 Where/Data 条件构建）
  - dao/ - Data Access Object（单例模式，链式调用 API）
- logic/ -业务逻辑目录
- typing/ 结构目录，用于定义请求参数、响应的数据结构
- util/ 工具目录，提供常用的辅助函数，一般不包含业务逻辑和状态信息
- event/ 事件目录
    - listener/ 事件监听器
- rest/ 请求第三方服务的目录
- task/ 任务队列目录
- router/ 路由目录
- migration/ 数据库迁移目录
### 功能代码
- 控制器

    在`controller`目录下面创建控制器，例如`user_controller.go`
    ```go
    type userController struct {
    }

    var UserController = &userController{
    }

    func (c *userController) List(ctx *httpx.Context) (any, error) {
        var req types.ListReq
        l := logic.NewGetUsersLogic()
        return l.Handle(ctx, req)
    }
    ```
    然后在`controllers/init.go`文件定义路由即可
    ```go
    user_router := route.Group("/user")
	user_router.GET("/", UserController.Index)
    ```
    控制器直接返回logic层处理后的结果，不需要关心响应格式，减少了不必要的if，else判断，自己封装的gin框架底层会根据error自动判断渲染数据还是error数据
    ```
    封装响应的原因是定义了输出的响应结构，如下，永远返回包含code、data、message、trace_id四个字段的结构，使响应结果结构化
    ```shell
    {
        "code": 0,
        "data": {
            "data": "add user succcess ddddd=96"
        },
        "message": "操作成功",
        "trace_id": "dc119c64-d4b9-4af1-9e02-d15fc4ba2e42"
    }
    ```
    如果响应结构字段名字不符合你的预期，可以进行自定义
    ```go
    func main() {
        // to do something
        httpx.DefaultSuccessCodeValue = 0 // 定义成功的code默认值,默认是0，你也可以改成200
        httpx.DefaultSuccessMessageValue = "成功" // 定义成功的message默认值,默认是'操作成功'
        httpx.CodeFieldName = "code" // 定义响应结构的code字段名，你也可以改成status
        httpx.MessageFieldName="msg"// 定义响应结构的消息字段名
        httpx.ResultFieldName = "data"// 定义响应结构的数据字段名
        traceid.TraceIdFieldName="request_id" // 定义响应以及日志中traceid的字段名字
    }
    ```
    响应结果类似如下
    ```
    {
        "code": 10001,
        "data": null,
        "msg": "年龄为必填字段\n",
        "request_id": "8ddb97db-be44-4df0-8110-0d38a0cc4657"
    }
    ```
- 业务层

    业务层采用命令模式，一个logic只负责处理一个业务的处理，例如`getusers_logic.go`
    ```go
    type GetUsersLogic struct{}

    func NewGetUsersLogic() *GetUsersLogic {
        return &GetUsersLogic{}
    }

    func (l *GetUsersLogic) Handle(ctx context.Context, req typing.ListReq) (*typing.ListReply, error) {
        var users []entity.User
        err := dao.User.Ctx(ctx).
            Where(do.User{Status: 1}).
            Order("id DESC").
            All(&users)
        if err != nil {
            return nil, err
        }

        return &typing.ListReply{Users: users}, nil
    }
    ```
- 数据库

    推荐使用 GoFrame 风格的 DAO 链式调用 API：
    ```go
    // 查询
    var users []entity.User
    err := dao.User.Ctx(ctx).
        Where(do.User{Status: 1}).
        Where("age > ?", 18).
        Order("id DESC").
        Page(1, 10).
        All(&users)

    // 插入
    _, err := dao.User.Ctx(ctx).
        Data(do.User{Name: "test", Status: 1}).
        Insert()

    // 更新
    _, err := dao.User.Ctx(ctx).
        Data(do.User{Status: 2}).
        Where(do.User{Id: 1}).
        Update()

    // 删除
    _, err := dao.User.Ctx(ctx).
        Where(do.User{Id: 1}).
        Delete()

    // 使用 Columns 常量避免硬编码字段名
    cols := dao.User.Columns()
    err := dao.User.Ctx(ctx).
        Fields(cols.Id, cols.Name).
        Where(cols.Status+" = ?", 1).
        All(&users)
    ```

    底层仍可使用 `db.WithContext(ctx)` 直接操作 GORM（仅在 DAO 内部使用）：
    ```go
    db.WithContext(ctx).Find(&u).Error()
    ```
    Error()方法底层转成DBError，便于上层区分，以及响应判断
- redis

    系统的redis库用的是`go-redis`,没有进行过多的封装，获取redis连接后，使用方法上就跟`go-redis`一样了,调用`Client()`方法获取redis资源对象
    ```go
    redisx.Client().HSet(ctx, "name", "age", 43)
    ```
- 日志

    系统提供了debug、info、warn、error四种级别的日志，接口如下
    ```go
    type Logger interface {
        Debug(keyword string, message any)
        Debugf(keyword string, format string, message ...any)

        Info(keyword string, message any)
        Infof(keyword string, format string, message ...any)

        Warn(keyword string, message any)
        Warnf(keyword string, format string, message ...any)

        Error(keyword string, message any)
        Errorf(keyword string, format string, message ...any)
    }
    ```
    可以通过env文件指定日志存储路径和要记录的日志级别，使用方式如下，第一个参数是用于为要记录的日志起一个有意义的关键字，便于grep日志
    ```go
    logx.WithContext(ctx).Warn("ShouldBind异常", err)
    logx.WithContext(ctx).Warnf("这是日志%s", "我叫张三")
    ```
    最终日志文件中记录的内容如下格式，包含`trace_id`
    ```
    {"level":"WARN","keyword":"redis","data":"services/user_service.go:24 execute command:[hset name age 43], error=dial tcp 192.168.65.254:6379: connect: connection refused","time":"2024-06-22 23:24:10","trace_id":"5f8b1ee9-7daf-4269-806a-029ee7c3768f"}
    ```
    另外，常规日志文件的名字是`年-月-日.log`格式，如**2024-05-22.log**。值得注意的是warn、error级别日志会单独拿到`年-月-日-error.log`格式文件，如**2024-05-22-error.log**，这样一方面是便于很好的监控异常，另一方面可以很快的排查异常问题
    
    >此外，系统还提供记录请求access日志，会记录到env配置的路径下的access文件夹,文件以`年-月-日.log`格式命名，日志内容主要包含请求路径、get参数、请求Method、响应码、耗时、User-Agent几个重要参数,格式如下
    ```
    {"level":"INFO","path":"/user/list","method":"GET","ip":"127.0.0.1","cost":"227.238215ms","status":200,"proto":"HTTP/1.1","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36","time":"2024-06-22 23:28:20","trace_id":"f606d909-2f4c-4455-b4b9-5eea0684c49a"}
    ```

- 事件
  事件定义在`event/`目录，事件的监听器定义在`event/listener/`目录
  事件参考内容如下,使用`eventbus.NewEvent`方法注册事件
  ```golang
  package event

    import (
        "go-gin/internal/eventbus"
    )

    var SampleEventName = "event.sample"

    func NewSampleEvent(user string) *eventbus.Event {
        return eventbus.NewEvent(SampleEventName, user)
    }
  ```
  事件监听器`listener`参考内容如下,只需要实现`Handle`方法

  ```golang
    package listener
    import (
        "context"
        "fmt"
        "go-gin/internal/eventbus"
        "go-gin/model/entity"
    )

    type DemoAListener struct {
    }

    func (l DemoAListener) Handle(ctx context.Context, e *eventbus.Event) error {
        user := e.Payload().(*entity.User)
        fmt.Println(user.Name)
        return nil
    }
  ```
  在`event/init.go`文件进行事件与监听器的绑定，使用`eventbus.AddListener`方法`进行绑定,一个事件可以有多个监听器,如果前一个监听器返回error,后面的监听器不会执行
  ```golang
  	eventbus.AddListener(SampleEventName, &listener.SampleAListener{}, &listener.SampleBListener{})
  ```
  在控制器触发事件，触发方式
  ```golang
  eventbus.Fire(ctx, event.NewSampleEvent("hello 测试"))
    // 或者
    event.NewSampleEvent("333").Fire(ctx)
  ```
- 定时任务

    定时任务的入口文件为`cmd/cron/main.go`,具体业务代码在`cron`目录编写。定时任务业务代码可以像api模式一样使用`log`、`db`

    定义一个job首先要定义一个实现了`cronx.Job`的接口的结构,`cronx.Job`接口如下
    ```go
    type Job interface {
        Handle(ctx context.Context) error // 实现业务逻辑
    }
    ```
    例子如下
    ```go
    type SampleJob struct{}

    func (j *SampleJob) Handle(ctx context.Context) error {
        var u entity.User
        dao.User.Ctx(ctx).Where(do.User{Id: 1}).One(&u)
        return nil
    }
    ```
    然后在`cron/init.go`文件定义cron的任务执行频率即可，如下定义`SampleJob`每3s执行一次
    ```
    cronx.AddJob("@every 3s", &SampleJob{})
    // cronx.Schedule(&SampleJob{}).EveryMinute()
    // cronx.Schedule(&SampleJob{}).EveryFiveMinutes()
    ```
    提供了`EveryFiveMinutes`、`EveryThreeMinutes`、`EveryTenMinutes`等多种优雅的时候方法
    定时任务也支持func方式，只需要提供一个`cronx.JobFunc`类型的函数即可，也就是`func(context.Context) error`形式
    例子如下:
    ```go
    func SampleFunc(ctx context.Context) error {
        fmt.Println("this is a sample function")
        return nil
    }
    ```
    我只需要像注册结构体方式一样，将func添加的定时任务管理器即可
    ```go
    cronx.AddFunc("@every 5s", SampleFunc)
	cronx.ScheduleFunc(SampleFunc).EveryMinute()
    ```
    定时任务其它执行频率的定义方式可以参考[https://github.com/robfig/cron](https://github.com/robfig/cron)

- 验证器

    验证器主要是对`gin`内置的binding进行的扩展，支持中文化提示和自定义字段名
    ```go
    type AddUserReq struct {
        Name   string    `form:"name" binding:"required"`
        Age    int       `form:"age" binding:"required" label:"年龄"`
        Status bool      `form:"status"`
    }
    ```
    使用`label`标签定义字段名字，验证失败时提示`年龄为必填字段`而不是`Age为必填字段`

    验证失败返回`BizError`类型错误，错误码默认`10001`
- 参数、响应结构

    请求参数和响应结构定义在`typing/`目录，一个模块一个文件名，如`user.go`
    ```go
    type AddUserReq struct {
        Name   string `form:"name" binding:"required"`
        Age    int    `form:"age" binding:"required" label:"年龄"`
    }

    type AddUserReply struct {
        Message string `json:"message"`
    }
    ```
    Controller 使用 `httpx.ShouldBindHandle` 自动绑定参数并调用 Logic：
    ```go
    func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
        return httpx.ShouldBindHandle(ctx, logic.NewAddUserLogic())
    }
    ```

- 常量

    业务错误常量定义在`const/errcode`目录：
    ```go
    var ErrUserNotFound = errorx.New(20001, "用户不存在")
    ```
- 错误类型
    - `BizError`: 业务错误，HTTP 200，自定义 code - `errorx.New(20001, "用户不存在")`
    - `ServerError`: 服务错误，返回对应 HTTP 状态码 - `errorx.NewServerError(http.StatusNotFound)`
    - 业务错误码建议从 20000 开始（10000-20000 保留给框架）
- 枚举常量
  - 枚举常量定义在 `const/enum` 目录
  - 定义枚举只需两步：
    1. 定义结构体和常量：
        ```golang
        // const/enum/user_status.go
        package enum

        import "go-gin/internal/etype"

        type UserStatus struct {
            etype.BaseEnum
        }

        var (
            USER_STATUS_NORMAL   = etype.NewEnum[UserStatus](1, "正常")
            USER_STATUS_DISABLED = etype.NewEnum[UserStatus](2, "禁用")
        )
        ```
    2. 运行代码生成器自动生成 `gen_*.go` 文件（包含 Scan/Value/JSON 方法）：
        ```bash
        go run ./cmd/make/... make:enum
        ```
  - 使用示例：
    ```golang
    // Entity 中使用枚举
    type User struct {
        Id     int64            `gorm:"column:id" json:"id"`
        Name   string           `gorm:"column:name" json:"name"`
        Status *enum.UserStatus `gorm:"column:status" json:"status"`
    }

    // 创建记录
    user := entity.User{Name: "test", Status: enum.USER_STATUS_NORMAL}
    dao.User.Ctx(ctx).Data(&user).Insert()

    // 查询时自动解析枚举
    var u entity.User
    dao.User.Ctx(ctx).Where(do.User{Id: 1}).One(&u)
    fmt.Println(u.Status.Desc()) // 输出: 正常

    // JSON 序列化只输出 code 值
    // {"id":1,"name":"test","status":1}
    ```
- 请求第三方接口

    使用`httpc`库（基于 go-resty 封装）：
    ```go
    resp, err := httpc.POST(ctx, "http://localhost:8080/api/list").
        SetFormData(httpc.M{"username": "test"}).
        Send()
    ```
    复杂的第三方服务可在`rest/`目录定义服务接口

- 队列

    队列使用`github.com/hibiken/asynq`库，队列服务入口为`cmd/queue`，任务代码维护在`task/`目录
    ```go
    // 发送任务
    task.DispatchNow(task.NewSampleTask("测试"))           // 立即执行
    task.Dispatch(task.NewSampleTask("测试"), 3*time.Second) // 延迟3s执行

    // 在 task/init.go 注册 handler
    task.Handle(NewSampleTaskHandler())
    ```
- util方法
  - `util.IsTrue/IsFalse` - 标量判断
  - `util.When(condition, trueValue, falseValue)` - 三元表达式
  - `jsonx.Encode/Decode` - JSON 序列化/反序列化

### 快速启动

```shell
1. git clone git@github.com:fanqingxuan/go-gin.git
2. cd go-gin && go mod tidy
3. go run cmd/api/main.go  -f .env   // api启动方式
4. go run cmd/cron/main.go  -f .env   // 定时任务启动方式
5. go run cmd/queue/main.go -f .env // 队列服务入口
6. go run cmd/migrate/main.go -f .env // 数据库迁移入口
```

### 代码生成工具

统一的代码生成工具 `cmd/make/`，注意必须使用 `./cmd/make/...` 语法：

```shell
# 生成枚举方法 (扫描 const/enum/ 生成 gen_*.go)
go run ./cmd/make/... make:enum

# 生成 entity/do/dao 三层代码 (从数据库表结构)
go run ./cmd/make/... make:dao -f .env
go run ./cmd/make/... make:dao -f .env -t user,order  # 指定表

# 生成迁移文件模板
go run ./cmd/make/... make:migration create_orders
```
