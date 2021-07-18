package tree

import "testing"

var (
	isBstNode = TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	notBstNode = TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	zeroValNode = TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
)

func TestIsValidBST(t *testing.T) {
	isTrue := IsValidBST(&isBstNode)
	if isTrue != true {
		t.Errorf("expect: true, but get: %v", isTrue)
	}
	isFalse := IsValidBST(&notBstNode)
	if isFalse != false {
		t.Errorf("expect: false, but get: %v", isFalse)
	}
	zeroVal := IsValidBST(&zeroValNode)
	if zeroVal != true {
		t.Errorf("expect: true, but get: %v", zeroVal)
	}
}
