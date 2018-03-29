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
}



// 商品详情页
type ShopController struct {
	beego.Controller
}


func (c *ShopController) Get() {
	c.TplName = "shop.html"
}



type SingleController struct {
	beego.Controller
}

func (c *SingleController) Get() {
	c.TplName = "single-product.html"
}


type CartController struct {
	beego.Controller
}

func (c *CartController) Get() {
	c.TplName = "cart.html"
}


type CheckoutController struct {
	beego.Controller
}

func (c *CheckoutController) Get() {
	c.TplName = "checkout.html"
}