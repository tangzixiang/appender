package appender

import (
	"fmt"
)

type TODOList struct {
	Handler
}

type TODO struct {
	name string
}

func ExampleTODOListAppend() {
	list := TODOList{}
	list.Append() // not panic
	fmt.Printf("%+v\n", list.Values())

	list2 := TODOList{}
	list2.Append(TODO{"step 1"}, TODO{"step 2"}, TODO{"step 3"})
	fmt.Printf("%+v\n", list2.Values())

	// Output:
	// []
	// [{name:step 1} {name:step 2} {name:step 3}]
}

func ExampleTODOListAppendIf() {
	list := TODOList{}
	list.AppendIf(true)(TODO{"step 1"})
	list.AppendIf(false)(TODO{"step 2"})
	list.AppendIf(true)(TODO{"step 3"})
	fmt.Printf("%+v\n", list.Values())

	// Output:[{name:step 1} {name:step 3}]
}

func ExampleTODOListAppendWhen() {
	when := func(sources, appends []interface{}, arguments ...interface{}) bool {
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

	// source is empty list
	// appends is [ {name:step 2} {name:step 3}]
	// both source and appends not contain a item name is "step 1" so append success
	list := new(TODOList).AppendWhen(when, "step 1")(TODO{"step 2"}, TODO{"step 3"})
	fmt.Printf("%+v\n", list.Values())

	// source is empty list
	// appends is [ {name:step 4} {name:step 5} {name:step 6}]
	// appends contain a item name is "step 4" so append failed
	list2 := new(TODOList).AppendWhen(when, "step 4")(TODO{"step 4"}, TODO{"step 5"}, TODO{"step 6"})
	fmt.Printf("%+v\n", list2.Values())

	// source is [ {name:step 2} {name:step 3}]
	// appends is  [ {name:step 4} {name:step 5}]
	// both source and appends not contain a item name is "step 1" so append success
	list.AppendWhen(when, "step 1")(TODO{"step 4"}, TODO{"step 5"})
	fmt.Printf("%+v\n", list.Values())

	// Output:
	// [{name:step 2} {name:step 3}]
	// []
	// [{name:step 2} {name:step 3} {name:step 4} {name:step 5}]
}

func ExampleTODOListAppendFrom() {
	list := (&TODOList{}).AppendFrom(func(sources []interface{}, arguments ...interface{}) []interface{} {
		length := arguments[0].(int)
		ret := make([]interface{}, length)

		for i := 1; i <= length; i++ {
			ret[i-1] = TODO{fmt.Sprintf("step %v", i)}
		}

		return ret
	}, 2)

	fmt.Printf("%+v\n", list.Values())
	// Output:  [{name:step 1} {name:step 2}]
}
