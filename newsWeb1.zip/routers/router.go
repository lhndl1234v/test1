package routers

import (
	"newsWeb1.zip/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandleRegister")
    //登录业务处理
    beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
    //首页展示
    beego.Router("/index",&controllers.ArticleController{},"get,post:ShowIndex")
    //添加文章业务
    beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")
    //查看文章详情
    beego.Router("/content",&controllers.ArticleController{},"get:ShowContent")
    //编辑文章
    beego.Router("/update",&controllers.ArticleController{},"get:ShowUpdate;post:HandleUpdate")
    //删除文章
    beego.Router("/delete",&controllers.ArticleController{},"get:HandleDelete")
    //展示添加分类页面
    beego.Router("/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")
}
