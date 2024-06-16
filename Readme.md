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
```

### 目录结构

- cmd/ - web服务、cron的主入口目录
- config/ -配置文件目录
- consts/ -常量目录
- controllers/ - 控制器目录
- internal/ -内部功能目录,里面方法不建议修改
- jobs/ - 定时任务目录
- middlewares/ -中间件目录
- models/ -数据表结构目录
- services/ -业务逻辑目录
- types/ 结构目录，用于定义请求参数、响应的数据结构
- utils/ 工具目录，提供常用的辅助函数，一般不包含业务逻辑和状态信息

### 功能代码
- 控制器

    在`controllers`目录下面创建控制器，例如`user_controller.go`
    ```go
    type userController struct {
    }

    var UserController = &userController{
    }

    func (c *userController) Index(ctx *gin.Context) {
        httpx.Ok(ctx, "hello world")
    }
    ```
    然后在`controllers/init.go`文件定义路由即可
    ```go
    user_router := route.Group("/user")
	user_router.GET("/", UserController.Index)
    ```
    另外，对于控制器的响应封装了几个公共方法
    ```go
    httpx.Ok(ctx, "hello world") // 输出正常的响应
    httpx.OkWithMessage(ctx *gin.Context, data any, msg string)
    
    httpx.Error(ctx, err) //输出异常的响应

    httpx.Handle(ctx *gin.Context, data any, err error) //既可以输出正常的响应，又可以说出异常的响应
    ```
    封装响应的原因是定义了输出的响应结构，如下，永远返回包含code、data、message、trace_id四个字段的结构，使相应结果结构化
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
    如果相应结构字段名字不符合你的预期，可以进行自定义
    ```go
    func main() {
        // to do something
        httpx.DefaultSuccessCodeValue = 0 // 定义成功的code默认值,默认是0，你也可以改成200
        httpx.CodeFieldName = "code" // 定义响应结构的code字段名，你也可以改成status
        httpx.MessageFieldName="msg"// 定义响应结构的消息字段名
        httpx.ResultFieldName = "data"// 定义响应结构的数据字段名
        traceid.TraceIdFieldName="request_id" // 定义响应以及日志中traceid的字段名字
    }
    ```
- 服务层

    服务层代码没有什么特别的，需要说明的是方法的第一个参数建议是`context.Context`,一是统一规范，而是可以日志记录traceid
    ```go
    type UserService struct {
    }

    func NewUserService() *UserService {
        return &UserService{}
    }

    func (svc *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
        var u []models.User
        if err := db.WithContext(ctx).Find(&u).Error; err != nil {
            return nil, err
        }
        return u, nil

    }
    ```
- 数据库

    要使用数据库，为了记录traceid，以及防止乱调用，所以系统只定义了一种获取gorm连接的方式,必须先调用`WithContext(ctx)`才能获得gorm资源，如下
    ```go
    db.WithContext(ctx).Find(&u).Error
    ```
- redis

    系统的redis库用的是`go-redis`,没有进行过多的封装，所以获取redis连接后，使用方式参考`go-redis`手册，获取`go-redis`方式如下,需要调用`GetInstance()`方法
    ```go
    redisx.GetInstance().HSet(ctx, "name", "age", 43)
    ```
- 日志
- 定时任务
- 验证器
- 参数、响应结构
- 常量
### 快速启动

```shell
1. git clone git@github.com:fanqingxuan/go-gin.git
2. cd go-gin && go mod tidy
3. web启动方式 go run cmd/api/main.go  -f .env
4. 定时任务 go run cmd/cron/main.go  -f .env
```