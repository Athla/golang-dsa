# Heap // Priority Queue

- Binary tree where every chind and grand child is smaller(MaxHeap) or larger than the current node(MinHeap)
```
            50              MaxHeap (Has the highest priority)
         //    \\            
        71      100         Bigger than the MaxHeap, thus MinHeap (Second level priority)
      //||      || \\
    101 80      200 101     Bigger than the PreviousMinHeap, thus CurrentMinHeap (Third level priority)
```
- Always a complete tree, always filled in tree level without gaps, filled from left to right.
- The main operation is a swapping, making the lowest value (the max heap) is always at the top of the branch and the min heap always at the bottom
- Heapfy it up \\ Heapfy it down -> swapping operations to keep the tree order.
- It may be stored as a array list, that is ordered from the MaxHeap to the MinHeap
```
[50, 71, 100, 101, 80, 200, 101, 200] -> Values
[00, 01, 002, 003, 04, 005, 006, 007] -> Idx
OUT <- QUEUE <- IN <=> FIFO Structure, the first one to enter will be the first one to go out
Lowest idx means they are close to exiting the queue
```
## Parent & Child relationship
```
GOING DOWN
2*idx + 1 -> left hand side
2*idx + 1 -> right hand side

GOING UP
[(i - 1)/2]
```
