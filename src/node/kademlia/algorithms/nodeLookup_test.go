package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/bucket"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/kademlia"
	"errors"
	"strconv"
	"testing"
)

// Mockup of ping RPC that always succeed.
func pingSuccess(contact.Contact, rpc.UDPSender) bool {
	return true
}

// Mockup of ping RPC that always fail.
func pingFail(contact.Contact, rpc.UDPSender) bool {
	return false
}

// Mockup of find node rpc that will succeed.
func findNodeSuccess(contact.Contact, id.KademliaID, rpc.UDPSender) ([]contact.Contact, error) {
	return []contact.Contact{
		{ID: id.NewKademliaID("127.21.0.2"), Address: "127.21.0.2"},
		{ID: id.NewKademliaID("127.21.0.3"), Address: "127.21.0.3"},
		{ID: id.NewKademliaID("127.21.0.4"), Address: "127.21.0.4"},
		{ID: id.NewKademliaID("127.21.0.5"), Address: "127.21.0.5"}}, nil
}

// Mockup of find node rpc that will fail.
func findNodeFail(contact.Contact, id.KademliaID, rpc.UDPSender) ([]contact.Contact, error) {
	return []contact.Contact{}, errors.New("not found")
}

// Validate if correct when no responses received.
func TestNodeLookupRecSuccess(t *testing.T) {
	kademlia.GetInstance()
	batch := []contact.Contact{{ID: id.NewKademliaID("a")}}
	result := NodeLookupRec(
		*id.NewKademliaID("127.21.0.2"),
		batch,
		findNodeSuccess,
		pingSuccess)
	if len(result) > bucket.BucketSize {
		t.FailNow()
	} else if len(result) != 5 {
		t.FailNow()
	}
}

// Validate if correct when no responses received.
func TestNodeLookupRecFail(t *testing.T) {
	kademlia.GetInstance()
	contactID := id.NewRandomKademliaID()
	result := NodeLookupRec(
		*contactID,
		[]contact.Contact{{ID: contactID}},
		findNodeFail,
		pingSuccess)
	if len(result) != 1 {
		t.FailNow()
	} else if !contactID.Equals(result[0].ID) {
		t.FailNow()
	}
}

// Verify that all contacts in batch gets a correct distance.
func TestGetAllDistances(t *testing.T) {
	contact.GetInstance().ID = id.NewKademliaID("TEST")
	id1 := *contact.GetInstance().ID
	id2 := id1
	id3 := id1
	id2[id.IDLength-1] += 1
	id3[id.IDLength-1] += 2
	batch := []contact.Contact{{ID: &id2}, {ID: &id3}}
	result := getAllDistances(*contact.GetInstance().ID, batch)
	for i := 0; i < len(batch); i++ {
		num, err := strconv.Atoi(result[i].GetDistance().String())
		if err != nil {
			t.FailNow()
		} else if num != i+1 {
			t.FailNow()
		}
	}
}

// Verify that min works as intended.
func TestMin1(t *testing.T) {
	if min(1, 2) != 1 {
		t.FailNow()
	}
}

// Verify that min works as intended.
func TestMin2(t *testing.T) {
	if min(2, 1) != 1 {
		t.FailNow()
	}
}

// Verify that all nodes are found.
func TestFindNodesSuccess(t *testing.T) {
	kademlia.GetInstance()
	batch := []contact.Contact{{ID: id.NewRandomKademliaID()}}
	result := findNodes(*id.NewRandomKademliaID(), batch, findNodeSuccess)
	if !batch[0].Equals(&result[0][0]) {
		t.FailNow()
	} else if len(result[1]) != 4 {
		t.FailNow()
	}
}

// Verify that all nodes are found.
func TestFindNodesFail(t *testing.T) {
	kademlia.GetInstance()
	batch := []contact.Contact{{ID: id.NewRandomKademliaID()}}
	result := findNodes(*id.NewRandomKademliaID(), batch, findNodeFail)
	if !batch[0].Equals(&result[0][0]) {
		t.FailNow()
	} else if len(result) != 1 {
		t.FailNow()
	}
}

// Validate the merge batch input with output.
func TestMergeBatch(t *testing.T) {
	contacts1 := []contact.Contact{
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.2"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.3"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.4"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.5"}}
	contacts2 := []contact.Contact{
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.4"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.5"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.6"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.7"}}
	batch := [][]contact.Contact{contacts1, contacts2}
	result := mergeBatch(batch)

	validator := []contact.Contact{
		contacts1[0],
		contacts1[1],
		contacts1[2],
		contacts1[3],
		contacts2[0],
		contacts2[1],
		contacts2[2],
		contacts2[3]}

	for i := 0; i < len(result); i++ {
		if !result[i].Equals(&validator[i]) {
			t.FailNow()
		}
	}
}

// Verify that duplicates are removed.
func TestRemoveDuplicates(t *testing.T) {
	batch := []contact.Contact{
		{ID: id.NewKademliaID("1")},
		{ID: id.NewKademliaID("2")},
		{ID: id.NewKademliaID("1")},
		{ID: id.NewKademliaID("2")}}
	result := removeDuplicates(batch)
	if len(result) != 2 {
		t.FailNow()
	}
}

// Verify that alive nodes are kept.
func TestRemoveDeadNodesAllAlive(t *testing.T) {
	batch := []contact.Contact{
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.2"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.3"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.4"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.5"}}
	result := removeDeadNodes(batch, pingSuccess)
	if len(result) != len(batch) {
		t.FailNow()
	}
}

// Verify that dead nodes are deleted.
func TestRemoveDeadNodesAllDead(t *testing.T) {
	batch := []contact.Contact{
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.2"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.3"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.4"},
		{ID: id.NewRandomKademliaID(), Address: "127.21.0.5"}}
	result := removeDeadNodes(batch, pingFail)
	if len(result) != 0 {
		t.FailNow()
	}
}

// Test if batch is resized correctly, when larger then bucketSize.
func TestResizeLarger(t *testing.T) {
	var batch []contact.Contact
	for i := 0; i < 2*bucket.BucketSize; i++ {
		batch = append(batch, contact.Contact{})
	}
	result := resize(batch)
	if len(result) > bucket.BucketSize {
		t.FailNow()
	}
}

// Test if batch is resized correctly, when smaller then bucketSize.
func TestResizeSmaller(t *testing.T) {
	var batch []contact.Contact
	for i := 0; i < 0.5*bucket.BucketSize; i++ {
		batch = append(batch, contact.Contact{})
	}
	result := resize(batch)
	if len(result) > bucket.BucketSize {
		t.FailNow()
	}
}

// Test if two batches are the same.
func TestIsSameTrue(t *testing.T) {
	batch := []contact.Contact{{ID: id.NewRandomKademliaID()}, {ID: id.NewRandomKademliaID()}}
	if !isSame(batch, batch) {
		t.FailNow()
	}
}

// Test if two batches are the same.
func TestIsSameFail(t *testing.T) {
	batch1 := []contact.Contact{{ID: id.NewKademliaID("a")}, {ID: id.NewKademliaID("b")}}
	batch2 := []contact.Contact{{ID: id.NewKademliaID("b")}, {ID: id.NewKademliaID("a")}}
	if isSame(batch1, batch2) {
		t.FailNow()
	}
}
