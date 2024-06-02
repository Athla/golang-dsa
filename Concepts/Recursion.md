# Recursion
- Something that keeps calling itself constantly.
- A function that calls itself until it's solved. Involves the definition of a Base Case -> Where all things meet.
```typescript
function foo(n:number): number {
    if (n ===1) {
        return 1;
    }   
    return  n + foo(n-1);
}
```
- The base case is where N == 1, summing returning itself, if not, recurse until this condition is met. That's the *BASE CASE*

### Some SUDO Code

```text
foo(n)
    if n = 1
        ret 1

    ret n + foo(n-1)
```

- A recursive function must have a return Addr, a return Val and Arguments
- - Addr: How it got here
- - Val: What is, in fact, returned
- - Args: Memory

Calling foo(5), we have:

    rA = foo(5), rV=5+, A=5
    rA = foo(4), rV=4+, A=4
    rA = foo(3), rV=3+, A=3
    rA = foo(2), rV=2+, A=2
    rA = foo(1), rV=1, A=1

And then it starts going back up

    rA = foo(1), rV=1, A=1
    rA = foo(2), rV=2+1, A=2
    rA = foo(3), rV=3+3, A=3
    rA = foo(4), rV=4+6, A=4
    rA = foo(5), rV=5+10, A=5

We first have gone down, then up, making it possible to break recursion in three steps.

# Recurse
## Pre:
    N+
## Recurse
    Call of the function   
## Post
    Do something after it

# Simple path finding algo to better understand it
## The MazeSolver
- What can be possibly be done to traverse a maze and find the exit?
- - Go up, go right, go down, go left
    Apply the base case for each direction.
    

- - Base cases:
    1. It's a Wall! -> Can't go there.
    2. Off the map! -> Also can't go there, return.
    3. It's the end! -> Congrats! Finished recursing.
    4. If we have seen it? -> If yes, don't go there yaboi.

