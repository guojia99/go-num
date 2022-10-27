/*
Algorithm	Average	Worst case
Space	O(n)	O(n)
Search	O(log n)	O(log n)
Insert	O(log n)	O(log n)
Delete	O(log n)	O(log n)



Copy to References: https://en.wikipedia.org/wiki/B-tree
- Every node has at most m children.
- Every internal node has at least ⌈m/2⌉ children.
- Every non-leaf node has at least two children.
- All leaves appear on the same level and carry no information.
- A non-leaf node with k children contains k−1 keys.
*/

package btree

type BTree struct {
}
