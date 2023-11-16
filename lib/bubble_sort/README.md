# Bubble sort

This is a **really** simple algorithm where you linearly walk through the list checking a simple condition; If the current number is larger than the next one, we swap the number positions.

After the first iteration of bubble sort, we for sure know that the highest number on this list is at the end right? Well, on the next time we run this sorting algorithm, we can run **without considering the last one**

After another iteration of bubble sort, we can do without considering the two last ones, and that progressively

Time complexity: O(N^2)
  - First iteration: N
  - Second iteration: N-1
  - Third iteration: N-2
  - Nth iteration: N - N + 1

## Pseudo code

```
for i..n:
    for j..n-1-i:
        if arr[i] > arr[j]:
            swap(i ,j)
```
