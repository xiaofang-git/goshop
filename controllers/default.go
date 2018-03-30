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

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Title"] = "Gshop"
	c.TplName = "login.html"
}