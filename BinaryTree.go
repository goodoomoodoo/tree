package tree

import "fmt"

// TreeNode is node for binary search tree
type TreeNode struct {
	left    *TreeNode
	right   *TreeNode
	element int
	height  int
}

// BinaryTree is the main struct of this class
type BinaryTree struct {
	root *TreeNode
	size int
}

/**
* Constructor of Binary Tree
* @param -
* @return new BinaryTree
 */
func NewBinaryTree() BinaryTree {
	return BinaryTree{ nil, 0 }
}

/**
* Private
* Return the height of the given node
* @param t *TreeNode
* @return int height of node
*/
func ( bst *BinaryTree ) height( t *TreeNode ) int {
	if t == nil {
		return 0
	}

	return t.height
}

func ( bst *BinaryTree ) Clear( ) {

	bst.root = bst.clear( bst.root )
}

func ( bst *BinaryTree ) clear( t *TreeNode ) *TreeNode {

	/**
	* =================== !! REMINDER !! ========================
	*
	* t, the parameter here is the pointer, not the struct itself;
	* hence, writing code below would not work:
	*
	* 	e.g.
	*	t = nil
	*
	* This would only assign the t pointer to nil and not the 
	* bst.root to nil. In order to actually change the bst.root, 
	* it has to be done in the form of "bst.root"!!!!!
	* Otw the only thing thats has been changed is POINTER!!!!!
	* 
	* P.S.
	* It is important to also remember that the assignment, such as 
	* "t.left = ..." found in remove() and insert(), implies this 
	* "(*t).left = ...". With such implication, it's easy to forget
	* that t itself is a variable for a pointer and nothing more
	*/

	return nil
}

/**
* Insert integer value to the tree
* @param val int (integer element)
* @return void
 */
func ( bst *BinaryTree ) Insert( val int ) {

	bst.root = bst.insert( bst.root, val )
}

/**
* Private
* Insert integer value to the tree with recursive call
* @param t *TreeNode, val int
* @return t *TreeNode
 */
func ( bst *BinaryTree ) insert( t *TreeNode, val int ) *TreeNode {

	if t == nil {

		t = &TreeNode{ nil, nil, val, 0 }

	} else {

		if val > t.element {

			t.right = bst.insert( t.right, val )

		} else if val < t.element {

			t.left = bst.insert( t.left, val )
		}

		// do nothing when equals
	}

	lh := bst.height( t.left )
	rh := bst.height( t.right )

	if lh > rh {

		t.height = lh + 1

	} else {

		t.height = rh + 1
	}

	return t
}

/**
* TODO
 */
func ( bst *BinaryTree ) Remove( val int ) bool {

	var isRemoved bool 

	bst.root, isRemoved = bst.remove( bst.root, val )

	return isRemoved
}

/**
* TODO
 */
func ( bst *BinaryTree ) remove( t *TreeNode, val int ) ( *TreeNode, bool ) {

	var s bool

	if t == nil {

		return nil, false
	}
	
	if t.element > val {

		t.left, s = bst.remove( t.left, val )

	} else if t.element < val {

		t.right, s = bst.remove( t.right, val )

	} else if t.right != nil && t.left != nil {

		t.element = bst.findMax( t.left ).element 
		t.left, s = bst.remove( t.left, t.element )

	} else {

		if t.right != nil {

			t = t.right

		} else {

			t = t.left 
		}

		s = true
	}

	return t, s
}

/**
* Find the given arguement in the tree and return true if exist
* @param val int
* @return bool True if exist else false
 */
func ( bst *BinaryTree ) Find( val int ) bool {

	return bst.find( bst.root, val )
}

/**
* Private
* Find the given arguement in the tree recursively
* @param t *TreeNode, val int
* @return bool True if exist else false
 */
func ( bst *BinaryTree ) find( t *TreeNode, val int ) bool {

	if t != nil {

		if t.element > val {

			return bst.find( t.left, val )

		} else if t.element < val {

			return bst.find( t.right, val )

		} else {

			return true
		}
	}

	return false
}

/**
* Returns the maximum value of the Binary Tree
* @param -
* @return int
 */
func ( bst *BinaryTree ) FindMax() int {

	maxNode := bst.findMax( bst.root )
	return maxNode.element
}

/**
* Return the TreeNode with the maximum value
* @param t *TreeNode
* @return *TreeNode
 */
func ( bst *BinaryTree ) findMax( t *TreeNode ) *TreeNode {

	for t.right != nil {
		t = t.right
	}

	return t
}

/**
* Return the TreeNode with the minimum value
* @param t *TreeNode
* @return *TreeNode
 */
func ( bst *BinaryTree ) FindMin() int {

	minNode := bst.findMin( bst.root )
	return minNode.element
}

/**
* Return the TreeNode with the minimum value
* @param t *TreeNode
* @return *TreeNode
 */
func ( bst *BinaryTree ) findMin( t *TreeNode ) *TreeNode {

	for t.left != nil {

		t = t.left
	}

	return t
}

/**
* Prints all element in ascending order
* @param -
* @return void
 */
func ( bst *BinaryTree ) Println() {

	bst.printTreeInLine( bst.root )
}

/**
* Private
* Iterates through and print each node element of tree inorder
* @param t *TreeNode
* @return void
 */
func ( bst *BinaryTree ) printTreeInLine( t *TreeNode ) {

	if t != nil {

		bst.printTreeInLine( t.left )
		bst.printTreeNode( t )
		bst.printTreeInLine( t.right )
	}
}

/**
* Private
* Print the information about the passed through node
* @param t *TreeNode
* @return void
 */
func ( bst *BinaryTree ) printTreeNode( t *TreeNode ) {

	fmt.Printf( "%5s: %4d %6s: %4d \n", "Data", t.element, "Height", t.height )
}
