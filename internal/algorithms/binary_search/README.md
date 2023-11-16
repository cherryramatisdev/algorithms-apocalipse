# Binary Search

The binary search algorithm is a improvement of the Linear Search algorithm,
essentially we try to **never scan the list**, that way we're avoiding the
`O(n)` time complexity at all costs.

The functionality of this algorithm is the following:

- Assuming the list is ordered, we jump to the middle of the list and check if the value is larger, lesser or equal to the value
- If it's larger, we jump to the middle of the half from the value to the end
- if it's lesser, we jump to the middle of the half from the value to the begin
- If it's equal, we found our value

- This happen until repeatedly until we found the value!

---

- Initials assumptions: The array **must** be ordered
- Time complexity: O(log n)

## Pseudo code

```
search(arr, low, high, needle)

middle = floor(low + (high - low) / 2)
value = arr[middle]

while (low < high):
    if value == needle:
        return true
    else if value > middle:
        low = middle + 1
    else:
        high = middle

return false
```

## QA

1. Why the while condition is `low < high` and not `low <= high`?

Because the low part is inclusive, but the high part is exclusive. So essentially if the low is equal to the high, the point is behind the actual position, so it's in a fundamentally wrong position.
