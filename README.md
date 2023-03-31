# go-programming-crash-course
Repo for course: https://www.udemy.com/course/go-programming-language-crash-course

## Details:

### variable declaration

```
var variableName string
variableName = "some string"
```

Shorthand that infers variable type

```
variableName := "some string"
```

### initialize a module

```go mod init moduleName```

### build to executable

```
go build -o name main.go
go build -o name.exe main.go
```

### declaring a const

```
const variableName = "Some constant"
```

### string concatenation for printing

```
fmt.Println("Now multiply the result by", secondNumber, prompt)
```

### varible name to access level

If a variable has its first letter as a capital letter, e.g: `var PublicVar` then it is publicly accessible, otherwise, e.g: `var privateVar` it is only accessible inside the package where it was declared
Same applies for functions
