package main

import (
  "fmt"
  "hash/crc32"
)

func generateHash(key string) uint32{
	return crc32.ChecksumIEEE([]byte(key))
}
// Node encapsulates its identifier and hashValue
type Node struct {
	ID string;
	HashID uint32;
}
// Nodes is array of Node
type Nodes []Node;

// Ring keeps track of the number of Nodes
type Ring struct {
	CountOfNodes uint32;
	Nodes Nodes;
}

// NewNode creates a Node
func NewNode(id string) *Node {
	return &Node{
		ID : id,
		HashID : generateHash(id)};
}
// NewRing initializes the configuration of the ring
func NewRing() *Ring{
  return &Ring {
    CountOfNodes : uint32(0),
	Nodes : Nodes{}}
}

func (r* Ring) AddNode(id string) {
    r.CountOfNodes+=1;
    node := NewNode(id);
    r.Nodes = append(r.Nodes,*node);
}

func main () {
  ring := NewRing();

  members := []string{"127.0.0.1","127.0.0.2","127.0.0.3","127.0.0.4"};
  
  for _,ipAddress := range members {
    ring.AddNode(ipAddress)
  }
  fmt.Println(ring.CountOfNodes)
  fmt.Println(ring.Nodes)
}
