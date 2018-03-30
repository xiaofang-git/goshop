package routers

import (
	"goshop/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})
	beego.Router("/shop.html", &controllers.ShopController{})
	beego.Router("/single-product.html", &controllers.SingleController{})
	beego.Router("/cart.html", &controllers.CartController{})
	beego.Router("/checkout.html", &controllers.CheckoutController{})
	beego.Router("/login.html", &controllers.LoginController{})
	beego.Router("/user.html", &controllers.UserController{})
	
}
