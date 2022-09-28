package algorithms

import (
	"D7024E/node/contact"
	"D7024E/node/id"
	"fmt"
	"strconv"
	"testing"
)

// Validate successful merge batch.
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

// Verify that all contacts in batch gets a correct distance.
func TestGetAllDistances(t *testing.T) {
	contact.GetInstance().ID = id.NewKademliaID("TEST")
	id1 := *contact.GetInstance().ID
	id2 := id1
	id3 := id1
	id2[id.IDLength-1] += 1
	id3[id.IDLength-1] += 2
	batch := []contact.Contact{{ID: &id2}, {ID: &id3}}
	result := getAllDistances(batch)
	for i := 0; i < len(batch); i++ {
		num, err := strconv.Atoi(result[i].GetDistance().String())
		if err != nil {
			t.FailNow()
		} else if num != i+1 {
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
	fmt.Println(result)
	if len(result) != 2 {
		t.FailNow()
	}
}
