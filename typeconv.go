package typeconv

import (
	"fmt"
	"log"
	"reflect"
)

func Int(v interface{}) int64 {
	if n, ok := v.(uint64); ok {
		return int64(n)
	}

	if n, ok := v.(int64); ok {
		return n
	}

	if n, ok := v.(int); ok {
		return int64(n)
	}

	if n, ok := v.(float64); ok {
		if float64(int64(n)) != n {
			log.Panicf("Int: can not convert %v to int64", n)
		}

		return int64(n)
	}
	log.Panicf("Int: can not convert: %T %v", v, v)
	return 0
}

func IntTuple(v interface{}) []int64 {
	if t, ok := v.([]int64); ok {
		return t
	}
	if t, ok := v.([]int); ok {
		ret := make([]int64, len(t))
		for i, v := range t {
			ret[i] = int64(v)
		}
		return ret
	}

	if t, ok := v.([]interface{}); ok {
		ret := make([]int64, len(t))
		for i, v := range t {
			ret[i] = Int(v)
		}
		return ret
	}
	log.Panicf("IntTuple: can not convert: %T %v", v, v)
	return nil
}

func Float(v interface{}) float64 {
	if f, ok := v.(float64); ok {
		return f
	}
	if f, ok := v.(float32); ok {
		return float64(f)
	}

	log.Panicf("Float: can not convert: %T %v", v, v)
	return 0.0
}

func FloatTuple(v interface{}) []float64 {
	if t, ok := v.([]float64); ok {
		return t
	}
	if t, ok := v.([]float32); ok {
		ret := make([]float64, len(t))
		for i, v := range t {
			ret[i] = float64(v)
		}
		return ret
	}

	if t, ok := v.([]interface{}); ok {
		ret := make([]float64, len(t))
		for i, v := range t {
			ret[i] = Float(v)
		}
		return ret
	}

	log.Panicf("FloatTuple: can not convert: %T %v", v, v)
	return nil
}

func String(v interface{}) string {
	return v.(string)
}

// try to convert value to target type, panic if fail
func Convert(val reflect.Value, targetType reflect.Type) reflect.Value {
	valType := val.Type()
	if valType.ConvertibleTo(targetType) {
		return val.Convert(targetType)
	}

	//fmt.Printf("Value type is %v, emptyInterfaceType is %v, equals %v\n", valType, emptyInterfaceType, valType == emptyInterfaceType)
	interfaceVal := val.Interface()

	switch realVal := interfaceVal.(type) {
	case float64:
		return reflect.ValueOf(realVal).Convert(targetType)
	case []interface{}:
		// val is of type []interface{}, try to convert to typ
		sliceSize := len(realVal)
		targetSlice := reflect.MakeSlice(targetType, 0, sliceSize)
		elemType := targetType.Elem()
		for i := 0; i < sliceSize; i++ {
			targetSlice = reflect.Append(targetSlice, Convert(val.Index(i), elemType))
		}
		return targetSlice
	}

	panic(fmt.Errorf("convert from type %v to %v failed: %v", valType, targetType, val))
}
