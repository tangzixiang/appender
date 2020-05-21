package appender

// Append is a func use to receive values
type Append func(values ...interface{})

// Caller is use to hand function When and function From
type Caller interface {
	When(when func(list []interface{}, arguments ...interface{}) bool, arguments ...interface{}) Append
	From(func(arguments ...interface{}) []interface{}) Append
}

// Handler use to  provide some method to push data and cache them
type Handler struct {
	list []interface{}
}

// Append data to cache list
func (a *Handler) Append(values ...interface{}) {
	a.list = append(a.list, values...)
}

// AppendIf return function to receive values then cache them if sure
func (a *Handler) AppendIf(sure bool) Append {
	return func(values ...interface{}) {
		if !sure {
			return
		}
		a.list = append(a.list, values...)
	}
}

// AppendWhen return function to receive values then cache them if when return true
// param when receive list and arguments,list is a collection which hand cache data,
// appendData is receive values from return function,argument is come from another param
// param arguments use to as param use for when
func (a *Handler) AppendWhen(when func(list, appendData []interface{}, arguments ...interface{}) bool, arguments ...interface{}) Append {
	return func(values ...interface{}) {
		if !when(a.list, values, arguments...) {
			return
		}
		a.list = append(a.list, values...)
	}
}

// AppendFrom use to offer a closure function decide what value need cache
// param from receive list and arguments,list is a collection which hand cache data,
// // appendData is receive values from return function,argument is come from another param
// param arguments use to as param use for when
func (a *Handler) AppendFrom(from func(list []interface{}, arguments ...interface{}) []interface{}, arguments ...interface{}) {
	list := from(a.list, arguments...)
	if len(list) == 0 {
		return
	}

	a.list = append(a.list, list...)
}

// Values get cache data collection
func (a *Handler) Values() []interface{} {
	if len(a.list) == 0 {
		return []interface{}{}
	}

	values := make([]interface{}, len(a.list))
	for i, l := range a.list {
		values[i] = l
	}

	return values
}
