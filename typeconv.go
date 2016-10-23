package typeconv

import "github.com/Sirupsen/logrus"

func Int(v interface{}) int64 {
	n1, ok := v.(uint64)
	if ok {
		return int64(n1)
	}

	n2, ok := v.(int64)
	if ok {
		return n2
	}

	return int64(v.(int))
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
	logrus.Panicf("ToIntTuple: can not convert: %T %v", v, v)
	return nil
}
