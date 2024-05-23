function walk(curr: BinaryNode<number> | null, path: number[]): number[] {
    // Base case
    // Stop recursing if not
    if (!curr) {
        return path;
    }

    // recurse starting
    // pre - go left
    walk(curr.left, path)
    // mid - go right
    walk(curr.right, path)
    // post - go node
    path.push(curr.value);
    // post
    return path;
}
export default function post_order_search(head: BinaryNode<number>): number[] {
    return walk(head, [])
}
