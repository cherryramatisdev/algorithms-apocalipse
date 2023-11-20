# Array buffers / Ring buffers

- It's exactly like an array list with some differences

- Instead of using zero as the `head` and the length as `tail`, we'll define both head and tail as index based values that exist within our algorithm, that way everything outside the range defined between head and tail is considered null and we don't traverse at all.

- With this concept in mind, actions like `pop` on the beginning is done by simply adding 1 to the head, that way our contained space between head and tail it's smaller and the item outside these range is not even considered. The same way popping on the end is simply by subtracting 1 from the tail.

- The same way, `push` operation are just the opposite, we simply add 1 to the tail or subtract one from the head.

- All this actions are now `O(1)`

```
TODO: I didn't understand the whole ring stuff around this.
```
