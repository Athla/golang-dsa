# GRAPHS
- Everything is a graph. Even a graph.
- It's a series of nodes with connections
- All trees are graphs, even IRL

## Terminology
- Cicle: When the starting Node(x) is also the end when folloing the links. It's a complete cycle.
- Acyclic: Has no cycles.
- Connected: Every node connects to another node
- Directed: When there is a direction between connections, like arrows pointing
- Undirected: NOT directed
- Weighted: Edges haves weights associated with them, like maps
- Dag: Direct Acyclic Graph
### Impl Terms
- Node/Edges: points/vertex of a graph
- Edge: connection between two nodes
 

## Big O
- O(Vertices * Edges)

## Representation
```as
   0
  /|\
 1 | 2
  \|/
   3
```
## Adjacency List
- Node(0): B, C
- Node(1): A, D
- Node(2): A, D
 Node(3): B, C

## Adjacency Matrix
|   | 0 | 1 | 2 | 3 |
|---|---|---|---|---|
| 0 | 0 | 1 | 1 | 0 |
| 1 | 1 | 0 | 0 | 1 |
| 2 | 1 | 0 | 0 | 1 |
| 3 | 0 | 1 | 1 | 0 |

## Edge List
- (A, B)
- (A, C)
- (B, D)
- (C, D)

## Methods to go over a Graph
- DFS and BFS!

### BFS in the following
```
graph TD;
    0((0)) -->|1| 1((1));
    1((1)) -->|1| 0((0));
    0((0)) -->|4| 2((2));
    0((0)) -->|5| 3((3));
    2((2)) -->|2| 3((3));
    3((3)) -->|5| 4((4));
```

