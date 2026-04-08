[Gin框架笔记网址](https://gitee.com/moxi159753/LearningNotes/tree/master/Golang/Gin框架/1_Gin内容介绍)

### Gin返回JSON

方式一（临时）：

```go
r := gin.Default()

r.GET("/json", func(c *gin.Context){
	data := gin.H{"name":"小明", "message":"hello", "age":"18"}
	c.JSON(http.StatusOK, data)
})
```

方式二：

```go
type msg struct{
	Name string `json:"name"`
	Message string
	Age int
}

r.GET("/another_json", func(c *gin.Context){
	data := msg{
		"小明",
		"hello",
		18,
	}
	c.JSON(http.StatusOK, data) //json序列化
})
```



### Gin获取querystring参数

`querystring`指的是URL中`?`后面携带的参数，例如：`/user/search?username=小王子&address=沙河`。 获取请求的querystring参数的方法如下：

```go
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
        // 可以添加默认值
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username") 获取请求中携带的querystring参数
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
```

我们输入对应的URL，就能获取到对应的参数了

```
http://localhost:9090/web?username=小王子&address=沙河
```



### Gin获取form表单

什么是form表单：

```html
<!-- 这就是一个标准表单 -->
<form action="提交地址" method="提交方式">
  <!-- 输入框：用户名 -->
  <input type="text" placeholder="请输入用户名">
  
  <!-- 密码框 -->
  <input type="password" placeholder="请输入密码">
  
  <!-- 提交按钮 -->
  <button type="submit">登录</button>
</form>
```

关键属性解释：

1. **`<form>`**：表单的**容器**，所有要提交的内容都必须写在里面

2. **action**：数据要提交到的**后台地址**（后端接口）

3. method

   ：提交方式，最常用两种：

   - `get`：数据拼在网址里（搜索、查询用）
   - `post`：数据隐藏提交（登录、注册、上传文件用）

4. **`<input>`**：用户输入的控件（文本、密码、单选、多选、文件等）

5. **`type="submit"`**：提交按钮，点了就把数据发出去



请求的数据通过form表单来提交，例如向`/user/search`发送一个POST请求，获取请求数据的方式如下：

```go
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")获取form表单的值
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}
```



### Gin获取path参数

请求的参数通过URL路径传递，例如：`/user/search/小王子/沙河`。 获取请求URL路径中的参数的方式如下。

```go
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
    // 有多个获取路径的函数时注意url的匹配不要冲突，
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address") //Param返回的都是string类型
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8080")
}
```



### Gin定义中间件

Gin中的中间件必须是一个`gin.HandlerFunc`类型，即参数为c * gin.Context的函数。

**记录接口耗时的中间件**

例如我们像下面的代码一样定义一个统计请求耗时的中间件。

```go
// StatCost 是一个统计耗时请求耗时的中间件
// 这里采用一个闭包，返回一个gin.HandlerFunc类型，即func(c *gin.Context)
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}
```





### GORM教程

https://www.liwenzhou.com/posts/Go/gorm/





![image-20260328103704332](D:\Typora\网络图片\image-20260328103704332.png)





### GORM查询

- **`FirstOrCreate`**: 查不到就**创建并保存**到数据库。
- **`FirstOrInit`**: 查不到就**初始化**一个新对象，但**不会保存**，需要你手动调用 `Save()`。



所有查询指令：

```go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 定义模型
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 默认情况下，gorm创建的表将会是结构体名称的复数形式，如果不想让它自动复数，可以加一下禁用
	db.SingularTable(true)
	// 2, 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	// 3, 创建
	//u1 := User{Name: "eryajf", Age: 20}
	//db.Create(&u1)
	//u2 := User{Name: "jinzhu", Age: 22}
	//db.Create(&u2)

	// 4,查询
	var user []User
	db.Debug().First(&user) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键查询第一条记录：", user)

	db.Debug().Take(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL LIMIT 1
	fmt.Println("随机获取一条记录：", user)

	db.Debug().Last(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` DESC LIMIT 1
	fmt.Println("根据主键查询最后一条记录：", user)

	db.Debug().Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL
	fmt.Println("查询所有的记录：", user)

	db.Debug().First(&user, 2)      //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` = 2)) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询指定的某条记录：", user) //仅当主键为整型时可用

	Where
	db.Debug().Where("name = ?", "jinzhu").First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询第一条匹配条件记录：", user)

	db.Debug().Where("name = ?", "jinzhu").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu'))
	fmt.Println("查询所有匹配条件的记录：", user)

	db.Debug().Where("name <> ?", "jinzhu").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name <> 'jinzhu'))
	fmt.Println("查询name不等于jinzhu的所有记录：", user)

	db.Debug().Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name IN ('jinzhu','jinzhu 2')))
	fmt.Println("查询name在jinzhu和jinzhu 2的所有记录：", user)

	db.Debug().Where("name LIKE ?", "%jin%").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name LIKE '%jin%'))
	fmt.Println("查询name包含jin的所有记录：", user)

	db.Debug().Where("name = ? AND age >= ?", "jinzhu", "20").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu' AND age >= '20'))
	fmt.Println("查询两个条件都符合的所有记录：", user)

	oneDay, _ := time.ParseDuration("-24h")
	lastWeek := time.Now().Add(oneDay * 7)
	db.Debug().Where("updated_at > ?", lastWeek).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((updated_at > '2020-03-01 19:45:11'))
	fmt.Println("查询一周内更新的用户记录：", user)

	today := time.Now()
	db.Debug().Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((created_at BETWEEN '2020-03-01 19:52:51' AND '2020-03-08 19:52:51'))
	fmt.Println("查询一周内创建的记录：", user)

	db.Debug().Where(&User{Name: "jinzhu", Age: 22}).First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu') AND (`user`.`age` = 22)) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过结构体查询：", user)

	db.Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 22}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu') AND (`user`.`age` = 22))
	fmt.Println("通过map查询：", user)

	db.Debug().Where([]int64{1, 2}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` IN (1,2)))
	fmt.Println("通过主键的切片查询：", user)

	db.Debug().Not("name", "jinzhu").First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` NOT IN ('jinzhu'))) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询name不是jinzhu的第一条记录：", user)

	db.Debug().Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` NOT IN ('jinzhu','jinzhu 2')))
	fmt.Println("查询name不在jinzhu或jinzhu2的所有记录：", user)

	db.Debug().Not([]int64{1, 2, 3}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` NOT IN (1,2,3)))
	fmt.Println("查询主键不是1，2，3的所有记录：", user)

	db.Debug().Not([]int64{}).First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询所有用户中的第一个：", user)

	db.Debug().Not("name = ?", "jinzhu").First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND (NOT (name = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询name不是jinzhu的第一个用户：", user)

	db.Debug().Not(User{Name: "jinzhu"}).First(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` <> 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过结构体查询name不是jinzhu的第一个用户：", user)

	db.Debug().Where("age > ?", 25).Or("age < ?", 23).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((age > 25) OR (age < 23))
	fmt.Println("查询年龄小于23的或者大于25的所有记录：", user)

	// struct
	db.Debug().Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu') OR (`user`.`name` = 'jinzhu 2'))
	fmt.Println("结构体：查询名字是jinzhu的或者是jinzhu 2的所有记录：", user)

	// map
	db.Debug().Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu') OR (`user`.`name` = 'jinzhu 2'))
	fmt.Println("map：查询名字是jinzhu的或者是jinzhu 2的所有记录：", user)

	db.Debug().First(&user, 3)          //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`id` = 3)) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键查询指定的某条记录：", user) //仅当主键为整型时可用

	db.Debug().First(&user, "id = ?", "string_primary_key") //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((id = 'string_primary_key')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键是非整形主键获取记录：", user)

	db.Debug().Find(&user, "name = ?", "jinzhu") //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name = 'jinzhu'))
	fmt.Println("查询name为jinzhu的记录：", user)

	db.Debug().Find(&user, "name <> ? AND age > ? ", "jinzhu", "20") //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((name <> 'jinzhu' AND age > '20' ))
	fmt.Println("查询name不是jinzhu且年龄大于20的记录：", user)

	db.Debug().Find(&user, User{Age: 20}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`age` = 20))
	fmt.Println("通过结构体查询年龄是20的所有记录：", user)

	db.Debug().Find(&user, map[string]interface{}{"age": 20}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`age` = 20))
	fmt.Println("通过map查询年龄是20的所有记录：", user)

	db.Debug().FirstOrInit(&user, User{Name: "non_existing"}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("查询name为non_existing的记录：", user)

	db.Debug().Where(User{Name: "jinzhu"}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过结构体查询name为jinzhu的记录：", user)

	db.Debug().FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("通过map查询name为jinzhu的记录：", user)

	// 未找到
	db.Debug().Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	// 找到
	db.Debug().Where(User{Name: "jinzhu"}).Attrs(User{Age: 50}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	// 找到
	db.Debug().Where(User{Name: "jinzhu"}).Assign(User{Age: 50}).FirstOrInit(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().FirstOrCreate(&user, User{Name: "non_existing"}) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "jinzhu"}).FirstOrCreate(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'non_existing')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Where(User{Name: "jinzhu"}).Attrs(User{Age: 30}).FirstOrCreate(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((`user`.`name` = 'jinzhu')) ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)

	db.Debug().Select("name", "age").Find(&user) //SELECT name FROM `user`  WHERE `user`.`deleted_at` IS NULL'age'
	fmt.Println("查询表中name字段参数为age的记录：", user)

	db.Debug().Select([]string{"name", "age"}).Find(&user) //SELECT name, age FROM `user`  WHERE `user`.`deleted_at` IS NULL
	fmt.Println("列出表中name与age字段：", user)

	db.Debug().Order("age desc,name").Find(&user) //SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY age desc,name
	fmt.Println("根据年龄排序来查询：", user)

	db.Debug().Order("age desc").Order("name").Find(&user)
	fmt.Println("根据多个条件排序查询：", user)

}
```



### 项目结构：

```
your-project/
├── cmd/                  # 项目入口（main 函数）
│   └── server/
│       └── main.go       # 启动服务
├── router/
│   └── router.go         # 路由注册
├── api/                  # HTTP/gRPC 接口定义、handler
├── internal/             # 内部业务逻辑（核心）
│   ├── service/          # 业务逻辑层
│   ├── dao/              # 数据访问层（操作 DB/Redis）
│   ├── model/            # 数据结构、DO/POJO
│   └── middleware/        # 中间件
├── pkg/                  # 公共工具包（可被外部引用）
├── config/               # 配置文件、配置加载
├── scripts/              # 脚本、构建、部署
├── go.mod
└── go.sum
```



```
请求 → api(handler) → service(业务) → dao(数据) → DB/Redis
```

严格规则：

- handler 不直接调用 dao
- service 不写 HTTP 逻辑
- dao 不处理业务
- model 只存结构
- pkg 工具谁都能用
