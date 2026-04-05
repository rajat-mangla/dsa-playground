package bit_manipulation

const MOD = 1000000007

func modPow(a, b int) int {
	res := 1
	base := a % MOD

	for b > 0 {
		if b&1 == 1 {
			res = res * base % MOD
		}
		base = base * base % MOD
		b >>= 1
	}

	return res
}

// Fermat's little theorem states that if p is a prime and a is an integer not divisible by p, then a^(p-1) ≡ 1 (mod p).
func modInverse(a int) int {
	result := 1
	for b := MOD - 2; b > 0; b >>= 1 {
		if b&1 == 1 {
			result = result * a % MOD
		}
		a = a * a % MOD
	}
	return result
}
