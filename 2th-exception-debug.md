# Exception & Debug

错误和异常处理是每个程序语言中不可或缺的一块，golang和其他语言的错误处理理念不太一样，golang推崇在错误发生的地方检查并处理错误，而不是throw出错误，再在其他地方catch错误并处理。因此，golang中并不直接支持 try/catch/finally。

## 错误处理

在golang中使用最多的错误处理是error对象和nil的比较，这种模式要求被调用函数若执行出错，应返回error对象。如下面的用法：

```go
f, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
```

error对象是一个interface，其定义如下：

```go
type error interface {
    Error() string
}
```

我们可以利用errors包来实现自己的error对象，模拟golang的包中的类似行为，如：

```go
package main

import (
    "errors"
    "fmt"
)

func GetAge() (age int, err error) {
    return -1, errors.New("get age failed, error is not null")
}

func main() {
    _, err := GetAge()
    if err != nil {
        fmt.Println(err)
    }
}
```

程序最终会输出"get age failed, error is not null"。

## Defer，Panic，Recover

Defer含义是延迟执行的意思，通常用在函数中，待函数执行完成或失败进行一些善后的工作
