package models

import (
	"context"
	"fmt"
	"go-gin/svc/sqlx"

	"github.com/guregu/null/v5"
)

type User struct {
	Name       string    `db:"username" json:"username"`
	Age        int       `json:"age"`
	CreateTime null.Time `json:"ctime"`
}

type UserModel struct {
	ctx     context.Context
	sqlconn sqlx.SqlConn
}

func NewUserModel(ctx context.Context, sqlconn sqlx.SqlConn) UserModel {
	return UserModel{
		ctx:     ctx,
		sqlconn: sqlconn,
	}
}

func (u UserModel) FindAll(id uint64) ([]User, error) {
	var users []User
	err := u.sqlconn.QueryRowsCtx(u.ctx, &users, "select username,age,create_time from user where id>?", id)
	fmt.Println(err)
	return users, err
}
