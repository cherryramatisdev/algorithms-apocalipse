# Array list

- Essentially array list try to add the "growable" capability to plain arrays while maintaining a better performance than linked lists, since for every action on a linked list we need to traverse the whole structure (because we don't have the length).

- This is done by understanding two important concepts: "length" and "capability". Length is the actual number of alocated values on this array list and capability is the total size of the array that can be alocated.

- With this concept in mind is pretty constant to apply actions such as pop, get and push because since we're working with a **real** set of values (the length) we can go to any point of the structure like an array but we know for a fact that all those items are alocated values and not just an empty space, causing extra loops execution.

- A important situation is when we try to push over the capability of an array list, on this situation we create underneath a new array and move all the items to there with a capability + 1.
