# Maze solver

- This is a `path finding` algorithm

- Assuming we have an array with the following content:

```
[
  "#####E#",
  "#     #",
  "#S#####"
]
```

- Goal: We need to make `S` find `E`

- Our base case will be:
  - It's a wall (`#`)
  - It's off the map
  - You're on the target (the `E`)
  - If we have seen this specific tile, that way we don't go through the same position twice
- Our recurse will be essentially try to go to any location
