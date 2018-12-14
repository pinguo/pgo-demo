package Controller

import (
    "fmt"
    "time"

    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Client/Db"
)

type MysqlController struct {
    pgo.Controller
}

// 使用db.Exec/db.ExecContext在主库上执行非查询操作，
// pgo默认使用10s超时，可通过context来自定义超时。
func (m *MysqlController) ActionExec() {
    // 获取db的上下文适配对象
    db := m.GetObject(Db.AdapterClass).(*Db.Adapter)

    // 执行插入操作
    age := (time.Now().Nanosecond() / 1000) % 100
    name := fmt.Sprintf("name%d", age)
    query := "INSERT INTO `test` (`name`, `age`) VALUES (?, ?)"

    res := db.Exec(query, name, age)
    if res == nil {
        panic("exec insert failed, " + query)
    }

    lastId, _ := res.LastInsertId()
    numRow, _ := res.RowsAffected()
    fmt.Println("insert: ", lastId, numRow)

    // 执行修改操作
    query = "UPDATE test SET age=age+1 WHERE id=?"
    res = db.Exec(query, lastId)
    if res == nil {
        panic("exec update failed, " + query)
    }

    lastId1, _ := res.LastInsertId()
    numRow1, _ := res.RowsAffected()
    fmt.Println("update: ", lastId1, numRow1)

    // 执行删除操作
    query = "DELETE FROM test WHERE id=?"
    res = db.Exec(query, lastId)
    if res == nil {
        panic("exec delete failed, " + query)
    }

    lastId2, _ := res.LastInsertId()
    numRow2, _ := res.RowsAffected()
    fmt.Println("delete: ", lastId2, numRow2)
}

// 使用db.Query/QueryContext来查询数据，若当前db对象未执行过任何操作，
// 则使用从库进行查询，否则复用上一次获取的db连接。
func (m *MysqlController) ActionQuery() {
    // 获取db的上下文适配对象
    db := m.GetObject(Db.AdapterClass).(*Db.Adapter)

    query := "SELECT id, name, age FROM `test` WHERE age>?"
    rows := db.Query(query, 10)
    if rows == nil {
        panic("query failed, query: " + query)
    }

    defer rows.Close()

    for rows.Next() {
        id, name, age := 0, "", 0
        if err := rows.Scan(&id, &name, &age); err != nil {
            panic("query scan faild, err: " + err.Error())
        }
        fmt.Println("user: ", id, name, age)
    }
}

// 使用db.Prepare/db.PrepareContext来执行批量操作，默认查询操作在
// 从库上进行，其余操作在主库上进行，若当前db对象有过其它操作，则查询
// 操作会复用之前的连接。
func (m *MysqlController) ActionPrepare() {
    // 获取db的上下文适配对象
    db := m.GetObject(Db.AdapterClass).(*Db.Adapter)

    query := "INSERT INTO test (name, age) VALUES (?, ?)"
    stmt := db.Prepare(query)

    for i := 0; i < 10; i++ {
        name := fmt.Sprintf("name%d", i)
        res := stmt.Exec(name, i)
        if res == nil {
            panic("stmt exec failed, " + query)
        }
    }
}

// 使用db.Begin/BeginContext/Commit/Rollback来进行事务操作
func (m *MysqlController) ActionTransaction() {
    // 获取db的上下文适配对象
    db := m.GetObject(Db.AdapterClass).(*Db.Adapter)

    db.Begin()
    db.Exec("INSERT INTO test (name, age) VALUES (?, ?)", "name1", 1)
    db.Exec("UPDATE test SET age=age+1 WHERE id=?", 1)
    db.Commit()
}
