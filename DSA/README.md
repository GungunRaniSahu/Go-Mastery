# Data Structures and Algorithms in Go

This repository contains implementations of fundamental Data Structures and Algorithms (DSA) in Go. The code is organized into different categories for easy navigation.

## Folder Structure
```
DSA/                
│    ├── Sorting/          # Sorting algorithms
│    │    ├── Bubble_sort.go    # Bubble Sort implementation
│    │    ├── Quick_sort.go     # Quick Sort implementation
│    │    ├── Merge_sort.go     # Merge Sort implementation
│    │
│    ├── Searching/       # Searching algorithms
│    │    ├── binary_search.go  # Binary Search implementation
│    │    ├── linear_search.go  # Linear Search implementation
│    │
│    ├── stacks/          # Stack data structure
│    │    ├── stack.go          # Stack implementation
│    │
│    ├── queues/          # Queue data structure
│    │    ├── queue.go          # Queue implementation
│    │
│    ├── linked-lists/    # Linked Lists
│    │    ├── singly_linked_list.go  # Singly Linked List implementation
│    │    ├── doubly_linked_list.go  # Doubly Linked List implementation
│    │
│    ├── trees/           # Tree data structures
│    │    ├── binary_tree.go  # Binary Tree implementation
│    │    ├── bst.go          # Binary Search Tree (BST) implementation
│    │
│    ├── graphs/          # Graph algorithms
│    │    ├── bfs.go          # Breadth-First Search (BFS) implementation
│    │    ├── dfs.go          # Depth-First Search (DFS) implementation
```

## Sorting Algorithms
- **Bubble Sort**: A simple comparison-based sorting algorithm.
- **Quick Sort**: A divide-and-conquer sorting algorithm with an average time complexity of O(n log n).
- **Merge Sort**: A stable divide-and-conquer sorting algorithm.

## Searching Algorithms
- **Binary Search**: Efficient searching on a sorted list with O(log n) complexity.
- **Linear Search**: A basic search algorithm with O(n) complexity.

## Data Structures
- **Stacks**: A LIFO (Last In, First Out) data structure.
- **Queues**: A FIFO (First In, First Out) data structure.
- **Linked Lists**: Singly and Doubly Linked List implementations.
- **Trees**: Implementations of Binary Trees and Binary Search Trees.
- **Graphs**: Graph traversal using BFS and DFS.

## How to Run
To run any implementation, navigate to the respective directory and execute the Go file:
```sh
cd DSA/Sorting
go run bubble_sort.go
```





