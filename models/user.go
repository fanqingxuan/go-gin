package models

import (
	"context"
	"fmt"
	"go-gin/svc/sqlx"
)

type User struct {
	Name string `db:"username" json:"username"`
	Age  int    `db:"age" json:"age"`
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
	err := u.sqlconn.QueryRowsCtx(u.ctx, &users, "select username,age from user where id>?", id)
	fmt.Println(users, err)
	return users, err
}
