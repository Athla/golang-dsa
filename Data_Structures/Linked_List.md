# What is it?

- It's a list, which means a collection of things.
- Has a flow of things and sequence of objects.
- Better for data insertion and data removal/popping

## It's types

- Singly:
    - It's made of Nodes:
        Data | Link
    - Can go Forward Only:
        A -> B -> C -> D
    - The representation of the above structure in a Singly List follows:
    A|MemAddr(B) -> B|MemAddr(C) -> C|MemAddr(D) -> D|MemAddr(Nil)

- Doubly:
    - Can go both ways:
        A <-> B <-> C <-> D
    - The representation of the above Doubly Linked List follows:
    MemAddr(Null)|A|MemAddr(B) <-> MemAddr(A)|B|MemAddr(C) <-> MemAddr(B)|C|MemAddr(D) <-> MemAddr(C)|D|MemAddr(Null)

- Circular:
    - Can be cyclic:
        A -> B -> C -> D -> A
    - It can be represented by the following:
        MemAddr(D)|A|MemAddr(B) -> MemAddr(A)|B|MemAddr(C) -> MemAddr(B)|C|MemAddr(D) -> MemAddr(C)|D|MemAddr(A)

### Rules of a Linked List
- It must have a some pointers:
    - HEAD: Reference to the starting point
    - TAIL: Reference to the final node
- The process of removal or insertion of data follows the logic bellow:
    - Break the current links between nodes
        *IF HEAD POINTS HEAD TO NEW HEAD NODE
    - IF insertion:
        - After breaking the links, the PREVIOUS will link to the NEW as it's next NODE
        - The NEXT node will link tothe NEW as it's PREVIOUS  (if a Doubly linked list)
        - Examples, we are ADDING X:
            - Singly:
                A|MemAddr(B) -> B|MemAddr(C) -> C|MemAddr(D) -> D|MemAddr(Nil)

                A|MemAddr(B) -> B|MemAddr(X) -> B|MemAddr(X) -> C|MemAddr(D) -> D|MemAddr(Nil)
            - Doubly:
                MemAddr(Null)|A|MemAddr(B) <-> MemAddr(A)|B|MemAddr(X) <-> MemAddr(B)|X|MemAddr(D) <-> MemAddr(X)|C|MemAddr(D) <-> MemAddr(C)|D|MemAddr(Null)
            
        - For the circular one, it keeps the same logic, but since the  change is at the TAIL NODE
