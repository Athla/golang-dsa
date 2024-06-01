# If its not a priority queue, it's a trie (Retrieval Tree)
- Has a constant time lookup

## English Lang Trie
- Only has 26 chars
- Add word "cat"
            (26)

```
         (26)             
       //              
      (C) 
      || 
      (A)
      ||  \\ 
      (T)  (R)
      || \\  \\
      (S) (T) (D)
          ||
          (T)
          ||
          (L)
          ||
          (E)
```
### PSEUDO CODE
```
insertion(str)
            curr = head
            for c in str
                        if curr[c]
                                    curr = curr[c]
                        else
                                    node = cN()
                                    curr[c] node
                                    curr = node
            
            curr.IsWord = true

deletion()
```
