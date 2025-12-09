# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Run Commands

```bash
# API server
go run cmd/api/main.go -f .env

# Cron/scheduled tasks
go run cmd/cron/main.go -f .env

# Queue worker
go run cmd/queue/main.go -f .env

# Database migrations
go run cmd/migrate/main.go -f .env

# Run tests
go test ./test/...

# Run a single test
go test ./test/ -run TestIsTrue
```

## Architecture Overview

This is a Gin-based web framework with custom wrappers for cleaner code. Uses Go 1.21+.

### Layered Architecture

```
router/ → controller/ → logic/ → model/
```

- **router/**: Route definitions grouped by module (user.go, login.go, etc.)
- **controller/**: Thin layer returning `(any, error)` - framework handles response formatting
- **logic/**: Business logic using Command pattern - one logic struct per operation (e.g., `GetUsersLogic`)
- **model/**: GORM models with wrapped database access

### Key Conventions

**Controllers return data directly** - no manual JSON response handling:
```go
func (c *userController) List(ctx *httpx.Context) (any, error) {
    return logic.NewGetUsersLogic().Handle(ctx, req)
}
```

**Database access requires context** - always use `db.WithContext(ctx)`:
```go
db.WithContext(ctx).Find(&u).Error()  // Error() returns wrapped DBError
```

**Logging uses context for trace ID**:
```go
logx.WithContext(ctx).Info("keyword", message)
```

### Error Types

- `BizError`: Business errors (HTTP 200, custom code) - use `errorx.New(code, msg)`
- `ServerError`: Server errors (HTTP status code) - use `errorx.NewServerError(status)`
- Business error codes should start from 20000 (10000-20000 reserved for framework)

### Enum Pattern

Enums in `const/enum/` embed `etype.BaseEnum` and require:
1. A prefix constant for type registry
2. Implementation of `sql.Scanner` and `json.Unmarshaler` interfaces
3. Constructor function (e.g., `NewUserStatus(code, desc)`)

### Request/Response Types

Defined in `typing/` directory. Request structs use `form:` tags for binding, `label:` tag for Chinese validation messages.

### Event System

- Events in `event/`, listeners in `event/listener/`
- Register with `eventbus.AddListener(eventName, &Listener{})`
- Fire with `eventbus.Fire(ctx, event.NewSampleEvent(payload))`

### Queue Tasks

- Task handlers in `task/` directory
- Dispatch: `task.DispatchNow(task.NewSampleTask(data))`
- Register handlers in `task/init.go` with `task.Handle(NewSampleTaskHandler())`

### Cron Jobs

- Job structs implement `cronx.Job` interface with `Handle(ctx context.Context) error`
- Register in `cron/init.go`: `cronx.AddJob("@every 3s", &SampleJob{})`
- Also supports `cronx.ScheduleFunc(fn).EveryMinute()` pattern
