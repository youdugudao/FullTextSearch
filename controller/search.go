package controller

import (
	"github.com/gin-gonic/gin"
	"full_text_search/search"
	"full_text_search/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func Get(c *gin.Context) {

	ResponseOk("get", c)
}

func Post(c *gin.Context) {
	var requestData struct {
		Content string
	}
	if err := c.BindJSON(&requestData); err != nil {
		ResponseError("解析request json失败， 原因："+err.Error(), c, nil)
		return
	}
	if err := search.SetIndex([]byte(requestData.Content)); err != nil {
		ResponseError("索引添加失败， 原因："+err.Error(), c, nil)
		return
	}
	ResponseOk(nil, c)
}

func Put(c *gin.Context) {

	ResponseOk("put", c)
}

func Delete(c *gin.Context) {

	ResponseOk("delete", c)
}

func Test() {
	// 为了代码的可读性，没有验证错误error
	database.Collection("user", func(collection *mgo.Collection) {
		changeLog, err := collection.Upsert(bson.M{"x": 11}, bson.M{"$set": bson.M{"y": 1}})
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Matched:", changeLog.Matched)
		fmt.Println("Removed:", changeLog.Removed)
		fmt.Println("Updated:", changeLog.Updated)
		fmt.Println("UpsertedId:", changeLog.UpsertedId)
	})
}
