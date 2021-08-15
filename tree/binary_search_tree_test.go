package tree

import (
	"arithmetic/tool"
	"reflect"
	"testing"
)

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

func TestMinDiffInBST(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   49,
				Left:  nil,
				Right: nil,
			},
		},
	}

	expect := 1
	res := minDiffInBST(root)
	if res != expect {
		t.Errorf("expect is %v, but get %v", expect, res)
	}

	expect2 := 1
	res2 := minDiffInBST(root2)
	if res2 != expect2 {
		t.Errorf("expect2 is %v, but get %v", expect, res2)
	}
}

func Test_recoverTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	expect := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	recoverTree(root)
	if !reflect.DeepEqual(root, expect) {
		t.Errorf("expect: %v, but get: %v", tool.ToString(expect), tool.ToString(root))
	}
}

func Test_deleteNode(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	expect := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := deleteNode(root, 3)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("expect: %v, but get: %v", tool.ToString(expect), tool.ToString(res))
	}
}
