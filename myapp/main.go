package main

import (
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"myapp/internal/model/entity"
	_ "myapp/internal/packed"
)

func main() {
	//cmd.Main.Run(gctx.New())
	//configTest()
	//logTest()
	//errTest()
	testORM()
}
func configTest() {
	//单例获取配置文件对象
	fmt.Println(g.Cfg("protest").Get(gctx.New(), "biserver.http.port"))
	fmt.Println(gcfg.Instance("protest").Get(gctx.New(), "biserver.http.port"))
	//修改默认配置文件
	fmt.Println(g.Cfg().Get(gctx.New(), "server.address"))
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("protest.properties")
	fmt.Println(g.Cfg().Get(gctx.New(), "server.address"))
	//内容配置
}
func logTest() {
	g.Log().Info(context.TODO(), "info1")
	l := glog.New()
	l.SetLevel(glog.LEVEL_ALL ^ glog.LEVEL_ERRO)
	l.SetLevelStr("notice")
	g.Log().Info(context.TODO(), "info2")
}
func errTest() {
	fmt.Printf("%v", gerror.New("ss"))
}
func testORM() {
	mod := g.DB().Model("platNumber")

	v, _ := mod.Clone().Limit(10).All()
	fmt.Printf("platNumber:  %v \n", v)
	//链式安全操作
	v, _ = mod.Clone().Safe().Limit(10).All()
	fmt.Printf("platNumber:  %v \n", v)
	//写入
	pn := &entity.Platnumber{
		Platnid:   1212,
		Platn:     "12343444",
		NnameCN:   "test2",
		NnameEN:   "test2",
		Ndescript: " s",
		Nvalid:    1,
	}
	mod.Clone().Data(pn).InsertAndGetId()
	mod.Clone().Data(pn).Insert()
	mod.Clone().Data(pn).InsertIgnore()
	mod.Clone().Data(pn).Replace()
	mod.Clone().Data(pn).OnDuplicate("nnameCN,nnameen").Save()

}
