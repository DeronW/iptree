package iptree

import (
	"net"
)

type state int8
type IpType int8

type Node struct {
	/*
		0 bit: has IP exits
		1 bit: has empty
		2 bit: is range border, 0 means no, 1 means yes
		3 bit: border side, 0 mean left, 1 means right
		4 ~ 7 bits: prefix 0 ~ 15
	*/
	state    byte
	children map[byte]*Node
}

const (
	stateEmpty state = iota
	stateMixed
	stateFull
)

/*
create a new `IP Set` tree, there is no inital data
*/
func New() *Node {
	return &Node{}
}

/*
append ip into the tree, accept three kinds of ip
 1. single ip, like: 192.168.0.1, or ::
 2. CIDR ip, like: 192.168.0.0/24
 3. range of ip, a head one and a tail one, mean all ip between them,
    include them two. like: 1.1.1.1-2.2.2.2
*/
func (n *Node) Append(s string) error {
	// ip, err := Parse(s)
	// if err != nil {
	// 	return err
	// }
	// if ip.Type == IpTypeIP {
	// 	n.AppendIP(*ip.IP)
	// }
	// if ip.Type == IpTypeNet {
	// 	n.AppendNet(*ip.IPNet)
	// }
	// if ip.Type == IpTypeRange {
	// 	n.AppendRange(ip.Range[0], ip.Range[1])
	// }
	return nil
}

func (n *Node) Remove(s string) error {
	return nil
}

func (n *Node) level() int {
	return int(n.state & 0x0F)
}

func (n *Node) isFull() bool {
	return n.state&0x30 > 0
}

func (n *Node) setFull() {
	n.state = n.state&0x0f | 0x30
	n.children = nil
}

/*
after append or remove IP, recall the modified path, and
reset node, include reset state and remove node
*/
func (n *Node) recastRange(start [16]*Node, end [16]*Node) {

}

func (n *Node) recastIP(path [16]byte) {
	nodes := [16]*Node{}

	node := n
	for i := 0; i < 16; i++ {
		node := node.children[path[i]]
		nodes[i] = node
	}

	// for i := 15; i >= 0; i-- {
	// 	n := path[i]
	// 	if n == nil {
	// 		continue
	// 	}
	// 	if n.isFull() {
	// 		continue
	// 	}
	// 	if n.children == nil {

	// 	}
	// }
}

/*
append a single ip into set
*/
func (n *Node) AppendIP(ip net.IP) {
	path := [16]byte{}
	nodes := [16]*Node{}

	node := n
	for i, b := range ip.To16() {
		c, ok := node.children[b]
		if ok {
			// already in set
			if c.isFull() {
				return
			}
		} else {
			node.children[b] = &Node{}
			node = node.children[b]
		}

		nodes[i] = node
		path[i] = b
	}

	// recast current path
	for i := 14; i >= 0; i-- {
		fullChildren := false
		children := nodes[i].children
		if len(children) == 256 {
			for _, c := range children {
				if !c.isFull() {
					break
				}
			}
			fullChildren = true
		}
		if fullChildren {
			nodes[i].setFull()
		} else {
			break
		}
	}

}

/*
append a Net into set
*/
func (n *Node) AppendNet(ip net.IPNet) {

}

/*
append a range of ip into set
*/
func (n *Node) AppendRange(start net.IP, end net.IP) {
}

func (n *Node) RemoveIP(ip net.IP) {
	if !n.Has(ip) {
		return
	}
	node := n
	for _, b := range ip.To16() {
		c, ok := node.children[b]
		if ok {
			node = c
		} else {

			return
		}
	}
}
func (n *Node) RemoveNet(ip net.IPNet)               {}
func (n *Node) RemoveRange(start net.IP, end net.IP) {}

/*
tell if the ip is in set
*/
func (n *Node) Has(ip net.IP) bool {
	node := n
	for _, b := range ip.To16() {
		c, ok := node.children[b]
		if !ok {
			return false
		}
		if c.isFull() {
			return true
		}
		node = c
	}
	return false
}

/*
list all ips
*/
func (n *Node) Values() []string {
	ips := []string{}
	return ips
}
