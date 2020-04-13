# ormx
Object Relational Mapping + Log
==

<pre>
CREATE TABLE user  (
  id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  phone varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
  PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET=utf8mb4;

type User struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
}
</pre>

<pre>
profile, err := log.NewProfile(map[string]string{
    log.ProfileDirectory: "./log",
    log.ProfileChannel:   "orm",
})

if err != nil {
    fmt.Println(err)
    return
}

stream, err := log.NewStream(profile)
if err != nil {
    fmt.Println(err)
    return
}

logger, err := log.NewLogger(stream, "test", 1)
if err != nil {
    fmt.Println(err)
    return
}

idRegister := &log.IdRegister{}
idRegister.SetTraceId("trace-id-10001")
idRegister.SetBizId("biz-id-20002")
</pre>

<pre>
write, err := database.NewProfile(map[string]string{
    database.ProfileId:       "test",
    database.ProfileDriver:   "mysql",
    database.ProfileHost:     "127.0.0.1",
    database.ProfileDatabase: "test",
    database.ProfileUsername: "root",
    database.ProfilePassword: "123456",
    database.ProfileWrite:    "true",
})

if err != nil {
    fmt.Println(err)
    return
}

builder := database.DriversBuilder{}

err = builder.AddProfile(write)
if err != nil {
    fmt.Println(err)
    return
}

drivers, err := builder.Build()
if err != nil {
    fmt.Println(err)
    return
}

driver, err := drivers.GetWriter()
if err != nil {
    fmt.Println(err)
    return
}

orm, err := ormx.New(logger)
if err != nil {
    fmt.Println(err)
    return
}
</pre>

<pre>
var result []*User
affectedRows, err11 := orm.Find(idRegister, driver, &result, "SELECT id, name, phone FROM user")
if err11 != nil {
    fmt.Println(err11)
    return
}

fmt.Println("affectedRows: ", affectedRows)

for _, user := range result {
    fmt.Println(user)
}
</pre>

<pre>
var result User
affectedRows, err21 := orm.First(idRegister, driver, &result, "SELECT id, name, phone FROM user")
if err21 != nil {
    fmt.Println(err21)
    return
}

fmt.Println("affectedRows: ", affectedRows)
fmt.Println(result)
</pre>

<pre>
user := &User{
    Name:  "张三",
    Phone: "13000000001",
}

result, err31 := orm.Insert(idRegister, driver, "user", user)
if err31 != nil {
    fmt.Println(err31)
    return
}

affectedRows, _ := result.RowsAffected()
lastInsertId, _ := result.LastInsertId()
fmt.Println("affectedRows: ", affectedRows)
fmt.Println("lastInsertId: ", lastInsertId)
</pre>

<pre>
affectedRows, err41 := orm.Update(idRegister, driver, "user", map[string]interface{}{"phone": "18000000002"}, 2)
if err41 != nil {
    fmt.Println(err41)
    return
}

fmt.Println("affectedRows: ", affectedRows)
</pre>

<pre>
affectedRows, err51 := orm.Delete(idRegister, driver, "user", 1)
if err51 != nil {
    fmt.Println(err51)
    return
}

fmt.Println("affectedRows: ", affectedRows)
</pre>

<pre>
// 进程正常关闭前
err61 := drivers.Close()
fmt.Println(err61)
</pre>
