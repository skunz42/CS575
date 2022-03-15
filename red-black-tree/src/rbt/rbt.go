package rbt

import (
    "fmt"
    "strconv"
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

func GetCityInfo(node *Node) (string) {
    return node.city_info.name
}

func insertHelper(tree *Tree, root *Node, new_node *Node) {
    // Normal BST insertion stuff plus conflict detection
    if new_node.city_info.id < root.city_info.id {
        if root.left == nil {
            root.left = new_node
            root.left.parent = root
        } else {
            insertHelper(tree, root.left, new_node)
        }
    } else {
        if root.right == nil {
            root.right = new_node
            root.right.parent = root
        } else {
            insertHelper(tree, root.right, new_node)
        }
    }
}

func Insert(tree *Tree, row []string) {
    city := cityFactory(row)
    node := &Node{ city_info: *city, is_red: true, left: nil, right: nil, parent: nil }
    if tree.Root == nil {
        tree.Root = node
        tree.Root.is_red = false
    } else {
        insertHelper(tree, tree.Root, node)
    }
}

///////////////
// PRINTING //
//////////////

func PrintInorder(root *Node) {
    if root == nil {
        return
    }
    PrintInorder(root.left)
    fmt.Println(root.city_info.name)
    PrintInorder(root.right)
}

func height(root *Node) (int) {
    if root == nil {
        return 0
    }
    lheight := height(root.left)
    rheight := height(root.right)

    if lheight > rheight {
        return lheight + 1
    }
    return rheight + 1
}

func printCurrentLevel(root *Node, level int, width *int) {
    if root != nil {
        if level == 1 {
            *width = *width + 1
            if root.parent != nil {
                fmt.Print(root.city_info.name + " (parent: " + root.parent.city_info.name + ") | ")
            } else {
                fmt.Print(root.city_info.name + " | ")
            }
        } else if level > 1 {
            printCurrentLevel(root.left, level-1, width)
            printCurrentLevel(root.right, level-1, width)
        }
    }
}

func PrintLevelOrder(root *Node) {
    colorRed := "\033[31m"
    colorReset := "\033[0m"

    h := height(root)
    width := 0
    for i := 1; i <= h; i++ {
        fmt.Print("| ")
        printCurrentLevel(root, i, &width)
        fmt.Println(string(colorRed), "(Level Width:" + strconv.Itoa(width) + ")", string(colorReset))
        width = 0
    }
}
