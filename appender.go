package appender

// Append is a func use to receive values
type Append func(appends ...interface{}) Appender

// Appender hand some append method
type Appender interface {
	// Append data to inner list
	Append(appends ...interface{}) Appender

	// AppendFrom use to offer a closure function decide what value need to append by closure function return values
	// first param 'from' receive sources and arguments,sources is a interface slice which hand received data,
	// argument is come from another param
	// second param arguments use to as param use for from
	AppendFrom(from func(sources []interface{}, arguments ...interface{}) []interface{}, arguments ...interface{}) Appender

	// AppendIf return function to receive values then append them if sure
	AppendIf(sure bool) Append

	// AppendWhen return function to receive values then append them if when return true
	// first param 'when' receive sources and append-values and arguments,sources is a interface slice which hand received data,
	// appends is receive values which need to append to inner list from return function,argument is come from another param
	// second param arguments use to as param use for when
	AppendWhen(when func(sources, appends []interface{}, arguments ...interface{}) bool, arguments ...interface{}) Append

	// Values return inner list copy
	Values() []interface{}
}

// Handler use to  provide some method to push data and cache them
type Handler struct {
	list []interface{}
}

// Append data to inner list
func (a *Handler) Append(appends ...interface{}) Appender {
	a.list = append(a.list, appends...)
	return a
}

// AppendIf return function to receive values then append them if sure
func (a *Handler) AppendIf(sure bool) Append {
	return func(appends ...interface{}) Appender {
		if !sure {
			return a
		}

		a.list = append(a.list, appends...)
		return a
	}
}

// AppendWhen return function to receive values then append them if when return true
// first param 'when' receive sources and append-values and arguments,sources is a interface slice which hand received data,
// appends is receive values which need to append to inner list from return function,argument is come from another param
// second param arguments use to as param use for when
func (a *Handler) AppendWhen(when func(sources, appends []interface{}, arguments ...interface{}) bool, arguments ...interface{}) Append {
	return func(values ...interface{}) Appender {
		if !when(a.list, values, arguments...) {
			return a
		}
		a.list = append(a.list, values...)
		return a
	}
}

// AppendFrom use to offer a closure function decide what value need to append by closure function return values
// first param 'from' receive sources and arguments,sources is a interface slice which hand received data,
// argument is come from another param
// second param arguments use to as param use for from
func (a *Handler) AppendFrom(from func(sources []interface{}, arguments ...interface{}) []interface{}, arguments ...interface{}) Appender {
	list := from(a.list, arguments...)
	if len(list) == 0 {
		return a
	}

	a.list = append(a.list, list...)
	return a
}

// Values return inner list copy
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
