package common

import (
	"testing"
	"time"
)

var cel cell = cell{
	level: 1_000,
}

var transferLimit uint = 100
var dst chan uint = make(chan uint)
var cdt *conduit = NewConduit(transferLimit, &cel, dst)

func RunTransfer(
	t *testing.T,
	availableEnergy uint,
	cdt Conduit,
	amt uint,
	expectedRec *uint,
) {
	cel.level = availableEnergy

	go cdt.Transfer(amt)

	select {
	case rec := <-dst:
		if rec != *expectedRec {
			t.Errorf("transfer failed, expected to receive %d, received: %d", amt, rec)
		}

		if cel.level != availableEnergy-*expectedRec {
			t.Errorf(
				"transfer failed, expected to drain %d, drained: %d",
				expectedRec,
				availableEnergy-cel.level,
			)
		}
	case <-time.After(1 * time.Second):
		if expectedRec != nil {
			t.Error("transfer failed, energy was not received")
		}
	}
}

func TestConduit_Transfer_WithinLimit(t *testing.T) {
	var expected uint = 3
	RunTransfer(t, 1_000, cdt, 3, &expected)
}

func TestConduit_Transfer_AtLimit(t *testing.T) {
	RunTransfer(t, 1_000, cdt, transferLimit, &transferLimit)
}

func TestConduit_Transfer_ExceedLimit(t *testing.T) {
	RunTransfer(t, 1_000, cdt, transferLimit+1, &transferLimit)
}

func TestConduit_Transfer_AtAvailable(t *testing.T) {
	var expected uint = 50
	RunTransfer(t, 50, cdt, 50, &expected)
}

func TestConduit_Transfer_ExceedAvailable(t *testing.T) {
	var expected uint = 50
	RunTransfer(t, 50, cdt, 51, &expected)
}

func TestConduit_Transfer_ExceedLimitAndAvailable(t *testing.T) {
	var expected uint = 50
	RunTransfer(t, 50, cdt, transferLimit+1, &expected)
}

func TestConduit_Transfer_ZeroEnergy(t *testing.T) {
	RunTransfer(t, 1_000, cdt, 0, new(uint))
}
