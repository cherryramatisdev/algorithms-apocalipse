# Linked lists

- This data structure is a `Node based data structure`
- Essentially is a list of nodes where one node points to the reference of the next node
- This is considered a singly linked lists because you can't walk backwards,
  you just go forward because a node only points to the next one

```
A -> B -> C -> D

Node<T>
  val: T
  next?: Node<T>
```

A doubly linked lists add a bi-directional behavior by appending a new property to the node:

```
A <-> B <-> C <-> D

Node<T>
  val: T
  next?: Node<T>
  prev?: Node<T>
```
