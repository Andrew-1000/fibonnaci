package algorithms

type FibonacciIterator struct {
	current int64
	previous int64
}

func  NewFibonacciIterator() FibonacciIterator {
	return FibonacciIterator {
		current: 0,
		previous: 0,
	}
}

func (fi *FibonacciIterator) Next() bool {
	if fi.current == 0 && fi.previous == 0 {
		fi.current = 1
		return true
	}
	fi.previous, fi.current = fi.current, fi.previous+fi.current
	return true
}

func (fi FibonacciIterator) Value() int64 {
	if fi.current == 0 && fi.previous == 0 {
		fi.current =1
		return fi.previous
	}
	fi.previous, fi.current = fi.current, fi.previous+fi.current
	return fi.previous
}