function walk(curr: BinaryNode<number> | null, path: number[]): number[] {
    // Base case
    // Stop recursing if not
    if (!curr) {
        return path;
    }

    // recurse starting
    // pre
    path.push(curr.value);
    // recurse
    // from left to right
    walk(curr.left, path)
    walk(curr.right, path)
    // post
    return path;
}
export default function pre_order_search(head: BinaryNode<number>): number[] {
    return walk(head, [])
}
