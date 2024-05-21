# What are **Two Pointers
- It's a technic that does set two variables pointing to different points of a structure, normally it's start (prt 1) and it's end (prt 2)


# When to use?
-  When it's needed to go over a array or a list of things, analyzing each element to other elements or similar things.

## Two Sum Problem, but with Two Pointers
```python
arr = [1, 2, 3, 4, 5]

target = 6

pointer_one = 0
pointer_two = len(arr) - 1

while pointer_one < pointer_two:
    total = arr[pointer_one] + arr[pointer_two]
    
    if total == target:
        return True
    # The sum is less than the target, then we walk up with the lower_pointer.
    elif total < target:
        pointer_one += 1
    # In this case the sum is more than the target, so we decrease our high pointer.
    else:
        pointer_two -= 1

    return false
```
