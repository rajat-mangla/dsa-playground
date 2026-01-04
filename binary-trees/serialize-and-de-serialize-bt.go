package binary_trees

import (
	"strconv"
	"strings"
)

//https://takeuforward.org/plus/dsa/problems/serialize-and-de-serialize-bt
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Data int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * func NewTreeNode(val int) *TreeNode {
 *     return &TreeNode{Data: val}
 * }
 */

func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	arr := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		str := "#"
		if curr != nil {
			str = strconv.Itoa(curr.Data)
			queue = append(queue, curr.Left, curr.Right)
		}
		arr = append(arr, str)
	}

	return strings.Join(arr, ",")
}

func Deserialize(data string) *TreeNode {
	if len(data) == 0 || data == "#" {
		return nil
	}

	arr := strings.Split(data, ",")
	val, _ := strconv.ParseInt(arr[0], 10, 32)
	root := &TreeNode{
		Data: int(val),
	}

	queue := []*TreeNode{root}
	idx := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		idx += 1
		if idx < len(arr) && arr[idx] != "#" {
			val, _ := strconv.ParseInt(arr[idx], 10, 32)
			curr.Left = &TreeNode{
				Data: int(val),
			}
			queue = append(queue, curr.Left)
		}

		idx += 1
		if idx < len(arr) && arr[idx] != "#" {
			val, _ := strconv.ParseInt(arr[idx], 10, 32)
			curr.Right = &TreeNode{
				Data: int(val),
			}
			queue = append(queue, curr.Right)
		}
	}

	return root
}
