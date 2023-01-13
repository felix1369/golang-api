package util

func CompareFloat(f1 float64, f2 float64, op string) bool {
	switch op {
	case "==":
		return f1 == f2
	case "=":
		return f1 == f2
	case ">=":
		return f1 >= f2
	case "<=":
		return f1 <= f2
	case ">":
		return f1 > f2
	case "<":
		return f1 < f2
	default:
		return false
	}
}
