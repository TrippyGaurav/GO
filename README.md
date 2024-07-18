# Go Basics

This README covers the basics of Go, including user input, conversion, time handling, pointers, arrays, slices, maps, structs, conditionals, loops, and functions.

## Table of Contents

1. [Installation](#installation)
2. [Hello World](#hello-world)
3. [User Input](#user-input)
4. [Conversion](#conversion)
5. [Time](#time)
6. [Pointers](#pointers)
7. [Arrays](#arrays)
8. [Slices](#slices)
9. [Maps](#maps)
10. [Structs](#structs)
11. [Conditionals](#conditionals)
12. [Loops](#loops)
13. [Functions](#functions)

## Installation

To install Go, follow the instructions on the official [Go installation page](https://golang.org/doc/install).

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## User Input

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello, ", name)
}
```

## Conversion

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str string = "123"
    var num int
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(num)
    }
}
```

## Time

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("Current time: ", now)

    formattedTime := now.Format("2006-01-02 15:04:05")
    fmt.Println("Formatted time: ", formattedTime)
}
```

## Pointers

```go
package main

import "fmt"

func main() {
    var a int = 10
    var ptr *int = &a

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", ptr)
    fmt.Println("Value at ptr:", *ptr)
}
```

## Arrays

```go
package main

import "fmt"

func main() {
    var arr [5]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    arr[3] = 4
    arr[4] = 5

    fmt.Println("Array:", arr)
}
```

## Slices

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice:", slice)

    slice = append(slice, 6, 7)
    fmt.Println("Updated Slice:", slice)
}
```

## Maps

```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    m["one"] = 1
    m["two"] = 2

    fmt.Println("Map:", m)
}
```

## Structs

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    var p Person
    p.Name = "John"
    p.Age = 30

    fmt.Println("Person:", p)
}
```

## Conditionals

### If-Else

```go
package main

import "fmt"

func main() {
    x := 10

    if x < 0 {
        fmt.Println("x is negative")
    } else if x == 0 {
        fmt.Println("x is zero")
    } else {
        fmt.Println("x is positive")
    }
}
```

### Switch-Case

```go
package main

import "fmt"

func main() {
    x := 2

    switch x {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    default:
        fmt.Println("Other")
    }
}
```

## Loops

### For Loop

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

### While Loop (for loop as a while loop)

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

### Range Loop

```go
package main

import "fmt"

func main() {
    nums := []int{2, 4, 6, 8}

    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

## Functions

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(3, 4)
    fmt.Println("Sum:", sum)
}
```

---

This guide provides a foundational understanding of Go programming concepts. For more in-depth knowledge, refer to the [official Go documentation](https://golang.org/doc/).
