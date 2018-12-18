package Controller

import (
    "fmt"

    "github.com/globalsign/mgo/bson"
    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Client/Mongo"
)

type MongoController struct {
    pgo.Controller
}

// curl -v http://127.0.0.1:8000/mongo/insert
func (m *MongoController) ActionInsert() {
    // 获取mongo的上下文适配对象
    mongo := m.GetObject(Mongo.AdapterClass, "test", "test").(*Mongo.Adapter)

    // 通过map插入
    doc1 := pgo.Map{"f1": "val1", "f2": true, "f3": 99}
    err := mongo.InsertOne(doc1)
    fmt.Println("insert one doc1:", err)

    // 通过bson.M插入
    doc2 := bson.M{"f1": "val2", "f2": false, "f3": 10}
    err = mongo.InsertOne(doc2)
    fmt.Println("insert one doc2:", err)

    // 通过struct插入
    doc3 := struct {
        F1 string `bson:"f1"`
        F2 bool   `bson:"f2"`
        F3 int    `bson:"f3"`
    }{"val3", false, 6}
    err = mongo.InsertOne(doc3)
    fmt.Println("insert one  doc3:", err)

    // 批量插入
    docs := []interface{}{
        bson.M{"f1": "val4", "f2": true, "f3": 7},
        bson.M{"f1": "val5", "f2": false, "f3": 8},
        bson.M{"f1": "val6", "f2": true, "f3": 9},
    }
    err = mongo.InsertAll(docs)
    fmt.Println("insert all docs:", err)
}

// curl -v http://127.0.0.1:8000/mongo/update
func (m *MongoController) ActionUpdate() {
    // 获取mongo的上下文适配对象
    mongo := m.GetObject(Mongo.AdapterClass, "test", "test").(*Mongo.Adapter)

    // 更新单个文档
    query := bson.M{"f1": "val1"}
    update := bson.M{"$inc": bson.M{"f3": 2}}
    err := mongo.UpdateOne(query, update)
    fmt.Println("update one f1==val1:", err)

    // 更新多个文档
    query = bson.M{"f3": bson.M{"$gte": 7}}
    update = bson.M{"$set": bson.M{"f2": false}}
    err = mongo.UpdateAll(query, update)
    fmt.Println("update all f3>=7:", err)

    // 更新或插入
    query = bson.M{"f1": "val10"}
    update = bson.M{"$inc": bson.M{"f3": 2}}
    err = mongo.UpdateOrInsert(query, update)
    fmt.Println("update or insert f1==val10:", err)
}

// curl -v http://127.0.0.1:8000/mongo/query
func (m *MongoController) ActionQuery() {
    // 获取mongo的上下文适配对象
    mongo := m.GetObject(Mongo.AdapterClass, "test", "test").(*Mongo.Adapter)

    // 查询单个文档(未指定结果类型，结果为bson.M)
    var v1 interface{}
    err := mongo.FindOne(bson.M{"f1": "val1"}, &v1)
    fmt.Println("query one f1==val1:", v1, err)

    // 查询单个文档(结果类型为map)
    var v2 pgo.Map
    err = mongo.FindOne(bson.M{"f1": "val2"}, &v2)
    fmt.Println("query one f1==val2:", v2, err)

    // 查询单个文档(结果类型为struct)
    var v3 struct {
        Id bson.ObjectId `bson:"_id"`
        F1 string        `bson:"f1"`
        F2 bool          `bson:"f2"`
        F3 int           `bson:"f3"`
    }
    err = mongo.FindOne(bson.M{"f1": "val3"}, &v3)
    fmt.Println("query one f1==val3:", v3, err)

    // 查询多个文档(指定结果为map)
    var docs []pgo.Map
    err = mongo.FindAll(bson.M{"f3": bson.M{"$gte": 6}}, &docs)
    fmt.Println("query all f3>=6:", docs, err)
}
