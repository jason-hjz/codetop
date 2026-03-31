网站：https://golangstar.cn/



代理：GOPROXY=https://goproxy.cn,direct



### var 和 type

**用 `var` 的核心场景**：需要创建 “存储数据的变量” 时（函数内 / 外声明变量、初始化 / 暂不初始化变量），`:=` 是 `var` 的简化写法（仅函数内可用）。

**用 `type` 的核心场景**：需要 “定义新的类型规则” 时（定义结构体、接口、自定义类型、类型别名），是创建自定义类型的唯一方式。



### 匿名函数

`func(i int) { ... }(i)` 中，`{}` 是匿名函数的**函数体**，后面的 `(i)` 是**立即调用这个匿名函数，并把当前循环的变量 `i` 作为参数传递进去**。



### sync锁

sync.RWMutex

读写锁互斥、写写锁互斥、读读可以同时加锁



### goroutine

每个goroutine结束前要做的两件事：

```
defer func(){
	wg.Done
	mu.Unlock
}()
```



### 输入输出

```go
//输出
// 1. 打印内容，不换行
	fmt.Print("Hello ")
	fmt.Print("Go\n") // 手动换行

// 2. 打印内容，自动换行（最常用！）
	fmt.Println("Hello Go")
	fmt.Println(123, 456, "abc") // 多个值用逗号分隔

//输入
// 1. fmt.Scan 遇到空格就会停止
	fmt.Print("请输入姓名和年龄：")
	fmt.Scan(&name, &age) // 空格/回车分隔输入
	fmt.Printf("姓名：%s 年龄：%d\n", name, age)
	
// 2. 读一整行
    import (
        "bufio"
        "fmt"
        "os"
    )

    func main() {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("请输入一句话：")
        scanner.Scan() // 读取一行
        text := scanner.Text()
        fmt.Println("你输入的是：", text)
    }
```

%T 打印类型

%v 打印任意类型

%#v 打印任何类型，常用于打印空字符串“ ”



### 接口

```
    var phone Phone      // 声明一个接口类型phone
    phone = new(Apple)   // 注意这种创建方式，new函数参数是接口的实现
```

首先声明phone为一个接口，然后用new方法定义phone，new的参数必须是phone的一种实现，即这里的Apple结构体必须实现接口Phone中的方法。



###  `range` 遍历规则

|    遍历类型     | 第一个返回值 |  第二个返回值  |                       示例（核心代码）                       |
| :-------------: | :----------: | :------------: | :----------------------------------------------------------: |
|   数组 / 切片   | 索引（int）  |     元素值     | `nums := []int{1,2}; for i, v := range nums { // i=0,v=1; i=1,v=2 }` |
|     字符串      |   字节索引   | 字符值（rune） | `s := "你好"; for i, c := range s { // i=0,c='你'; i=3,c='好' }` |
|       map       |  key（键）   |  value（值）   | `m := map[string]int{"a":1}; for k, v := range m { // k="a",v=1 }` |
| 通道（channel） |    元素值    |       无       | `ch := make(chan int); go func(){ch<-10; close(ch)}(); for v := range ch { // v=10 }` |

注意：

1、遍历数组 / 切片时，value 是副本：修改 `v` 不会改变原数组 / 切片的值，如需修改，要通过索引操作；

2、遍历map时是无序的



### return

`return`不是原子级操作，执行过程是: 设置返回值—>执行`defer`语句—>将结果返回



### new和make的区别

new：一个通用的内存分配器，为任何类型（除了下面三种）分配内存并返回**指针**。
make：一个专用的初始化工具，只为 slice、map、channel 这三种引用类型服务，并返回一个可直接使用的**值**。
