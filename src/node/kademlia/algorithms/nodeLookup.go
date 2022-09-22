package kademlia

import (
	"D7024E/config"
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/log"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"
)

func NodeLookup(destNode id.KademliaID, batch []contact.Contact) (closest []contact.Contact) {
	rt := bucket.GetInstance()
	me := rt.Me

	if len(batch) == 0 {
		batch = rt.FindClosestContacts(&destNode, config.Alpha)
	}
	var newBatch [][]contact.Contact
	// For each of the alpha nodes in "batch", send a findNode RPC and append the result to "newBatch"
	for i := 0; i < len(batch); i++ {
		var kN []contact.Contact
		kN, err := rpc.FindNode(batch[i], destNode)
		if err != nil {
			log.ERROR("%v", err)
		} else {
			newBatch = append(newBatch, kN)
		}
	}

	// Convert the contact batch into a single slice.
	batch = mergeBatch(newBatch)

	// Calculate the distance to each node in the batch and remove duplicates.
	distBatch := getAllDistances(*me.ID, batch)
	cleanedBatch := removeDuplicates(distBatch)
	fmt.Print(cleanedBatch)

	// Sort the cleaned batch and extract the alpha closest nodes.

	return
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

	var mergedTest []contact.Contact
	mergedTest = mergeBatch(testSet)

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
		currentDist := batch[i].GetDistance()
		for j := 0; j < len(cleanedBatch); j++ {
			if currentDist == cleanedBatch[j].GetDistance() {
				dupe = true
			}
		}
		// If no match is found, we add the i:th contact from batch to cleanedBatch.
		if dupe == false {
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
	fmt.Println("")
	fmt.Println("")
	fmt.Println("The test set is:")
	for i := 0; i < len(testSet); i++ {
		fmt.Println(testSet[i])
	}
	fmt.Println("There are", len(testSet), "elements in the list")
	fmt.Println("")
	fmt.Println("")
	testSet = removeDuplicates(testSet)
	fmt.Println("The test set after running removeDuplicates is:")
	for i := 0; i < len(testSet); i++ {
		fmt.Println(testSet[i])
	}
	fmt.Println("There are", len(testSet), "elements in the list")
}
