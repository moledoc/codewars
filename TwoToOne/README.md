## Two to one

Take 2 strings `s1` and `s2` including only letters from a to z.
Return a new **sorted** string, the longest possible, containing distinct letters - each taken only once - coming from `s1` or `s2`.

```sh
a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
TwoToOne(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
TwoToOne(a, a) -> "abcdefghijklmnopqrstuvwxyz"
```