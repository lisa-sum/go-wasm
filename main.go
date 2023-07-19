package main

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

// User 表示一个用户
type User struct {
	ID   int
	Name string
}

// users 保存所有用户
var users = make(map[int]User)

// idCounter 用于为新用户生成 ID
var idCounter = 0

func main() {
	// 启动 Vecty 并渲染 UserList 组件
	vecty.RenderBody(&UserList{})
}

// UserList 是一个 Vecty 组件，它显示一个用户列表和一个添加新用户的表单
type UserList struct {
	vecty.Core
	NewUserName string
}

// Render 渲染 UserList
func (u *UserList) Render() vecty.ComponentOrHTML {
	// 创建一个 vecty.List 来保存所有用户的 UserView
	userViews := vecty.List{}
	for id := range users {
		// 使用当前用户的 ID 创建一个新的 UserView
		userView := &UserView{ID: id}
		userViews = append(userViews, userView)
	}

	// 返回 UserList 的 HTML
	return elem.Body(
		// 显示所有用户
		userViews,

		// 添加新用户的表单
		elem.Form(
			elem.Input(
				vecty.Markup(
					vecty.Property("value", u.NewUserName),
					vecty.Property("placeholder", "New user name"),
					event.Input(func(e *vecty.Event) {
						// 更新 NewUserName 属性当输入框内容改变
						u.NewUserName = e.Value.Get("value").String()
						vecty.Rerender(u)
					}),
				),
			),
			elem.Button(
				vecty.Text("Add user"),
				vecty.Markup(
					event.Click(func(e *vecty.Event) {
						// 防止表单提交
						e.Call("preventDefault")

						// 添加新用户
						users[idCounter] = User{ID: idCounter, Name: u.NewUserName}
						idCounter++

						// 清除输入框
						u.NewUserName = ""
						vecty.Rerender(u)
					}),
				),
			),
		),
	)
}

// UserView 是一个 Vecty 组件，它显示一个用户和一个删除按钮
type UserView struct {
	vecty.Core
	ID int
}

// Render 渲染 UserView
func (u *UserView) Render() vecty.ComponentOrHTML {
	// 获取用户
	user := users[u.ID]

	// 返回 UserView 的 HTML
	return elem.Div(
		elem.Span(
			vecty.Text(strconv.Itoa(user.ID)+": "+user.Name),
		),
		elem.Button(
			vecty.Text("Delete"),
			vecty.Markup(
				event.Click(func(e *vecty.Event) {
					// 删除用户
					delete(users, u.ID)

					// 重新渲染 UserList
					vecty.Rerender(&UserList{})
				}),
			),
		),
	)
}
