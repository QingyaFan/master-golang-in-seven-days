# 语法

特点：有指针

1、类型与变量

array、slice、map、struct、interface

2、程序控制

- 条件判断

```go
if condition {
    // codes
}
```

```go
switch expression {
case condition:
    // codes
}
```

- 循环

```go
for index := 0; index < count; index++ {
    // codes
}

for _, var := range var {
    // codes
}
```

3、函数

```go
func func_name(params) return_val {
    // codes
}
```

4、类与对象

golang中没有类的概念，但像C语言一样，提供了struct，C语言中的struct的成员变量都是公有的，而golang中变量或方法首字母大写，才是公有的。下面上一个例子：

```go
package main

import "fmt"

type Teacher struct {
    Name  string
    age   int
    title string
}

func (teacher *Teacher) getName() string {
    return teacher.Name
}

func main() {
    tea1 := new(Teacher)
    tea1.Name = "cheerfun"

    fmt.Println(tea1.getName())
}
```

5、模块与包

同一个包（package）中的不同文件定义的变量和方法如果首字母大写，那么可以在其他文件中使用。
