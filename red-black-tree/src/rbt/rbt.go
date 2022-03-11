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
}

func insertHelper(tree *Tree, root *Node, new_node *Node) {
    if tree.Root == nil {
        tree.Root = new_node
    } else {
        if new_node.city_info.id < root.city_info.id {
            if root.left == nil {
                root.left = new_node
            } else {
                insertHelper(tree, root.left, new_node)
            }
        } else {
            if root.right == nil {
                root.right = new_node
            } else {
                insertHelper(tree, root.right, new_node)
            }
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
