# Arrays algorithmic concepts
- This directory contains implementations of sorting algorithms 
- Each file focuses on a specific algorithm or operation, providing a clear and concise implementation.

1. **Bubble Sort** (`bubble_sort.go`):
   - Bubble Sort is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order. 
     - The pass through the list is repeated until the list is sorted.
   - Time Complexity: O(n^2)
   - Space Complexity: O(1)

2. **Selection Sort** (`selection_sort.go`):
   - Selection Sort is an in-place comparison sorting algorithm. It divides the input list into two parts: a sorted and an unsorted part. 
   - It repeatedly selects the smallest (or largest) element from the unsorted part and moves it to the end of the sorted part.
   - Time Complexity: O(n^2)
   - Space Complexity: O(1)

3. **Insertion Sort** (`insertion_sort.go`):
   - Insertion Sort is a simple sorting algorithm that builds the final sorted array one item at a time. 
   - It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.
   - Time Complexity: O(n^2)
   - Space Complexity: O(1)

4. **Merge Sort** (`merge_sort.go`):
   - Merge Sort is an efficient, stable, and comparison-based sorting algorithm. 
   - It divides the unsorted list into n sublists, each containing one element, and then repeatedly merges sublists to produce new sorted sublists until there is only one sublist remaining.
   - Time Complexity: O(n log n)
   - Space Complexity: O(n)

5. **Quick Sort** (`quick_sort.go`):
   - Quick Sort is an efficient sorting algorithm that uses a divide-and-conquer approach. 
   - It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays according to whether they are less than or greater than the pivot.
   - Time Complexity: O(n log n) on average, O(n^2) in the worst case
   - Space Complexity: O(log n) due to recursive stack space

6. **Heap Sort** (`heap_sort.go`):
   - Heap Sort is a comparison-based sorting algorithm that uses a binary heap data structure. 
   - It first builds a max heap from the input data, then repeatedly extracts the maximum element from the heap and reconstructs the heap until all elements are sorted.
   - Time Complexity: O(n log n)
   - Space Complexity: O(1)
