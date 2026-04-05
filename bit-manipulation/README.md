**Bit manipulation interesting finds**

Fermat's litte theorem
```
a^(p-1) ≡ 1 (mod p)
```
Where p is a prime number and a is an integer not divisible by p.
This can be used to find the modular inverse of a number a modulo p, which is given by:
```
a^(-1) ≡ a^(p-2) (mod p)
```
This is particularly useful in competitive programming when you need to perform division in modular arithmetic, as it allows you to compute the modular inverse efficiently using exponentiation by squaring.
