// @APIVersion 1.0.0
// @Title COG token API
// @Description Kênh thanh toán COG token
// @Contact phanvanthanh.mrt@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/cog",
			beego.NSInclude(
				&controllers.CogController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
