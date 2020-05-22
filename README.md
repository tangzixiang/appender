# appender

appender let you struct convenient append data 

## Install

```
go get github.com/tangzixiang/appender
```

## Usage

step 1: define a struct and embed appender.Handler

```go
// custom list hand some append method by appender.Handler
type TODOList struct {
    appender.Handler
}

type TODO struct {
    name string
}
```

step 2: append direct such as slice

```go
list := TODOList{}
list.Append(TODO{"step 1"}, TODO{"step 2"}) 
fmt.Printf("%+v\n", list.Values())
// Output: [{name:step 1} {name:step 2}]

// or all in one step
fmt.Printf("%+v\n", new(TODOList).Append(TODO{"step 1"}, TODO{"step 2"}).Values())
// Output: [{name:step 1} {name:step 2}]

// or
fmt.Printf("%+v\n", (&TODOList{}).Append(TODO{"step 1"}, TODO{"step 2"}).Values())
// Output: [{name:step 1} {name:step 2}]
```

or append data when you sure

```go
list := (&TODOList{}).Append(TODO{"step 1"}, TODO{"step 2"}) 

// append when you sure
fmt.Printf("%+v\n", list.AppendIf(true)(TODO{"step 3"}).Values())
// Output: [{name:step 1} {name:step 2} {name:step 3}]
```

not use appender to implment above mybe write down：

```go
var list []interface{}{TODO{"step 1"}, TODO{"step 2"}}
var verb bool

if(verb){
  list = append(list,TODO{"step 3"})
}
```

maybe you can use a callback to decide whether to append use AppendWhen

```go
// sources is a empty slice here
// append is [{name:step 3} {name:step 4} {name:step 5}]
// arguments contain a string "step 6"
when := func(sources,appends []interface{}, arguments ...interface{}) bool {
	// ensure compStr not in sources or appends item name
	compStr := arguments[0].(string)

	for _, d := range sources {
		if compStr == d.(TODO).name {
			return false
		}
	}

	for _, d := range appends {
		if compStr == d.(TODO).name {
			return false
		}
	}

    return true 
}

// both source and appends not contain "step 6" so step3/4/5 will append success
list := (&TODOList{}).Append(TODO{"step 1"}, TODO{"step 2"}).
AppendWhen(when, "step 6")(TODO{"step 3"}, TODO{"step 4"}, TODO{"step 5"})

fmt.Printf("%+v\n", list.Values())
// Output: [{name:step 1} {name:step 2} {name:step 3} {name:step 4} {name:step 5}]
```

or use a  closure to decide  append what

```go
// sources is a empty slice here
// arguments contain a num which is 2
(&TODOList{}).AppendFrom(func(sources []interface{}, arguments ...interface{}) []interface{} {
    length := arguments[0].(int)
    ret := make([]interface{}, length)

    for i := 0; i < length; i++ {
        ret[i] = TODO{fmt.Sprintf("step %v", i+1)}
    }

    return ret
}, 2)

fmt.Printf("%+v\n", list.Values())
// Output: [{name:step 1} {name:step 2}]
```

## Contributing

## License

MIT © tangzixiang
