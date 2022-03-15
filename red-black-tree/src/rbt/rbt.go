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

func recolor(tree *Tree, grandfather *Node) {
    grandfather.right.is_red = false
    grandfather.left.is_red = false

    if grandfather != tree.Root {
        grandfather.is_red = true
    }
    rebalance(tree, grandfather)
}

func update_parent(tree *Tree, node *Node, old_child *Node, new_parent *Node) {
    node.parent = new_parent

    if new_parent != nil {
        if new_parent.city_info.id > old_child.city_info.id {
            new_parent.left = node
        } else {
            new_parent.right = node
        }
    } else {
        tree.Root = node
    }
}

func rotate_right(tree *Tree, node *Node, parent *Node, grandfather *Node, recolor bool) {
    great_grandfather := grandfather.parent
    update_parent(tree, parent, grandfather, great_grandfather)

    old_right := parent.right
    parent.right = grandfather
    grandfather.parent = parent

    grandfather.left = old_right
    if old_right != nil {
        old_right.parent = grandfather
    }

    if recolor == true {
        parent.is_red = false
        node.is_red = true
        grandfather.is_red = true
    }
}

func rotate_left(tree *Tree, node *Node, parent *Node, grandfather *Node, recolor bool) {
    great_grandfather := grandfather.parent
    update_parent(tree, parent, grandfather, great_grandfather)

    old_left := parent.left
    parent.left = grandfather
    grandfather.parent = parent

    grandfather.right = old_left
    if old_left != nil {
        old_left.parent = grandfather
    }

    if recolor == true {
        parent.is_red = false
        node.is_red = true
        grandfather.is_red = true
    }
}

func rebalance(tree *Tree, node *Node) {
    parent := node.parent

    if parent == nil || parent.parent == nil || (node.is_red == false || parent.is_red == false) {
        return
    }

    //fmt.Println("Node: " + node.city_info.name + ", Parent: " + parent.city_info.name)

    grandfather := parent.parent
    var node_dir string // 0 - left, 1 - right
    var parent_dir string

    if parent.city_info.id > node.city_info.id {
        node_dir = "L"
    } else {
        node_dir = "R"
    }

    if grandfather.city_info.id > parent.city_info.id {
        parent_dir = "L"
    } else {
        parent_dir = "R"
    }

    var uncle *Node

    if parent_dir == "L" {
        uncle = grandfather.right
    } else {
        uncle = grandfather.left
    }

    summary_direction := node_dir + parent_dir

    if uncle == nil || uncle.is_red == false {
        if summary_direction == "LL" {
            rotate_right(tree, node, parent, grandfather, true)
        } else if summary_direction == "RR" {
            rotate_left(tree, node, parent, grandfather, true)
        } else if summary_direction == "LR" {
            rotate_right(tree, nil, node, parent, false)
            rotate_left(tree, parent, node, grandfather, true)
        } else if summary_direction == "RL" {
            rotate_left(tree, nil, node, parent, false)
            rotate_right(tree, parent, node, grandfather, true)
        } else {
            fmt.Println("Invalid direction")
        }
    } else {
        recolor(tree, grandfather)
    }
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

    //TODO Rebalance
    rebalance(tree, new_node)
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
