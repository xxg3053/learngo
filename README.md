# study go

#### 特别
- 没有对象，没有继承多态，没有泛型，没有try/catch
- 有接口，函数式编程，CSP并发模型（goroutine + channel）
- 学习go很简单，因为语法简单
- 学好go不简单

#### 涵盖内容
基本语法：变量、选择、循环、指针、数组、容器   
面向接口：结构体、dock typing概念、组合思想   
函数式编程：闭包的概念、多样立例题   
工程化：资源管理、错误处理、测试和文档、性能调优   
并发编程：goroutine和channel、理解调度器、多样示例   

##### 变量
[代码](https://github.com/xxg3053/learngo/blob/master/lang/base.go)
- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有char，只有rune
- 原生支持复数
- 枚举、iota关键字

##### 条件
[代码](https://github.com/xxg3053/learngo/blob/master/lang/branch.go)
- for，if后面的条件没有括号
- if条件里可以定义变量
- 没有while
- switch不需要break， 也可以直接switch多个条件

##### 数组、切片、map
推荐用range遍历   
###### 数组
[代码](https://github.com/xxg3053/learngo/blob/master/lang/arrays.go)
- 数组是值类型
- ```[10]int```和```[20]int```是不同类型
- 调用```func f(arr [10]int)```会拷贝数组 
- 在go语言中一般不直接使用数组

###### 切片(Slice)
[代码](https://github.com/xxg3053/learngo/blob/master/lang/slice.go)
- Slice本身没有数据，是对底层array的一个view
- reslice
- Slice的扩展

###### map

###### 字符

##### 函数
[代码](https://github.com/xxg3053/learngo/blob/master/lang/func.go)
- 返回多个值，可以起名字
- 函数作为参数
- 没有默认参数，可选参数
- 可变参数列表

##### 指针
[代码](https://github.com/xxg3053/learngo/blob/master/lang/pointer.go)
- 指针不能运算
- 值传递和引用传递？go语言只有值传递一种方式


##### goroutine
[代码](https://github.com/xxg3053/learngo/blob/master/lang/goroutine/goroutine.go)
协程Coroutine
- 轻量级"线程"
- 非抢占式多任务处理，由协程主动交出控制权（runtime.Gosched()//交出协程控制权）
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或多个线程上运行

##### channel
[代码](https://github.com/xxg3053/learngo/blob/master/lang/channel/done.go)
- channel
- buffered channel
- range
- 理论基础：Communication Sequential Process(CSP)
- Don`t communicate by sharing memory; share memory by communicating
- 不要通过共享内存来通信；通过通信来共享内存

#### 实战项目
分布式爬虫   

#### 第三方库
转码包 golang.org/x/text    
检查html编码  golang.org/x/net/html   

#### 网站资源选择
css选择器   
xpath  
正则   


#### 基本流程
Seed --request--Engine--任务队列  
Engine--URL---Fetcher---text   
Engine--text---Parser---items   

