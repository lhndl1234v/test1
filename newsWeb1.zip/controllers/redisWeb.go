package controllers

import (
	"github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego"

)

func init(){
	conn,err:=redis.Dial("tcp",":6379")
	if err !=nil{
		beego.Error("redis连接失败")
		return
	}
	defer conn.Close()
	//conn.Send("set","redis","test")
	//conn.Flush()
	//conn.Receive()

	resp,err:=conn.Do("mget","t1","t2","t3")

	result,_:=redis.Values(resp,err)

	var v1,v2 int
	var v3 string

	redis.Scan(result,&v1,&v2,&v3)

	beego.Info("获取的数据为:",v1,v2,v3)
}
