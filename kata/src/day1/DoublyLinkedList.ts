type Node<T> = {
    value: T,
    prev?: Node<T>,
    next?: Node<T>,
}
export default class DoublyLinkedList<T> {
    public length: number;
    private head?: Node<T>;
    private tail?: Node<T>;



    constructor() {
        this.length = 0;
        this.head = undefined;
        this.tail = undefined;
    }
    // Adds at the start
    // G -> A
    // A -> G
    // Head -> G
    prepend(item: T): void {
        const node = { value: item } as Node<T>;

        this.length++
        if (!this.head) {
            this.head = this.tail = node;
            return
        }
        node.next = this.head;
        this.head.prev = node;
        this.head = node;

    }
    // B <-> C <-> D
    // Insert F at C IDX (between B and C)
    // B <-> F
    // F <-> C
    // B <-|-> C
    insertAt(item: T, idx: number): void {
        if (idx > this.length) {
            throw new Error("F")
        }
        this.length++
        if (idx === this.length) {
            this.append(item);
            return;
        } else if (idx === 0) {
            this.prepend(item);
            return;
        }

        let curr = this.getAt(idx) as Node<T>;

        const node = { value: item } as Node<T>;

        node.next = curr;
        node.prev = curr.prev;
        curr.prev = node;

        if (node.prev) {
            node.prev.next = curr;
        }

    }
    // Adds at the end
    // Add F at the end, After Z
    // Y <-> Z <-> F
    // Z <-> F
    // F <-> Z
    // Tail -> Z
    append(item: T): void {
        const node = { value: item } as Node<T>;

        this.length++
        if (!this.tail) {
            this.tail = this.head = node;
            return;
        }
        node.prev = this.tail;
        this.tail.next = node;
        this.tail = node;
    }
    // Redirect links
    // Y <-> Z <-> F
    // Y <-> Z <-> F && Y <-> Z
    // Y <-|-> F && F <-|-> Z
    // Y == NULL
    // To do that we have to check it there is a previous and if there is a next
    remove(item: T): T | undefined {
        let curr = this.head;
        for (let i = 0; curr && i < this.length; i++) {
            if (curr.value === item) {
                break
            }
            curr = curr.next;
        }

        if (!curr) {
            return undefined;
        }

        return this.removeNode(curr);
        // const node = { value: item } as Node<T>;
        // if (node.prev && node.next) {
        // }
    }
    get(idx: number): T | undefined {
        return this.getAt(idx)?.value;

    }
    removeAt(idx: number): T | undefined {
        const node = this.getAt(idx);
        if (!node) {
            return undefined
        }

        return this.removeNode(node)
    }

    private removeNode(node: Node<T>): T | undefined {
        this.length--;
        if (this.length === 0) {
            const out = this.head?.value;
            this.head = this.tail = undefined
            return out;
        }

        if (node.prev) {
            node.prev = node.next;
        }

        if (node.next) {
            node.next = node.prev;
        }

        if (node === this.head) {
            this.head = node.next;
        }
        if (node === this.tail) {
            this.tail = node.prev;
        }
        node.prev = node.next = undefined;
        return node.value;
    }

    private getAt(idx: number): Node<T> | undefined {
        let curr = this.head;
        for (let i = 0; curr && i < idx; i++) {
            curr = curr.next;
        }

        return curr;

    }
}
