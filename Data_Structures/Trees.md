# Trees
- In the end, everything is a tree, which is a  graph, and so on.
- A tree has a root (A) and child nodes, it's the top most node.
- Height: longest path to the most child node, means that the tree does not need to be balanced.
- Binary Tree: A balanced tree with the same amount of childs in each level and side, at max two children for each node.
- General Tree: A tree with 0 or more Childreen.
- Binary Search Tree: A ordered Binary Tree.
-

            7
        23       3
    5      4  18     21

# Traverse

- Might be a Pre Order, InOrder and PostOrder
- Watch changes between them is  where the node is printed/visited

## Pre Order
- - Visit a Node
- - Recurse Left
- - Recurse Right
    ### Output
    *7*, 23, 5, 4, 3, 18, 21
## In Order
- - Recurse Left
- - Visit a Node
- - Recurse Right
    ### Output
    5, 23, 4, *7*, 17, 3, 21
## Post Order
- - Recurse
- - Recurse Right
- - Visit a Node
    ### Output
    5, 4, 23, 18, 21, 3, *7*

# Assintotic Analysis
- O(N)
