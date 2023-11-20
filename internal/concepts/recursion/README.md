# Recursion

Recursion is a base concept of programming where your function execute itselfs basically.

1. It's always important to define a **base case** for your recursion, the base case is the condition that will break out recursion: `ALWAYS DEFINE A BASE CASE, DONT BREAK YOUR COMPUTER`
2. The forwarding function execution should have a slightly change to the initial parameter so the function actually proceed for something, like decreasing a number on every iteration of the recursion.
3. The recurse action can be broke down into three steps:
  1. Pre: Any action that happen **before the recurse** (on the `foo` case is the "n+")
  2. Recurse: The action recursion
  3. Post: Any action that happen **after the recurse** (we can for example instantiate a variable with the `n + foo(n - 1)` and do any other action)

```python
def foo(n):
  if n = 1: # base case
    return 1

  return n + foo(n - 1) # recursion with : slightly change to the initial parameter
```

## Execution of this algorithm

| rA    | rV | A |
|-------|----|---|
| foo 5 | 5+ | 5 |
| foo4  | 4+ | 4 |
| foo3  | 3+ | 3 |
| foo2  | 2+ | 2 |
| foo1  | 1  | 1 |

After this execution downwards, we go upwards replacing the + signs with the result

| rA    | rV   | A |
|-------|------|---|
| foo 5 | 5+10 | 5 |
| foo4  | 4+6  | 4 |
| foo3  | 3+3  | 3 |
| foo2  | 2+1  | 2 |
| foo1  | 1    | 1 |

