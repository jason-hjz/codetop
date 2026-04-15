[RabbitMQ笔记文档](https://www.fengfengzhidao.com/article/B-EpYZgB8lppN5cbEgGQ)



### 简单流程

生产者将消息发给交换器的时候，一般会指定一个 **RoutingKey(路由键)**，交换机根据这个key决定把消息发给哪些队列

RabbitMQ 中通过 **Binding(绑定)** 将 **Exchange(交换器)** 与 **Queue(消息队列)** 关联起来，在绑定的时候一般会指定一个 **BindingKey(绑定键)** ,这样 RabbitMQ 就知道如何正确将消息路由到队列了



```
生产者
   ↓
发消息时指定：【Routing Key】
   ↓
Exchange（交换机）
   ↓
根据 Binding Key 做匹配
   ↓
Queue（队列） ← 消费者监听
```

------

完整工作流程

1. 先做绑定（提前建好）

- 交换机：`ex_demo`

- 队列：`queue_order`

- 绑定关系

  ：把 

  ```
  queue_order
  ```

   绑到 

  ```
  ex_demo
  ```

  ，并指定 Binding Key = order.create

2. 生产者发消息

- 发到交换机：`ex_demo`
- 带上 **Routing Key = order.create**

3. 交换机做匹配

- 交换机拿着消息的 **Routing Key**去对比所有绑定的 **Binding Key**

- 匹配上 → 消息扔进对应队列

4. 消费者消费

消费者监听 **队列名 queue_order** 拿到消息



Routing Key

- **生产者指定**
- 是消息身上的 **“地址标签”**
- 作用：告诉交换机 “我要往哪条路发”

Binding Key

- **绑定时指定**
- 是队列在交换机那登记的 **“收件规则”**
- 作用：告诉交换机 “什么样的消息我要收”



### SSL安全配置

注意：mq的协议是amqps，端口是5671

```
// 1. 加载客户端证书和密钥（双向认证时需要）
  cert, err := tls.LoadX509KeyPair("D:\\client_certificate.pem", "D:\\client_key.pem")
  if err != nil {
    log.Fatalf("加载客户端证书失败: %v", err)
  }

  // 2. 加载CA证书（验证服务器证书）
  caCert, err := os.ReadFile("D:\\ca_certificate.pem")
  if err != nil {
    log.Fatalf("读取CA证书失败: %v", err)
  }
  caCertPool := x509.NewCertPool()
  caCertPool.AppendCertsFromPEM(caCert)

  // 3. 配置TLS
  tlsConfig := &tls.Config{
    Certificates:       []tls.Certificate{cert}, // 客户端证书（双向认证时需要）
    RootCAs:            caCertPool,              // 信任的CA
    InsecureSkipVerify: false,                   // 必须验证服务器证书
  }

  // 连接 RabbitMQ
  conn, err := amqp.DialTLS("amqps://admin:password@192.168.80.165:5671/", tlsConfig)
  if err != nil {
    log.Fatalf("无法连接到 RabbitMQ: %v", err)
  }
```



### Fanout模式

广播：把所有发送到该 Exchange 的消息路由到所有与它绑定的 Queue 中，**忽略 BindingKey**

下列这种方式可以缓存，消费者可以消费到它运行之前生产者生产的消息

| 生产者                | 消费者              |
| --------------------- | ------------------- |
| 1、连接MQ 创建通道    | 1、连接MQ、创建通道 |
| 2、声明交换器         | 2、消费消息         |
| 3、创建队列           |                     |
| 4、将队列绑定到交换器 |                     |
| 5、发送消息           |                     |

**交换机、队列、绑定关系，是存在于 RabbitMQ 服务器（Broker）上的，不是存在于代码里的！**

只要**生产者创建过一次**，服务器上就永久保存了，**消费者直接拿来用就行**。



### Direct模式

把消息路由到那些 BindingKey 与 RoutingKey **完全匹配**的 Queue 中



### Topic模式

基于路由键的 “模糊匹配” 分发，支持通配符，比路由模式更灵活

绑定队列时可使用通配符：

- `*`：匹配一个单词（如 `order.*` 匹配 `order.create`，但不匹配 `order.create.success`）。
- `#`：匹配零个或多个单词（如 `order.#` 匹配 `order.create`、`order.create.success`）。

```
//topic 模式用来获取命令行参数的代码

// 从命令行参数中获取消息体内容
func bodyFrom(args []string) string {
	// 声明一个字符串变量 s，用来存储最终的消息内容
	var s string

	// 判断：如果命令行参数个数小于3  或者  第2个参数为空
	// args[0] 是程序名，args[1] 是路由键，args[2]开始才是消息内容
	if (len(args) < 3) || os.Args[2] == "" {
		// 不满足条件时，默认消息内容为 "hello"
		s = "hello"
	} else {
		// 满足条件：从 args[2] 开始到最后所有参数，用空格拼接成字符串
		// 比如命令行输：go run main.go info hello world  → 得到 "hello world"
		s = strings.Join(args[2:], " ")
	}

	// 返回最终的消息内容
	return s
}

// 从命令行参数中获取 路由键/级别（routing key/severity）
func severityFrom(args []string) string {
	// 声明字符串变量 s，存储路由键
	var s string

	// 判断：如果命令行参数个数小于2  或者  第1个参数为空
	if (len(args) < 2) || os.Args[1] == "" {
		// 默认路由键：anonymous.info
		s = "anonymous.info"
	} else {
		// 使用命令行传入的第1个参数作为路由键
		// 例如：go run main.go error  → s = "error"
		s = os.Args[1]
	}

	// 返回路由键（severity/routing key）
	return s
}
```



### 死信队列

处理 “无法被正常消费” 的消息（如消费超时、消费失败且重试次数耗尽），将其转移到专门的 “死信队列”，便于后续排查问题。

- 为普通队列设置死信交换器（`x-dead-letter-exchange`）和死信路由键（`x-dead-letter-routing-key`）。
- 当消息满足死信条件（如被拒绝 `basic.reject` 且 `requeue=false`、TTL 过期、队列满）时，会被路由到死信队列。

设置：生产者在声明队列时设置死信队列参数

```
 // 声明普通队列并设置死信交换器
  args := amqp.Table{
    "x-dead-letter-exchange":    "dlx_exchange",
    "x-dead-letter-routing-key": "dlx_key",
  }
  _, err = ch.QueueDeclare(
    "normal_queue", // name
    true,           // durable
    false,          // delete when unused
    false,          // exclusive
    false,          // no-wait
    args,           // arguments
  )
```

死信消费者：

```
//	声明死信交换器
//	声明死信队列
//	绑定死信队列到死信交换器
//	消费死信队列中的消息
```



### 消息优先级

声明队列时 用 amqp.Table{"x-max-priority": 10,} 设置最大优先级

在publish的时候设置Priority字段来设置消息优先级

```
err = ch.Publish(
      "",     // 交换机
      q.Name, // 路由键（队列名称）
      false,  // 非强制
      false,  // 非立即
      amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(msg.body),
        Priority:    msg.priority, // 设置消息优先级
      })
```



### 回调队列

生产者声明一个回调队列 并用Consume监听回调队列消费消息

消费者消费正常队列消息后 Publish发送回调消息到回调队列



### 流量控制

生产者 通过 ch.QueueInspect(q.Name) 获取队列状态

消费者通过 ch.Qos()  设置消费的能力