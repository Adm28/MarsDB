package consistenthashring

import (
  "fmt"
  "hash/crc32"
  "sort"
  "errors"
  "log"
)

/*
Todo
Create an interface with common methods so that it's easier to implement different hashing poilicies

List of Methods available after importing this package
This package contains the implementatiopn of Consistent Hashing
NewRing  - Allocates a  new ring and returns a pointer to that ring
NewNode  - Allocates a new Node and returns the pointer to that node
AddNode  - Adds the Node in the right position in the ring
RemoveNode - Delets the Node from the Ring
GetNode  - Hashes the key and returns the ip address of the Node to which the value has to be inserted

*/

type SearchIndexMessage struct {
  desc string;
  msgType int;
}
func generateHash(key* string) uint32{
	return crc32.ChecksumIEEE([]byte(*key))
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
	Nodes Nodes;
}

// NewNode creates a Node
func NewNode(id* string) *Node {
  return &Node{
	  ID : *id,
		HashID : generateHash(id)};
}
// NewRing initializes the configuration of the ring
func NewRing() *Ring{
  return &Ring {
    Nodes : Nodes{} }
}
func searchIndex(nodes Nodes,hashID* uint32)(int,*SearchIndexMessage) {
  if(len(nodes) == 0) {
    return -1,&SearchIndexMessage{"Empty Array",0};
  }
  ind := sort.Search(len(nodes), func(i int) bool { return nodes[i].HashID >= *hashID })
  if(ind == -1 || ind >= len(nodes)) {
    return -1,&SearchIndexMessage{"Greater than all the elements of the array",1};
  } 
  return ind,nil;
}

func (r* Ring) AddNode(id* string) {
  node := NewNode(id);
  idx,msg := searchIndex((*r).Nodes,&(*node).HashID);
  (*r).Nodes = append((*r).Nodes,*node)
  nodesSlice := (*r).Nodes[:]
  lenNodes := len(nodesSlice)
  if(msg==nil) {
    for i:=lenNodes-1;i>idx;i-- {
      nodesSlice[i] = nodesSlice[i-1] 
    }
    nodesSlice[idx] = *node
  } else if(msg.msgType == 0) {
      nodesSlice[0] = *node
  } else if(msg.msgType == 1) {
      nodesSlice[len((*r).Nodes)-1] = *node
  }
}

func (r* Ring) RemoveNode(node Node){ 
  idx,msg := searchIndex(r.Nodes,&node.HashID);
  lenNodes := len((*r).Nodes);
  if msg==nil && idx < lenNodes && (*r).Nodes[idx].HashID == (node).HashID{
    fmt.Println("Yes")
    for i:=idx+1;i<lenNodes;i++ {
      (*r).Nodes[i-1] = (*r).Nodes[i]
    }
    (*r).Nodes = (*r).Nodes[:lenNodes-1]
  }  
}
func (r *Ring) GetNode (key* string) (*Node,error) {
  hashedKey := generateHash(key)
  log.Printf("Key : %s Generated Hash %s",key,hashedKey)
  // fmt.Printf("GetNode- Calculated Hashed Value %d",hashedKey)
  idx,msg := searchIndex(r.Nodes,&hashedKey)
  nodesArr := (*r).Nodes[:]
  if msg==nil {
    return &nodesArr[idx],nil;
  } else if msg.msgType == 0 {
    return nil,errors.New("No Nodes in the Cluster ")
  } else {
    return &nodesArr[0],nil
  }
}



/*func main () {
  ring := NewRing();

  members := []string{"127.0.0.1","127.0.0.2","127.0.0.3","127.0.0.4"};
  
  for _,ipAddress := range members {
    ring.AddNode(&ipAddress)
  }
  fmt.Println(ring.Nodes)
  // key := "HellYeah...fuckingyes"
  // fmt.Println(ring.GetNode(&key))
  All_Nodes := ring.Nodes
  All_Nodes = All_Nodes[:1];

  ring.RemoveNode(ring.Nodes[0])
  ring.RemoveNode(ring.Nodes[0])
  //ring.RemoveNode(ring.Nodes[0])
}*/
