package jmespath

import (
	"fmt"
	"reflect"
)

func validateArg(f string, arguments []interface{}, index int, expectedType reflect.Kind) (reflect.Value, error) {
	if index >= len(arguments) {
		return reflect.Value{}, fmt.Errorf(argOutOfBoundsError, f, index+1, len(arguments))
	}
	arg := reflect.ValueOf(arguments[index])
	if arg.Type().Kind() != expectedType {
		return reflect.Value{}, fmt.Errorf(invalidArgumentTypeError, f, index+1, expectedType.String())
	}
	return arg, nil
}