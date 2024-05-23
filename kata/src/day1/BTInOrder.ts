function walk(curr: BinaryNode<number> | null, path: number[]): number[] {
    // Base case
    // Stop recursing if not
    if (!curr) {
        return path;
    }

    // recurse starting
    // pre - go left
    walk(curr.left, path)
    // mid - go node
    path.push(curr.value);
    // from left to right
    walk(curr.right, path)
    // post
    return path;
}
export default function in_order_search(head: BinaryNode<number>): number[] {
    return walk(head, [])
}
