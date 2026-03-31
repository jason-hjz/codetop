# MySQL笔记

[toc]

## 什么是SQL和MySQL

> **资料库管理系统**（Data base management system）
>
> ![image-20230311100807057](C:\Users\Hjz\AppData\Roaming\Typora\typora-user-images\image-20230311100807057.png)
>
> 资料库分为	SQL（关联式资料库）和  noSQL（非关联式资料库）
>
> 相应地，DBMS也分为	RDBMS（关联式资料库管理系统） 和 NRDBMS（非关联式...）
>
> 😍要学习的MySQL就属于RDBMS的一种



> 什么是SQL？🤔
>
> **SQL**	Structured Query Language (一种用于与RDBMS软件做沟通的语言)
>
> ![image-20230311101932302](C:\Users\Hjz\AppData\Roaming\Typora\typora-user-images\image-20230311101932302.png)
>
> PS：NRDBMS软件没有统一的语言



## 资料库（增删改查）



> primary key （主键） ：唯一区分每一个元组（每一笔资料）的属性或属性集
>
> foreign key（外键）：表R1的非主键属性是表R2的主键，用以实现表格之间关联



<font color="black" size = 5>**基础命令**</font>：

------

### 库

- 自定义的变量：	   \` 变量名 \`

- 创建资料库：            CREATE DATABASE \` 数据库名称 `；

- 展示已有资料库：    SHOW DATABASES；

- 删除资料库：            DROP DATABASE \` 数据库名称 `；

- 使用资料库：            USE \` 数据库名称 `；

- 资料类型：

  | INT              | 整数                                           |
  | ---------------- | ---------------------------------------------- |
  | **DECIMAL(m,n)** | **浮点数（有小数点的数）；m位数，小数占了n位** |
  | **VARCHAR(n)**   | **字串；最多存放n个字原**                      |
  | **BLOB**         | **（Binary Large Object）图片 影片 档案**      |
  | **DATE**         | **‘ YYYY-MM-DD ’ 日期**                        |
  | **TIMESTAMP**    | **‘ YYYY-MM-DD  HH:MM:SS ’ 时间**              |



### 表格

- 创建表格：

```
CREATE TABLE `表格名称`(
	`主键名称` 类型 PRIMARY KEY 限制（可有可无）,
	`属性1名称` 类型 限制,
	`属性2名称` 类型 限制
  //PRIMARY KEY(`主键名称`)
  //FOREIGN KEY(`外键名称`) REFERENCES `外键链接的表格`(相应的主键) ON DELETE CASCADE
);
```

- 展示表格：				DESCRIBE \`表格名称`;
- 删除表格：                DROP TABLE \`表格名称`;
- 添加属性：                ALTER TABLE \`表格名称` ADD 属性名 类型；
- 删除属性：                ALTER TABLE \`表格名称` DROP COLUMN 属性名；
- 添加外键：                ALTER TABLE \`表格名称\` ADD FOREIGN KEY （\`外键名称\`)

​										  REFERENCES \`外键链接的表格\`（\`相应的主键\`）ON DELETE SET NULL

> ON DELETE SET NULL 表示当删除该主键，相应的外键置为NULL
>
> ON DELETE CASCADE 表示当删除该主键，删除相应外键对应的这笔资料



- 填入资料(元组)：  　INSERT INTO \`表格名称`  VALUES (分量1，分量2，‘字串分量3’);

  ​								   INSERT INTO \`表格名称`（指定需要填写的分量)  VALUES (分量值)；

  [^填入资料指令]: 默认按顺序依次填入表格，分量是字串时要加 ‘ ’ 或 “ ”，空分量用NULL；

  > <font color="blue" size = 3>**特别得，当表格1中有外键且其链接的表格2为空(未填写)时，无法填入资料，可以先将外键属性值全置NULL，将外键链接的表格2资料填入后，再表格1中的外键属性值依次修改。**</font>

- 查看表格资料：   　 SELECT * FROM \`表格名称`;



- constraints 限制：

  | NOT NULL                      | 不为空                                           |
  | ----------------------------- | ------------------------------------------------ |
  | **UNIQUE**                    | **值唯一不可重复**                               |
  | **DEFAULT  ‘ 值 ’**（不是``） | **预设值，当某元组该属性未填入时自动填入预设值** |
  | **AUTO_INCREMENT**            | **该属性每次自动加一**                           |
  | **check（属性in（a，b））**   | **限制该属性一定在范围a，b中**                   |



### 修改删除

- <font color="green" size = 4>**以下修改和删除操作需要先关闭预设的更新模式**</font>：

```
用　SET SQL_SAFE_UPDATES = 0；关闭预设的更新模式
```

- 修改资料(分量或元组)：      

```
UPDATE `表格名称`
SET `要修改的分量的属性1` = `修改后的新值`，`要修改的分量的属性2` = `修改后的新值`，...，修改n
WHERE 要修改的属性满足的条件；
```

- 删除资料(元组或表格)：

```
DELETE FROM `表格名称`
WHERE 要删除的元组满足的条件；
```



- 在父表中用DELETE或UPDATE操作删除父表与子表有匹配行的键值时

| CASCADE         | 删除父表中的行，且自动删除子表中匹配的行                     |
| --------------- | ------------------------------------------------------------ |
| **SET DEFAULT** | **删除父表中的行，且自动设置子表中的外键为缺省值（创建时设定的DEFAULT值）** |
| **SET NULL**    | **删除父表中的行，且自动设置子表中的外键值为NULL**           |
| **RESTRICT**    | **默认操作，拒绝对父表进行操作**                             |

- 在子表中用INSERT或UPDATE操作插入与父表中键值不匹配的外键值时

| RESTRICT | 拒绝对子表进行操作 |
| -------- | ------------------ |



### 查询

- SELECT DISTINCT（不重复）\`属性名1\`， \`属性名2\`，...  ，\`属性名n\` FROM \`表格名称\`；

  [^select * from `表格名称`]: 这里的 * 指的是所有属性

- 排序：

```
SELECT * FROM `表格名称`
ORDER BY `属性名1`，`属性名2` DESC   /*将表格根据属性名1的值排序，默认由低到高↓，加 DESC限制后由高到低↓；如果两值相等，再比较属性名2的值进行排序*/
LIMIT n  	//加入LIMIT n后可以只传回资料（元组）的前n行
WHERE `属性名` IN（‘属性1’，‘属性2’，‘属性3’）;	//当该属性为属性123中的任意一个时
GROUP BY `属性名` //按该属性的值分组，值相同的为一组
HAVING 组应该满足的条件
```

> HAVING 必须在分好组的前提下才能用	group by、having、order by的使用顺序：group by 、having、order by
>
> having后可以用聚合函数而where不行



### 聚合函数 

**aggregate functions**：

- SELECT COUNT（*）FROM \`表格名称\`  					     计算资料(元组)数量
- SELECT AVG（\`属性名\`）FROM \`表格名称\`                   计算属性平均值 
- SELECT SUM（\`属性名\`）FROM \`表格名称\`                  计算属性值的和
- SELECT MIN / MAX（\`属性名\`）FROM \`表格名称\`       计算属性最小值/最大值



### 万用字元

**wildcards**：

> % 代表多个字元        _ 代表一个字元

- eg：取得电话尾数为138的客户

​               SELECT * FROM \`client\` WHERE \`number\` LIKE '%138';

- eg：取得生日在12月的员工

​			   SELECT * FROM \`employee\` WHERE \`birth_date\` LIKE '\_\____12%' ;（ps：5个\_）



### 联合

**union**（把两个查询结果合并 两列合成一列）

- SELECT \`属性名1\`AS \`新名字\`，...，\`属性名n\` AS \`新名字\`  FROM  \`表格名称`

​		UNION

​		SELECT \`属性名1\`，...，\`属性名m\`FROM \`表格名称`

​		......

> 联合的各个查询中 属性数必须相等即n=m，且合并的属性类型必须相同
>
> 联合后的属性名为新名字 如果没有定义新名字则默认为第一个查询中的属性名



### 连接

**join**

- SELECT * FROM \`表格1\` JOIN \`表格2\`

​	   ON \`表格1\`.\`主键\` = \`表格2\`.\`对应外键\`

​	   WHERE    条件  ;

> **LEFT / RIGHT  JOIN 表示左/右边无论条件是否成立都会回传所有资料，右/左边只有条件（ON）成立才会回传资料，否则回传NULL** 



### 子查询

- SELECT \`属性名\` FROM \`表格\`

​		WHERE \`属性名\` [= < > in]\(

​		子查询语句

​		)；



### 视图

- **创建视图**

  视图可以从一张表、几张表或其他视图中创建

  ```
  CREATE VIEW `视图名`(属性名1，属性名2，...)     //如果没有列出属性名则视图展示下面select子查询所有属性
  AS
  SELECT 语句
  WITH CHECK OPTION；		//可有可无
  ```

- **删除视图**

  ```
  DROP VIEW `视图名`
  RESTRICT	//默认，存在依赖被删除视图的其他对象，不允许删除
  CASCADE；   //删除视图可能产生级联删除
  ```

- 查询更新语句与表格类似



### SQL的完整性约束

- 非过程性完整性约束

  1. 数据取值要求CONSTRAINTS
  2. 域约束
  3. 实体完整性（主键值唯一且不能为空）
  4. 参照完整性（外键值为对应主键的值或者为空）
  5. 一般约束（企业约束）

- 过程性完整性约束

  6. 触发器

     

### 域

> **域本质上是一种可带约束CONSTRAINTS的数据类型，域是一组具有相同数据类型的值的集合**

- 创建

  ```
  CREATE DOMIN `域名` AS `数据类型`
  CHECK()；
  ```
  
- 删除

  ```
  DROP DOMIN `域名`；
  ```



### 授权与回收

- 授权

  ```
  GRANK ALL PRIVILEGES | 特权1，特权2，...
  ON `表名`
  TO `授权对象的名字`
  WITH GRANK OPTION；	//授予授权的权限
  ```

- 回收

  ```
  REVOKE GRANK OPTION FOR		//收回授权的权限
  ALL PRIVILEGES | 特权1，特权2，...
  ON `表名`
  FROM `授权对象的名字`
  ```



### 事务

```
BEGIN TRANSACTION	//事务开始
COMMIT				//事务正常结束，保存事务操作
ROLLBACK			//事务异常结束，回滚到事务开始时的状态
```

