## Linked-list

- A linked list is a linear data structure where each element is a separate object. 
- Each element (node) of a list is comprising of two items - the data and a reference to the next node. 
- The last node has a reference to null. 
- The entry point into a linked list is called the head of the list. 
- It should be noted that head is not a separate node, but the reference to the first node. 
- If the list is empty then the head is a null reference.

## Concepts

### Slow and Fast pointers
- Slow and fast pointers is a technique used to solve problems related to linked lists.
- The idea is to have two pointers, one moving at a slow pace (slow pointer) and the other moving at a fast pace (fast pointer).
- This technique is used to detect cycles in a linked list, find the middle of a linked list, and other problems related to linked lists.

#### Recommended Practice problems:
1. LC 2130: https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
2. 