//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var Global js.Value
var tzOffset int

func log(format string, args ...any) error {
	console := Global.Get(`console`)
	if console.Type() != js.TypeObject {
		return fmt.Errorf(`could not get console`)
	}
	console.Call(`log`, fmt.Sprintf(format, args...))
	return nil
}

func getTimezoneOffset() (int, error) {
	dateClass := Global.Get(`Date`)
	if dateClass.Type() != js.TypeFunction {
		return 0, fmt.Errorf(`Date is not a function: %v`, dateClass)
	}
	date := dateClass.New()
	if date.Type() != js.TypeObject {
		return 0, fmt.Errorf(`new Date() is not an object: %v`, date)
	}
	tzOffset := date.Call(`getTimezoneOffset`)
	if tzOffset.Type() != js.TypeNumber {
		return 0, fmt.Errorf(`getTimezoneOffset() did not return a number: %v`, tzOffset)
	}
	f := tzOffset.Float()
	return int(f), nil
}

func getElementById(id string) (js.Value, error) {
	doc := Global.Get(`document`)
	if doc.Type() != js.TypeObject {
		return js.Null(), fmt.Errorf(`"document" is not an object`)
	}
	if doc.Get(`getElementById`).Type() != js.TypeFunction {
		return js.Null(), fmt.Errorf(`"document.getElementById" is not a function`)
	}
	val := doc.Call(`getElementById`, id)
	if !val.Truthy() {
		return js.Null(), fmt.Errorf(`no such element %q`, id)
	}
	if val.Type() != js.TypeObject {
		return js.Null(), fmt.Errorf(`returned value was not an object`)
	}
	return val, nil
}

func setDisplay(el js.Value, show bool) error {
	style := el.Get(`style`)
	if !style.Truthy() || style.Type() != js.TypeObject {
		return fmt.Errorf(`no style object found`)
	}
	if show {
		style.Set(`display`, `block`)
	} else {
		style.Set(`display`, `none`)
	}
	return nil
}

func display(unit rune, multiplier, max, value int) error {
	for i := max; i > 0; i -= multiplier {
		id := fmt.Sprintf(`%s-%d-%d`, string(unit), multiplier, i)
		el, err := getElementById(id)
		if err != nil {
			return fmt.Errorf(`could not display element: %w`, err)
		}
		if err := setDisplay(el, value >= i); err != nil {
			return fmt.Errorf(`could not set display of %q: %w`, id, err)
		}
	}
	return nil
}

func tick() error {
	now := time.Now()
	s := now.Second()
	m := now.Minute()
	h := now.Hour()

	if err := display('s', 1, 4, s%5); err != nil {
		return fmt.Errorf(`couldn't set 1s block: %w`, err)
	}
	if err := display('s', 5, 10, s%15); err != nil {
		return fmt.Errorf(`couldn't set 5s block: %w`, err)
	}
	if err := display('s', 15, 45, s%60); err != nil {
		return fmt.Errorf(`couldn't set 15s block: %w`, err)
	}

	if err := display('m', 1, 4, m%5); err != nil {
		return fmt.Errorf(`couldn't set 1m block: %w`, err)
	}
	if err := display('m', 5, 10, m%15); err != nil {
		return fmt.Errorf(`couldn't set 5m block: %w`, err)
	}
	if err := display('m', 15, 45, m%60); err != nil {
		return fmt.Errorf(`couldn't set 15m block: %w`, err)
	}

	if err := display('h', 1, 2, h%3); err != nil {
		return fmt.Errorf(`couldn't set 1h block: %w`, err)
	}
	if err := display('h', 3, 9, h%12); err != nil {
		return fmt.Errorf(`couldn't set 3h block: %w`, err)
	}
	if err := display('h', 12, 12, h%24); err != nil {
		return fmt.Errorf(`couldn't set 12h block: %w`, err)
	}
	return nil
}

func main() {
	var err error
	done := make(chan struct{}, 0)
	Global = js.Global()
	tzOffset, err = getTimezoneOffset()
	if err != nil {
		log(`could not get timezone offset: %s`, err.Error)
		return
	}
	Global.Set(`tick`, js.FuncOf(func(this js.Value, args []js.Value) any {
		if err := tick(); err != nil {
			if eerr := log(err.Error()); eerr != nil {
				panic(err)
			}
		}
		return js.Null()
	}))
	<-done
}
