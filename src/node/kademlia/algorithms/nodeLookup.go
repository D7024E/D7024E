package algorithms

import (
	"D7024E/config"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"D7024E/node/kademlia/kademliaSort"

	"fmt"
	"math/rand"
	"sort"
	"sync"
)

func NodeLookup(destNode id.KademliaID) []contact.Contact {
	batch := kademlia.GetInstance().RoutingTable.FindClosestContacts(&destNode, config.Alpha)
	// batch = append(batch, kademlia.GetInstance().Me)
	return NodeLookupRec(destNode, batch)
}

func NodeLookupRec(destNode id.KademliaID, batch []contact.Contact) []contact.Contact {
	rt := bucket.GetInstance()
	me := contact.GetInstance()
	alpha := config.Alpha

	if len(batch) == 0 {
		batch = rt.FindClosestContacts(&destNode, alpha)
	}

	var newBatch [][]contact.Contact
	// For each of the alpha nodes in "batch", send a findNode RPC and append the result to "newBatch"
	for i := 0; i < len(batch); i++ {
		var kN []contact.Contact
		kN, err := rpc.FindNodeRequest(*me, batch[i], destNode)
		if err != nil {
			log.ERROR("%v", err)
		} else {
			newBatch = append(newBatch, kN)
		}
	}

	// Convert the contact batch into a single slice.
	updatedBatch := mergeBatch(newBatch)

	// Calculate the distance to each node in the batch and remove duplicates.
	updatedBatch = getAllDistances(*me.ID, updatedBatch)
	updatedBatch = removeDuplicates(updatedBatch)

	// Sort the cleaned batch and extract the alpha closest nodes.
	updatedBatch = kademliaSort.SortContacts(updatedBatch)
	updatedBatch = removeDeadNodes(updatedBatch, rpc.Ping)

	if len(updatedBatch) >= alpha {
		updatedBatch = updatedBatch[:alpha]
	}

	var sameBatch bool = true
	for i := 0; i < min(len(batch), len(updatedBatch)); i++ {
		if !batch[i].GetDistance().Equals(updatedBatch[i].GetDistance()) {
			sameBatch = false
		}
	}
	if sameBatch {
		return updatedBatch
	} else {
		return NodeLookupRec(destNode, updatedBatch)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// Note that this merge do not take duplicates into account.
func mergeBatch(batch [][]contact.Contact) []contact.Contact {
	var mergedBatch []contact.Contact
	for i := 0; i < len(batch); i++ {
		mergedBatch = append(mergedBatch, batch[i]...)
	}
	return mergedBatch
}

func TestMergeBatch(n int) {
	var testSetA []contact.Contact
	var testSetB []contact.Contact
	for i := 0; i < n; i++ {
		randomContactA := contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}

		randomContactB := contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}

		randomContactA.SetDistance(id.NewRandomKademliaID())
		randomContactB.SetDistance(id.NewRandomKademliaID())

		testSetA = append(testSetA, randomContactA)
		testSetB = append(testSetB, randomContactB)
	}

	var testSet [][]contact.Contact
	testSet = append(testSet, testSetA)
	testSet = append(testSet, testSetB)

	mergedTest := mergeBatch(testSet)

	fmt.Println("")
	fmt.Println("The test set is:")
	for i := 0; i < len(testSet); i++ {
		for j := 0; j < len(testSet[i]); j++ {
			fmt.Println(testSet[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("The merged test set is:")
	for i := 0; i < len(mergedTest); i++ {
		fmt.Println(mergedTest[i])
	}
}

// Given two kademlia ids' returns the distance between them
func getDistance(nodeA id.KademliaID, nodeB id.KademliaID) id.KademliaID {
	return *nodeA.CalcDistance(&nodeB)
}

func TestDistance() {
	nodeA := id.NewRandomKademliaID()
	nodeB := id.NewRandomKademliaID()

	fmt.Println("The test nodes are:", *nodeA, *nodeB)

	fmt.Println("The distance from node A to node B is:", getDistance(*nodeA, *nodeB))

	fmt.Println("The distance from node A to itself is:", getDistance(*nodeA, *nodeA))
}

// Updates the distances in "batch" to be the distances to the current node then returns the new batch.
func getAllDistances(me id.KademliaID, batch []contact.Contact) []contact.Contact {
	for i := 0; i < len(batch); i++ {
		relativeDistance := getDistance(*batch[i].ID, me)
		batch[i].SetDistance(&relativeDistance)
	}
	return batch
}

func TestAllDistances(n int) {
	var testSet []contact.Contact
	for i := 0; i < n; i++ {
		randomContact := contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}
		randomContact.SetDistance(id.NewRandomKademliaID())
		testSet = append(testSet, randomContact)
	}
	me := id.NewRandomKademliaID()
	testSet = getAllDistances(*me, testSet)
	fmt.Println("")
	fmt.Println("The base node is:", *me)
	for i := 0; i < len(testSet); i++ {
		fmt.Println("node-id is", *testSet[i].ID, "distance in base 10 is -", *testSet[i].GetDistance())
	}
}

func removeDuplicates(batch []contact.Contact) []contact.Contact {
	var cleanedBatch []contact.Contact
	// For each element in batch, check if the distance already exists in cleanedBatch.
	for i := 0; i < len(batch); i++ {
		var dupe bool = false
		currentID := batch[i].ID
		for j := 0; j < len(cleanedBatch); j++ {
			if currentID == cleanedBatch[j].ID {
				dupe = true
			}
		}
		// If no match is found, we add the i:th contact from batch to cleanedBatch.
		if !dupe {
			cleanedBatch = append(cleanedBatch, batch[i])
		}
	}
	return cleanedBatch
}

func TestDupe(n int, dupes int) {
	var testSet []contact.Contact
	for i := 0; i < dupes; i++ {
		randomContact := contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}
		randomContact.SetDistance(id.NewRandomKademliaID())
		testSet = append(testSet, randomContact)
		testSet = append(testSet, randomContact)
	}
	for i := 0; i < n; i++ {
		randomContact := contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}
		randomContact.SetDistance(id.NewRandomKademliaID())
		testSet = append(testSet, randomContact)
	}
	// fmt.Println("")
	// fmt.Println("")
	// fmt.Println("The test set is:")
	for i := 0; i < len(testSet); i++ {
		fmt.Println(testSet[i])
	}
	// fmt.Println("There are", len(testSet), "elements in the list")
	// fmt.Println("")
	// fmt.Println("")
	testSet = removeDuplicates(testSet)
	fmt.Println("The test set after running removeDuplicates is:")
	for i := 0; i < len(testSet); i++ {
		fmt.Println(testSet[i])
	}
	// fmt.Println("There are", len(testSet), "elements in the list")
}

// Takes a list of contacts, and a function as arguments. The function should be Ping() or a test function.
func removeDeadNodes(batch []contact.Contact, fn func(contact.Contact, contact.Contact) bool) []contact.Contact {
	me := contact.GetInstance()

	var deadNodes []int
	var wg sync.WaitGroup

	for i := 0; i < len(batch); i++ {
		wg.Add(1)
		n := i
		go func() {
			alive := fn(*me, batch[n])
			if !alive {
				deadNodes = append(deadNodes, n)
			} else {
				AddContact(batch[n])
			}
			wg.Done()
		}()
	}

	wg.Wait()

	sort.Ints(deadNodes)

	for i := 0; i < len(deadNodes); i++ {
		idx := deadNodes[i] - i
		if idx != len(batch)-1 {
			batch = append(batch[:idx], batch[idx+1:]...)
		} else {
			batch = batch[:idx]
		}
	}

	return batch
}

func TestRemoveDeadNodes(n int) {
	var testSet []contact.Contact
	for i := 0; i < n; i++ {
		randomContact := contact.Contact{
			ID:      id.NewRandomKademliaID(),
			Address: id.NewRandomKademliaID().String(),
		}
		testSet = append(testSet, randomContact)
	}

	fmt.Println("")
	fmt.Println("The test set is", len(testSet), "randomly generated contacts.")
	fmt.Println("")

	for i := 0; i < len(testSet); i++ {
		fmt.Println(testSet[i])
	}

	fmt.Println("")
	fmt.Println("")

	var resSet []contact.Contact = removeDeadNodes(testSet, FakePing)

	fmt.Println("")
	fmt.Println("After using remove nodes on them with a fake ping function, which has 50% chance to return a fail.")
	fmt.Println("")

	for i := 0; i < len(resSet); i++ {
		fmt.Println(resSet[i])
	}

}

func FakePing(nodeA contact.Contact, nodeB contact.Contact) bool {
	if rand.Float32() < float32(0.5) {
		return true
	} else {
		return false
	}
}
