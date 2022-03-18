package rbt

import (
    "fmt"
    "strconv"
)

///////////////
// STRUCTS ///
/////////////

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

///////////////
// UTILITIES /
/////////////

func GetCityInfo(node *Node) (string) {
    return node.city_info.name
}

func get_sibling(node *Node) (*Node, string) {
    parent := node.parent
    if node.city_info.id > parent.city_info.id {
        sibling := parent.left
        direction := "L"
        return sibling, direction
    } else {
        sibling := parent.right
        direction := "R"
        return sibling, direction
    }
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

func find_node_helper(root *Node, id int) (*Node) {
    if root == nil || root.city_info.id == id {
        return root
    } else if id < root.city_info.id {
        return find_node_helper(root.left, id)
    } else {
        return find_node_helper(root.right, id)
    }
}

func find_node(tree *Tree, city_id_map map[string]int, city string) (*Node) {
    id := city_id_map[city]
    return find_node_helper(tree.Root, id)
}

func num_of_children(node *Node) (int) {
    if node == nil {
        return 0
    }
    if node.left != nil && node.right != nil {
        return 2
    } else if (node.left == nil && node.right != nil) || (node.left != nil && node.right == nil) {
        return 1
    }
    return 0
}

func find_in_order_succ(node *Node) (*Node) {
    right_node := node.right
    left_node := right_node.left

    if left_node == nil {
        return right_node
    }

    for left_node.left != nil {
        left_node = left_node.left
    }
    return left_node
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

///////////////
/// INSERT ///
/////////////

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
// DELETING //
/////////////

func remove_case_6(tree *Tree, node *Node) {

}

func remove_case_5(tree *Tree, node *Node) {

}

func remove_case_4(tree *Tree, node *Node) {

}

func remove_case_3(tree *Tree, node *Node) {

}

func remove_case_2(tree *Tree, node *Node) {
    parent := node.parent
    sibling, direction := get_sibling(node)
    if sibling.is_red == true && parent.is_red == false && sibling.left.is_red == false && sibling.right.is_red == false {
        if direction == "L" {
            rotate_right(tree, nil, sibling, parent, false)
        } else if direction == "R" {
            rotate_left(tree, nil, sibling, parent, false)
        }
        parent.is_red = true
        sibling.is_red = false
        remove_case_1(tree, node)
    }
    remove_case_3(tree, node)
}

func remove_case_1(tree *Tree, node *Node) {
    // recursively call other cases
    if tree.Root == node {
        node.is_red = false
        return
    }
    remove_case_2(tree, node)
}

func remove_leaf(node *Node) {
    if node.city_info.id > node.parent.city_info.id {
        node.parent.right = nil
    } else {
        node.parent.left = nil
    }
}

func remove_black_node(tree *Tree, node *Node) {
    remove_case_1(tree, node)
    remove_leaf(node)
}

func deleteHelper(tree *Tree, node *Node) {
    left_child := node.left
    right_child := node.right

    var not_nil_child *Node

//    fmt.Println(left_child.city_info.name)
//    fmt.Println(right_child.city_info.name)

    if left_child != nil {
        not_nil_child = left_child
    } else {
        not_nil_child = right_child
    }

    if node == tree.Root {
        if not_nil_child != nil {
            tree.Root = not_nil_child
            tree.Root.parent = nil
            tree.Root.is_red = false
        } else {
            tree.Root = nil
        }
    } else if node.is_red == true {
        if num_of_children(node) == 0 {
            remove_leaf(node)
        } else {
            fmt.Println("ERROR: Cannot have children")
        }
    } else {
        if num_of_children(right_child) > 0 || num_of_children(left_child) > 0 {
            fmt.Println("ERROR: Red child of black cannot have children")
        } else if not_nil_child != nil && not_nil_child.is_red == true {
            node.city_info = not_nil_child.city_info
            node.left = not_nil_child.left
            node.right = not_nil_child.right
        } else {
            remove_black_node(tree, node)
        }
    }
}

func Delete(tree *Tree, city_id_map map[string]int, city string) {
    remove_node := find_node(tree, city_id_map, city)

    if remove_node == nil {
        return
    }

    fmt.Println(num_of_children(remove_node))

    if num_of_children(remove_node) == 2 {
        succ := find_in_order_succ(remove_node)
        remove_node.city_info = succ.city_info
        remove_node = succ
    }

    fmt.Println(remove_node.city_info.name)

    deleteHelper(tree, remove_node)
}

///////////////
// PRINTING //
/////////////

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
