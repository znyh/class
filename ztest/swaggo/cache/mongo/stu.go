package mongo

import (
	"swaggo/db/student"
	"swaggo/util"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	mgoSess *mgo.Session
)

func MongdbInit(addr string) error {
	var err error
	mgoSess, err = mgo.Dial(addr)
	return err
}

func QueryStudentInfo(uid uint64) (info *student.StudentInfo, err error) {
	info = &student.StudentInfo{}
	result := bson.M{}
	err = mgoSess.DB("").C("infos").Find(bson.M{"uid": uid}).One(result)
	if err != nil {
		return nil, err
	}
	info.Gold = 1000
	info.Sex = 1
	err = util.Bson2struct(result, info)
	return
}
