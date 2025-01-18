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
- 引入`gorm.io/gorm`操作数据库
# 亮点，很好的优雅度
- 封装了db error，redis error, db层减少了很多if error，else error的判断
- 封装了gin框架，减少了控制器层error相关if else的判断，使代码更美观，清晰，简洁
- 封装了db error，redis error,http error,业务error，以及golang自身的error类型，响应不同的错误码结果
- 封装了gin的参数绑定，提供了shouldbindhandle、shouldbindqueryhandle等方法，使controller代码更优雅

### 依赖库如下
```shell
github.com/gin-gonic/gin v1.9.1
github.com/go-playground/locales v0.14.1
github.com/go-playground/universal-translator v0.18.1
github.com/go-playground/validator/v10 v10.14.0
github.com/go-sql-driver/mysql v1.8.1
github.com/google/uuid v1.6.0
github.com/labstack/gommon v0.4.2
github.com/redis/go-redis/v9 v9.5.1
github.com/robfig/cron v1.2.0
github.com/rs/zerolog v1.33.0
golang.org/x/text v0.14.0
gopkg.in/yaml.v3 v3.0.1
gorm.io/driver/mysql v1.5.7
gorm.io/gorm v1.25.10
github.com/go-resty/resty/v2 v2.13.1
```

### 目录结构

- cmd/ - web服务、cron的主入口目录
- config/ -配置文件目录
- const/ -常量目录
- controller/ - 控制器目录
- internal/ -内部功能目录,里面方法不建议修改
- cron/ - 定时任务目录
- middleware/ -中间件目录
- model/ -数据表结构目录
- logic/ -业务逻辑目录
- typeing/ 结构目录，用于定义请求参数、响应的数据结构
- util/ 工具目录，提供常用的辅助函数，一般不包含业务逻辑和状态信息
- event/ 事件目录
    - listener/ 事件监听器
- rest/ 请求第三方服务的目录
- task/ 任务队列目录
- router/ 路由目录
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
    type GetUsersLogic struct {
        model *models.UserModel
    }

    func NewGetUsersLogic() *GetUsersLogic {
        return &GetUsersLogic{
            model: models.NewUserModel(),
        }
    }

    func (l *GetUsersLogic) Handle(ctx context.Context, req types.ListReq) (resp *types.ListReply, err error) {
        var u []models.User
        if u, err = l.model.List(ctx); err != nil {
            return nil, err
        }

        redisx.GetInstance().HSet(ctx, "name", "age", 43)
        return &types.ListReply{
            Users: u,
        }, nil

    }
    
    ```
- 数据库

    要使用数据库，为了记录traceid，以及防止乱调用，所以系统只定义了一种获取gorm连接的方式,必须先调用`WithContext(ctx)`才能获得gorm资源，如下
    ```go
    db.WithContext(ctx).Find(&u).Error()
    ```
    Error()方法底层转成DBError，便于上层区分，以及响应判断
- redis

    系统的redis库用的是`go-redis`,没有进行过多的封装，获取redis连接后，使用方法上就跟`go-redis`一样了,调用`GetInstance()`方法获取redis资源对象
    ```go
    redisx.GetInstance().HSet(ctx, "name", "age", 43)
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
        "go-gin/model"
    )

    type DemoAListener struct {
    }

    func (l DemoAListener) Handle(ctx context.Context, e *eventbus.Event) error {
        user := e.Payload().(*model.User)
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

        var u model.User
        db.WithContext(ctx).Find(&u)

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

    验证器主要是对`gin`内置的binding进行的扩展
    - 支持中文化提示
        ```go
        type AddUserReq struct {
            Name   string    `form:"name" binding:"required"`
            Age    int       `form:"age" binding:"required"`
            Status bool      `form:"status"`
            Ctime  time.Time `form:"ctime"`
        }

        // controller
        var req typeing.AddUserReq
            if err := ctx.ShouldBind(&req); err != nil {
                logx.WithContext(ctx).Warn("ShouldBind异常", err)
                httpx.Error(ctx, err)
                return
            }
        ```
        如上如果参数不包括name的时候，会提示如下,自动进行了中文化处理
        ```
        {"code":10001,"data":null,"message":"Name为必填字段\n年龄为必填字段\n","trace_id":"695517e3-1b68-4845-839d-c0e58d8f3a43"}
    - 支持自定义提示语的字段名字

        使用`label`标签定义字段名字

        ```go
        type AddUserReq struct {
            Name   string    `form:"name" binding:"required"`
            Age    int       `form:"age" binding:"required" label:"年龄"`
            Status bool      `form:"status"`
            Ctime  time.Time `form:"ctime"`
        }
        ```
        如上提示语不再提示`Age为必填字段`，而是提示`年龄为必填字段`

    - 支持非`gin`框架方式使用验证器
        提供了`validators.Validate()`方法进行验证结构字段的值是否合理
        ```go
        var req = typeing.AddUserReq{
            Name: "测试",
        }
        if err := validators.Validate(&req); err != nil {
            httpx.Error(ctx, err)
            return
        }
        ```
        **注意:**`validators.Validate`和`ctx.ShouldBind`验证失败返回的是`BizError`类型错误,错误码是`ErrCodeValidateFailed`,默认值是`10001`，你也可以通过`errorx.ErrCodeValidateFailed = xxx`在main入口修改默认值
- 参数、响应结构

    定义了可以规范化请求参数、响应结构的目录，使代码更容易维护，结构定义在`typeing/`目录，一个模块一个文件名，如`user.go`
    
    结构定义如下
    ```go
        package typeing

        import (
            "time"
        )

        type AddUserReq struct {
            Name   string    `form:"name"`
            Age    int       `form:"age"`
            Status bool      `form:"status"`
            Ctime  time.Time `form:"ctime"`
        }

        type AddUserReply struct {
            Message string `json:"message"`
        }

    ```
    使用方式,在`controller`层使用
    ```go
    var req typeing.AddUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		httpx.Error(ctx, err)
		return
	}

    ```
    其实就是使用了`gin`框架本身提供的shouldbind特性，将参数绑定到结构体，后面逻辑直接可以使用结构体里面的字段进行操作了，参数需要包括那些字段，通过结构体很容易看到，实现了参数的可维护性
    ```go
    resp := typeing.AddUserReply{
		Message: fmt.Sprintf("add user succcess %s=%d", user.Name, user.Id),
	}
	httpx.Ok(ctx, resp)
    ```
    响应结构体如上，结构体数据响应中转成json渲染到`data`域，这样实现相应的结构化和可维护性，响应结果如下
    ```
    {"code":0,"data":{"message":"add user succcess ddddd=125"},"message":"成功","trace_id":"b1a9e4f8-7772-4c3a-bb3d-99a22d6a0ff6"}
    ```
    

- 常量

    未来系统中可能会存在很多业务常量，这里预先建立了目录，当前内置了一些关于错误的预定义常量，这样在业务逻辑中直接使用即可，不需要到处写相同的错误，另外使错误相关更加集中，方便管理，也提高了可维护性
    ```go
    var (
        ErrUserNotFound = errorx.New(2001, "用户不存在")
    )

    ```
- 错误类型
    系统内置了两种错误类型`BizError`和`ServerError`
    - `ServerError`主要是为了处理no method或者method not allowed以及其他服务上的错误，便于响应返回正确的http状态码和统一一致的响应结构,`errorx`包内置错误常量
    ```go
        ErrMethodNotAllowed    = NewServerError(http.StatusMethodNotAllowed)
        ErrNoRoute             = NewServerError(http.StatusNotFound)
        ErrInternalServerError = NewServerError(http.StatusInternalServerError)
    ```
    - `BizError`是我们业务开发中使用更多的错误结构，就是业务中定义的异常错误类型，这种类型返回的http状态码都是200，响应结构的状态码、消息均来源于`BizError`变量中。`BizError`的变量定义方式如下
        ```go
        errorx.New(20001, "用户不存在")
        errorx.NewDefault("用户不存在")  // code默认值为ErrCodeDefaultCommon的值，也就是10000
        ```
        注意，新增的业务错误码建议从20000开始，因为`internal`底层可能会定义10000-20000之内的业务错误码，例如校验失败的错误码是`ErrCodeValidateFailed`值为10001,通用错误`ErrCodeDefaultCommon`值为10000
    - `error`,error应该是其他错误的超类，如果非上述两种错误，我们统一用`error`捕获，并且返回响应http状态码200,code为默认值`ErrCodeDefaultCommon`，也就是10000
        ```
        {
            "code": 10000,
            "data": null,
            "message": "用户不存在",
            "trace_id": "dc119c64-d4b9-4af1-9e02-d15fc4ba2e42"
        }
        ```
- 请求第三方接口
接入了`go-resty`库，并做了简单封装，便于开箱即用
    
    - 原生方式
        ```
        resp, err := httpc.POST(ctx, "http://localhost:8080/api/list").
            SetFormData(httpc.M{"username": "aaaa", "age": "55555"}).
            Send()
        ```
        如上,主要对go-resty进行了简单封装，封装成了`httpc`库,并提供了`POST`,`GET`常用两种请求方式
    - 服务方式
    
        如果第三方接口交互较多，可以作为服务进行对接，首先在`main.go`文件配置第三方服务地址,例如
        ```go
        user.Init("http://localhost:8080")
        ```
        然后在`rest`目录定义服务相关文件主要包括
        - `init.go`启动文件
        - `response.go`接口返回格式以及解析响应结果
        - `svc.go` 定义服务接口、参数以及响应结构，进行明确要求，便于代码的可维护性
        - `svc.impl.go` 对svc.go中接口的实现
        定义要上面几个文件之后，便可以在自己的业务文件中发起请求了
            ```go
            hash := md5.Sum([]byte("abcd"))
            pwd := hex.EncodeToString(hash[:])
            resp, err := login.Svc.Login(ctx, &login.LoginReq{Username: "1", Pwd: pwd})
            if err != nil {
                httpx.Error(ctx, err)
                return
            }
            ```

- 队列

    队列使用的是比较热门的库`github.com/hibiken/asynq`,本项目稍微进行了一点点儿封装，简化使用，更加结构化，便于代码的维护,弱化了client和server端指定taskname
    - 队列server目录为`cmd/task`
    - 队列代码维护在`task/`目录
    - 将数据写入队列的方式，封装了3个方法
        ```go
        task.Dispatch(queue.NewSampleTask("测试3333"),3*time.Secord) // 使用task包下的Dispatch方法,并添加延迟时间3s后执行
        task.DispatchWithRetry(queue.NewSampleTask("测试3333"),)// 使用task包下的Dispatch方法,并添加延迟时间和失败后的重试次数
        task.DispatchNow(queue.NewSampleTask("测试3333")) // 使用task包下的Dispatch方法,立即执行
        tasqueuekx.NewSampleTask("测试3333").DispatchNow() // 使用task结构的DispatchNow方法
        
        task.NewOption().Queue(task.HIGH).TaskID("test").Dispatch(queue.NewSampleBTask("hello")) // 指定发送队列

        ```
    - server端handler处理,首先需要将没一个task的handler维护到server端,在`task/init.go`文件进行添加
        ```go
        task.Handle(NewSampleTaskHandler()) // Handle是封装的一个方法
        ```
### 快速启动

```shell
1. git clone git@github.com:fanqingxuan/go-gin.git
2. cd go-gin && go mod tidy
3. web启动方式 go run cmd/api/main.go  -f .env
4. 定时任务 go run cmd/cron/main.go  -f .env
```
