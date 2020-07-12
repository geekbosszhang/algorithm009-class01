package homework

/*
Bit operation
Base-10 time case:
	ans = ans * 10 + n % 10
	n /= 10
Base-2 time case:
	ans = ans * 2 + n % 2
	n /= 2
Or use bit operators
	ans = (ans<<1) | (n & 1) 
	or ans = (ans<<1) + (n & 1)
	n >>= 1
*/
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i:= 0; i < 32; i++ {
		ans=(ans << 1) | (num&1)
		num>>=1
	}
	return ans
 }