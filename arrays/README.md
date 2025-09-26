# Arrays algorithmic concepts
- This directory contains implementations of various algorithms and operations related to arrays. 
- Each file focuses on a specific algorithm or operation, providing a clear and concise implementation.

1. **K frequent Elements** (`top_k_frequent_elements.go`):
   - There are mostly two approaches to solve such problem:
     1. Bucketing sort to group elements by their frequency, then iterate from the highest frequency to get the top k elements.
        1. Time Complexity: O(n)
        2. Space Complexity: O(n)
     2. Using a min-heap (priority queue) to sort the elements.
        1. Time Complexity: O(n log n)
        2. Space Complexity: O(n)
     3. Using a min-heap (priority queue) to keep track of the top k elements.
        1. Time Complexity: O(n log k)
        2. Space Complexity: O(n)
     3. Using the Quickselect algorithm to partition the array and find the k most frequent elements.

[//]: # ()
[//]: # (2. **Two Sum** &#40;`two_sum.go`&#41;:)

[//]: # (   - This file contains an implementation of the Two Sum problem, where the goal is to find two numbers in an array that add up to a specific target sum.)

[//]: # (   - The function `TwoSum` takes an array of integers and a target sum as input and returns the indices of the two numbers that add up to the target sum.)

[//]: # (   - It uses a hash map to store the complement of each number and checks if the complement exists in the map.)

[//]: # (   - Time Complexity: O&#40;n&#41;)

[//]: # (   - Space Complexity: O&#40;n&#41;)

