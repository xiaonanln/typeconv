package typeconv

import "github.com/Sirupsen/logrus"

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
			logrus.Panicf("Int: can not convert %v to int64", n)
		}

		return int64(n)
	}

	logrus.Panicf("Int: can not convert: %T %v", v, v)
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
	logrus.Panicf("IntTuple: can not convert: %T %v", v, v)
	return nil
}

func Float(v interface{}) float64 {
	if f, ok := v.(float64); ok {
		return f
	}
	if f, ok := v.(float32); ok {
		return float64(f)
	}

	logrus.Panicf("Float: can not convert: %T %v", v, v)
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

	logrus.Panicf("FloatTuple: can not convert: %T %v", v, v)
	return nil
}
