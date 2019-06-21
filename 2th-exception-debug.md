# Exception & Debug

## 错误处理

在golang中见过最多的错误处理是error对象和nil的比较：

```go
f, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
```

这要求函数调用返回error对象，error对象是一个interface，其定义如下：

```go
type error interface {
    Error() string
}
```

我们可以利用errors包来模拟golang的包中的类似行为，如：

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

## try/catch

## Defer，Panic
