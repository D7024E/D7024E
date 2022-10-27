package algorithms

import (
	rpc "D7024E/kademliaRPC/RPC"
	"D7024E/node/contact"
	"D7024E/node/id"
	"D7024E/node/stored"
	"testing"
	"time"
)

func RefreshMockupSuccess(id.KademliaID, contact.Contact, rpc.UDPSender) bool {
	return true
}

func RefreshMockupFail(id.KademliaID, contact.Contact, rpc.UDPSender) bool {
	return false
}

// Verify that the refresh does refresh the value.
func TestNodeRefreshRecSuccessfulRefresh(t *testing.T) {
	value := stored.Value{
		Data:   "SuccessfulRefresh",
		ID:     *id.NewKademliaID("SuccessfulRefresh"),
		TTL:    time.Second,
		DeadAt: time.Now().Add(time.Hour),
	}
	stored.GetInstance().AddRefresh(value.ID)
	res := NodeRefreshRec(value, nil, RefreshMockupSuccess)
	stored.GetInstance().StopRefresh(value.ID)
	if !res {
		t.Fatalf("value was not being refreshed after being added to refreshed")
	}
}

// Verify that the refresh will store value when refresh rpc fails.
func TestNodeRefreshRecFailedRefresh(t *testing.T) {
	value := stored.Value{
		Data:   "FailedRefresh",
		ID:     *id.NewKademliaID("FailedRefresh"),
		TTL:    time.Second,
		DeadAt: time.Now().Add(time.Hour),
	}
	stored.GetInstance().AddRefresh(value.ID)
	res := NodeRefreshRec(value, nil, RefreshMockupFail)
	stored.GetInstance().StopRefresh(value.ID)
	if !res {
		t.Fatalf("value was not being refreshed after being added to refreshed")
	}
}

// Test what happens when value is no longer being refreshed.
func TestNodeRefreshRecNotRefreshed(t *testing.T) {
	value := stored.Value{
		Data:   "NotRefreshed",
		ID:     *id.NewKademliaID("NotRefreshed"),
		TTL:    time.Second,
		DeadAt: time.Now().Add(time.Hour),
	}
	res := NodeRefreshRec(value, nil, RefreshMockupFail)
	if res {
		t.Fatalf("value not being refreshed found and refreshed")
	}
}
