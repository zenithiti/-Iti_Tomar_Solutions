package algorithm

func FibonacciNumber(n uint64) uint64 {
	f := make([]uint64, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 1
	f[1] = 1
	for i := uint64(2); i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
func SumOfEvenFibonacciNumber(n uint64)  uint64{
	var prev, cur, sum uint64 = 0, 1, 0
	for cur+prev <= n {
		next := prev + cur
		if next%2 == 0 {
			sum += next
		}
		prev = cur
		cur = next
	}
	return sum
}