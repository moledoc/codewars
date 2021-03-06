The ludic numbers are a set of numbers related to prime numbers, as they are generated by sieving. Below is the method to generate such numbers.

```
First take the natural numbers starting from 2.
[2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19 ...]
1 is the first ludic number -> Ludic = {1}.

Take the first element from the resulting array: 2.
Remove every 2nd indexed number.
[3, 5, 7, 9, 11, 13, 15, 17, 19 ...]
The first deleted number is 2 -> Ludic = {1, 2}.

Take the first element from the resulting array: 3.
Remove every 3rd indexed number.
[5, 7, 11, 13, 17, 19 ...]
The first deleted number is 3 -> Ludic = {1, 2, 3}.

Take the first element from the resulting array: 5.
Remove every 5th number.
[7, 11, 13, 19 ...]
The first deleted number is 5 -> Ludic = {1, 2, 3, 5}.

etc...
```

Write a function to return an integer: the sum of the first n ludic numbers.
You can assume that 1 ≤ n ≤ 10000.
