package main

import (
//	"fmt"
//	"crypto/md5"	
	//"strings"
)

// Represents the Member node in a conisistent hash ring
type Member interface {
	String() string;
}

// Holds Information about the Members of the Nodes in the consistent hash circle
type ConsistentHash struct {	
	sortedNodes []string;
	ring map[string] uint64;
	countNodes uint64;	
}

 	
