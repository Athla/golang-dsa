// Divide and Conquer
// Algo logic:
// - Select a Pivot (any point)
// - Any value that is less than or equal to the pivot put it in the first position of the arr
// - All the ones that are bigger than the array goes to the next side of it
// - Everything in one side will be <= and the other side >, considered a weak sort
// - Split between the pivot greater and pivot lower and repeat the process, picking a new pivot
// - Pick everything around but not the pivot
// - The array is sorted
// - One side will be the Low and the other one will be the High
// - Both sides are inclusive to avoid problems
// -- It kindof becomes a tree due  the splitting in two different binary sides
// -- Binary Tree
// It keeps diving in half and half and half, calling itself continuosly, like mitosis
// Finally, it sort the nodes itself, creating dead leaves when it's len==1
// - It has O(n log n) in the average case, might change with the pivot point
// - - Reverse sorted binary arrays are the achilles heels of it, becoming n^2
// - - Worst average case goes between [O(nlogn) ~ O(n^2)]

function partition(arr: number[], lo: number, hi: number): number {
    const pivot = arr[hi];

    let idx = lo - 1;

    // [8, 7, 6, 4, 5]
    for (let i = lo; i < hi; i++) {
        // arr[i] == 4
        if (arr[i] <= pivot) {
            idx++;
            // tmp = 4, idx = 1
            const tmp = arr[i];
            // arr[3] == 8
            arr[i] = arr[idx];
            // arr[0] == 4
            arr[idx] = tmp;

        }
    }
    // [4, 7, 6, 8, 5]

    // idx == 1
    idx++;
    //arr[4] = 7
    arr[hi] = arr[idx];
    //arr[1] = 5
    arr[idx] = pivot;
    // [4, 5, 6, 8, 7]
    //
    return idx;
}

function qs(arr: number[], lo: number, hi: number) {
    if (lo >= hi) {
        return;
    }

    const pivotIdx = partition(arr, lo, hi)

    qs(arr, lo, pivotIdx - 1)
    qs(arr, pivotIdx + 1, hi)
}
export default function quick_sort(arr: number[]): void {
    qs(arr, 0, arr.length - 1);
}
