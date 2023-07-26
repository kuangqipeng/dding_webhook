package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// 1.简单钉钉的 webhook接收器
func main() {
	accessToken := "<dingding robot accessToken>"
	secert := "<dingding robot secret>"

	robot := NewRobot(accessToken, secert)
	e := gin.Default()
	tpl := template.Must(template.ParseFiles("./notify.tpl"))
	// 1. 接受到报警通知
	// 2. 解析模板填充数据
	// 3. 发送钉钉
	e.Handle(http.MethodPost, "/webhook", func(c *gin.Context) {
		// 1.接受报警通知
		data := new(Notify)
		err := c.BindJSON(data)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
		}

		// 2.解析模板填充数据
		notify := buildNotifyTpl(data)
		buf := bytes.NewBuffer([]byte{})
		err = tpl.ExecuteTemplate(buf, "alert-notify", notify)
		if err != nil {
			fmt.Printf("err: %+v\n", err)
		}
		err = robot.BuildMsgAndSend(WithMarkDown("告警信息通知", buf.String()))
		if err != nil {
			fmt.Printf("err: %+v\n", err)
		}
		return
	})
	e.Run(":20000")
}

func buildNotifyTpl(data *Notify) NotifyTpl {
	tpl := NotifyTpl{
		CommonAnnotations: data.CommonAnnotations,
		CommonLabels:      data.CommonLabels,
		Status:            data.Status,
	}

	// UTC时区转换
	local, _ := time.LoadLocation("Local")
	startTime := data.Alerts[0].StartsAt.In(local).Format(time.DateTime)

	// instance获取
	ins := data.GroupLabels["instance"]
	instance := strings.Split(ins, ":")[0]

	// alertName获取
	alertName := data.CommonLabels["alertname"].(string)

	tpl.AlertNotifyName = alertName
	tpl.Instance = instance
	tpl.AlertTime = startTime
	return tpl
}
