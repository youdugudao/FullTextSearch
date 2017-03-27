package database

import (
	"full_text_search/config"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func init() {
	if Session == nil {
		var err error
		Session, err = mgo.Dial(config.MongoDB.Host + ":" + config.MongoDB.Port)
		if err != nil {
			panic("连接mongodb失败，原因：" + err.Error())
		}
		err = Session.DB("admin").Login(config.MongoDB.User, config.MongoDB.Password)
		if err != nil{
			panic("登录mongodb失败，原因：" + err.Error())
		}
	}
}

func Collection(collectionName string, f func(collection *mgo.Collection)) (err error) {
	session := Session.Clone()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			return
		}
	}()
	collection := session.DB(config.MongoDB.Database).C(collectionName)
	f(collection)
	return
}
