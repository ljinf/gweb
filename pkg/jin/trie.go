package jin

import (
	"strings"
)

type Tree struct {
	root *node //根节点
}

type node struct {
	uri     string     //uri  例如：/user/login
	segment string     //uri的一部分  例如： user
	child   []*node    //子节点
	handler HandleFunc //处理函数
}

func NewTree() *Tree {
	return &Tree{
		root: newNode("", "", nil),
	}
}

func newNode(uri string, segment string, handler HandleFunc) *node {
	return &node{
		uri:     uri,
		segment: segment,
		child:   make([]*node, 0),
		handler: handler,
	}
}

//从根节点匹配，并添加
func (t *Tree) Set(uri string, index int, handler HandleFunc) {
	segments := strings.Split(uri, "/")
	if index >= len(segments) {
		return
	}
	//根
	root := t.root
	var temp *node
	for _, child := range root.child {
		if child.segment == segments[index] {
			temp = child
			break
		}
	}

	if temp == nil {
		n := newNode("", segments[index], nil)
		root.child = append(root.child, n)
		temp = n
	}

	temp.insert(uri, index+1, handler)
}

//添加
func (n *node) insert(uri string, index int, handler HandleFunc) {
	segments := strings.Split(uri, "/")
	if index >= len(segments) {
		//历遍完以后，保存对应的处理函数
		n.handler = handler
		n.uri = uri
		return
	}
	var root *node
	for _, child := range n.child {
		if child.segment == segments[index] {
			root = child
			break
		}
	}

	if root == nil {
		child := newNode("", segments[index], nil)
		n.child = append(n.child, child)
		root = child
	}

	root.insert(uri, index+1, handler)
}

//匹配路由，返回处理函数
func (t *Tree) Get(uri string) (pattern string, handler HandleFunc) {
	if len(t.root.child) < 1 {
		return "", nil
	}

	segments := strings.SplitN(uri, "/", 2)
	//根
	var temp *node
	for _, child := range t.root.child {
		if child.segment == segments[0] {
			temp = child
			break
		}
	}

	if temp != nil {
		if node := temp.find(segments[1]); node != nil {
			return node.uri, node.handler
		}
	}

	return "", nil
}

func (n *node) find(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)

	if len(segments) == 1 {
		//叶节点
		if n.uri != "" {
			if n.segment == segments[0] || hasPrefix(n.segment, ":") {
				return n
			}
		} else {
			//  类似：/user  直接匹配第二层了
			matchs := n.filterNode(segments[0])
			if len(matchs) < 1 {
				return nil
			}
			for _, child := range matchs {
				if child.segment == segments[0] || hasPrefix(child.segment, ":") {
					return child
				}
			}
		}
		return nil
	}

	//匹配下一层
	matchs := n.filterNode(segments[0])
	if len(matchs) < 1 {
		return nil
	}
	for _, child := range matchs {
		temp := child.find(segments[1])
		if temp != nil {
			return temp
		}
	}

	return nil
}

//查找
func (n *node) filterNode(segment string) []*node {
	//含有通配符，其子节点全部符合
	if hasPrefix(n.segment, ":") || n.segment == segment {
		return n.child
	}

	nodes := make([]*node, 0)
	for _, child := range n.child {
		if child.segment == segment || hasPrefix(child.segment, ":") {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func hasPrefix(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}
