package controllers

import (
	"github.com/astaxie/beego"
)


// 主页
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Data["Title"] = "Gshop"
}


// 商品详情页
type ShopController struct {
	beego.Controller
}


func (c *ShopController) Get() {
	c.Data["Title"] = "Gshop"
	c.TplName = "shop.html"
}



type SingleController struct {
	beego.Controller
}

func (c *SingleController) Get() {
	c.Data["Title"] = "Gshop"
	c.TplName = "single-product.html"
}


type CartController struct {
	beego.Controller
}

func (c *CartController) Get() {
	c.Data["Title"] = "Gshop"
	c.TplName = "cart.html"
}


type CheckoutController struct {
	beego.Controller
}

func (c *CheckoutController) Get() {
	c.Data["Title"] = "Gshop"
	c.TplName = "checkout.html"
}


// 登录相关处理逻辑
type LoginController struct {
	beego.Controller
}

// get
func (c *LoginController) Get() {

		// 检查cookie
	name := c.Ctx.GetCookie("name")

	if name == "xiaofang@163.com" {
		// cookie验证成功
		c.TplName = "user.html"
	}else{
		// 验证失败
		c.TplName = "login.html"
	}
}


// post
func (c *LoginController) Post() {
	username := c.GetString("username")
	passwd := c.GetString("passwd")
	// 判断用户名密码是否正确
	if username == "xiaofang@163.com" && passwd == "0805" {
		// c.Ctx.WriteString("登录成功")

		// 设置cookie
		c.Ctx.SetCookie("name", username)
		// c.TplName = "user_info.html"
		c.Ctx.Redirect(302, "user.html")
	}else{
		c.Ctx.WriteString("用户名或密码错误")
	}
}


// 用户信息
type UserController struct {
	beego.Controller
}

// get
func (c *UserController) Get() {
	// 检查cookie
	name := c.Ctx.GetCookie("name")

	if name == "xiaofang@163.com" {
		// cookie验证成功
		c.TplName = "user.html"
	}else{
		// 验证失败
		c.Ctx.Redirect(302, "login.html")
	}
}