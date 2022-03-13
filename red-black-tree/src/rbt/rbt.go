package rbt

import (
    "fmt"
)

type Node struct {
    city_info City
    is_red bool
    left *Node
    right *Node
    parent *Node
}

type Tree struct {
    Root *Node
    ll bool
    rr bool
    lr bool
    rl bool
}

func SetRotations(tree *Tree) {
    tree.ll = false
    tree.rr = false
    tree.lr = false
    tree.rl = false
}

func rotateLeft(node *Node) {
    x := node.right
    y := node.left

    x.left = node
    node.right = y
    node.parent = x

    if y != nil {
        y.parent = node
    }
}

func rotateRight(node *Node) {
    x := node.left
    y := node.right

    x.right = node
    node.left = y
    node.parent = x

    if y != nil {
        y.parent = node
    }
}

func insertHelper(tree *Tree, root *Node, new_node *Node) {
    red_red_conflict := false

    // Normal BST insertion stuff plus conflict detection

    if tree.Root == nil {
        tree.Root = new_node
    } else {
        if new_node.city_info.id < root.city_info.id {
            if root.left == nil {
                root.left = new_node
            } else {
                insertHelper(tree, root.left, new_node)
                root.left.parent = root

                if root.city_info.id != tree.Root.city_info.id {
                    if root.is_red == true && root.left.is_red == true {
                        red_red_conflict = true
                    }
                }
            }
        } else {
            if root.right == nil {
                root.right = new_node
            } else {
                insertHelper(tree, root.right, new_node)
                root.right.parent = root

                if root.city_info.id != tree.Root.city_info.id {
                    if root.is_red == true && root.right.is_red == true {
                        red_red_conflict = true
                    }
                }
            }
        }

        // roatations

        if tree.ll == true {
            rotateLeft(root)
            root.is_red = false
            root.left.is_red = true
            tree.ll = false
        } else if tree.rr == true {
            rotateRight(root)
            root.is_red = false
            root.right.is_red = true
            tree.rr = false
        } else if tree.rl == true {
            rotateRight(root.right)
            root.right.parent = root
            rotateLeft(root)
            root.is_red = false
            root.left.is_red = true
            tree.rl = false
        } else if tree.lr == true {
            rotateLeft(root.left)
            root.left.parent = root
            rotateRight(root)
            root.is_red = false
            root.right.is_red = true
            tree.lr = false
        }

        if red_red_conflict == true {
            if root.parent.right == root {
                if root.parent.left == nil || root.parent.left.is_red == false {
                    if root.left != nil && root.left.is_red == true {
                        tree.rl = true
                    } else if root.right != nil && root.right.is_red == true {
                        tree.ll = true
                    }
                } else {
                    root.parent.left.is_red = false
                    root.is_red = false
                    if root.parent.city_info.id != tree.Root.city_info.id {
                        root.parent.is_red = true
                    }
                }
            } else {
                    if root.parent.right == nil || root.parent.right.is_red == false {
                        if root.left != nil && root.left.is_red == true {
                            tree.rr = true
                        } else if root.right != nil && root.right.is_red == true {
                            tree.lr = true
                        }
                    } else {
                        root.parent.right.is_red = false
                        root.is_red = false
                        if root.parent.city_info.id != tree.Root.city_info.id {
                            root.parent.is_red = false
                        }
                    }
            }
            red_red_conflict = true
        }
    }
}

func PrintInorder(root *Node) {
    if root == nil {
        return
    }
    PrintInorder(root.left)
    fmt.Println(root.city_info.name)
    PrintInorder(root.right)
}

func Insert(tree *Tree, row []string) {
    city := cityFactory(row)
//    fmt.Println(city.name)
    node := &Node{ city_info: *city, is_red: true, left: nil, right: nil, parent: nil }
//    fmt.Println(node.city_info.id)
    if tree.Root == nil {
        tree.Root = node
        tree.Root.is_red = false
    } else {
        insertHelper(tree, tree.Root, node)
    }
}
