Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error

Интерфейс представляет собой структуру, которая хранит указатель на объект, которые реализует заданный интерфейс, а также указатель на объект, описывающий тип интерфейса и тип значения, на которое ссылается первый указатель. 

Интерфейс является nil, если оба указателя в него - nil. В данном случае err хранит в себе указатель на customError, поэтому не является nil.

```
