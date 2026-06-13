#### Pre-requisite: Binary Trees and Recursion.

## Binary Tree Structure:
A Binary Search Tree (BST) is a specialized form of a binary tree where:

- **Node Structure:** Each node in a BST contains a value and has up to two children: a left child and a right child. These children are also nodes, and the structure continues recursively.
- **Left and Right Subtrees:** The left child of a node holds values less than the node’s value, while the right child holds values greater than the node’s value. This property ensures that the tree is organized in a way that makes searching efficient.
- **Height and Depth:** The height of a node is the number of edges on the longest path from the node to a leaf. The depth of a node is the number of edges from the root to the node. In a well-balanced BST, the height is minimized to approximately log2N, where N is the number of nodes, ensuring efficient operations.
- **Binary Tree Property:** Each node in a BST follows the binary tree property: having at most two children. This binary structure is fundamental to the BST’s efficiency in searching, insertion, and deletion operations.
- **Subtree Property:** Each subtree of a BST is also a BST. This means that every node in the left subtree is less than the root node, and every node in the right subtree is greater, maintaining the BST property throughout the tree.

### Balanced Variants of BST
To ensure optimal performance, several balanced variants of BSTs are used, such as:

- **AVL Tree:** A self-balancing BST where the height difference between left and right subtrees (balance factor) of any node is at most 1. Rotations are used to maintain balance after insertions and deletions.
- **Red-Black Tree:** A self-balancing BST where each node has an extra bit for color (red or black). Balancing is ensured by following specific rules about node coloring and rotations.

### Practical Applications

- **Database Indexing:** BSTs are used in database indexing to quickly locate records based on keys.
- **Symbol Tables:** Used in compilers and interpreters to manage variable names and function definitions.
- **Memory Management:** Implementing data structures like heaps for efficient memory allocation and deallocation.