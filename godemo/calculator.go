package godemo

import "errors"

func Calculate(a *int, b *int, operation string) (int, error) {
	if a == nil || b == nil {
		return 0, errors.New("operands can's be NIL")
	}
	switch operation {
	case "addition":
		return *a + *b, nil
	case "multiplication":
		return *a * *b, nil
	case "division":
		return *a / *b, nil
	default:
		return 0, nil
	}

}
