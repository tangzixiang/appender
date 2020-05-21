# Title

appender is a convenient append data tool

## Install

```
go get github.com/tangzixiang/appender
```

## Usage

```go
// custom list hand some append method by appender.Handler
type TODOList struct {
    appender.Handler
}

type TODO struct {
    name string
}

list := TODOList{}

// append direct
// [{name:step 1} {name:step 2} {name:step 3}]
list.Append(TODO{"step 1"}, TODO{"step 2"}, TODO{"step 3"}) 

// append when you sure
// [{name:step 1} {name:step 2} {name:step 3} {name:step 4}]
list.AppendIf(true)(TODO{"step 4"}) // append success

 // [{name:step 1} {name:step 2} {name:step 3} {name:step 4}]
list.AppendIf(false)(TODO{"step 5"}) // append failed

// apend coustom, list is cached data collection
// [{name:step 1} {name:step 2} {name:step 3} {name:step 4} 
// {name:step 1} {name:step 2} {name:step 3}]
list.AppendWhen(func(list,appendData []interface{}, arguments ...interface{}) bool {
    compStr :=arguments[0].(string)

    for _,d := range appendData {
        if compStr == d.(TODO).name {
            return false
        }
    }

    return true
}, "step 4")(TODO{"step 1"}, TODO{"step 2"}, TODO{"step 3"})

// or
// [{name:step 1} {name:step 2} {name:step 3} {name:step 4} 
// {name:step 1} {name:step 2} {name:step 3} {name:step 1} {name:step 2}]
list.AppendFrom(func(list []interface{}, arguments ...interface{}) []interface{} {
    length := arguments[0].(int)
    ret := make([]interface{}, length)

    for i := 1; i <= length; i++ {
        ret[i-1] = TODO{fmt.Sprintf("step %v", i)}
    }

    return ret
}, 2)
```

## Contributing

## License

MIT © tangzixiang
