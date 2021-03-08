package util

func MaxString(a string, b string) string {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinString(a string, b string) string {
	if a > b {
		return b
	} else {
		return a
	}
}

func MaxLong(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinLong(a int64, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}
