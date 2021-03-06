// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remoteach

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/moov-io/paygate/internal/model"
	"github.com/moov-io/paygate/internal/secrets"
	"github.com/moov-io/paygate/pkg/id"
)

func TestWEBPaymentType(t *testing.T) {
	var paymentType model.WEBPaymentType
	if err := json.Unmarshal([]byte(`"SINGLE"`), &paymentType); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(`"ReoCCuRRing"`), &paymentType); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(`"other"`), &paymentType); err == nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte("1"), &paymentType); err == nil {
		t.Fatal("expected error")
	}
}

func TestWEB__createWEBBatch(t *testing.T) {
	depID, userID := base.ID(), id.User(base.ID())
	keeper := secrets.TestStringKeeper(t)

	receiverDep := &model.Depository{
		ID:            id.Depository(base.ID()),
		BankName:      "foo bank",
		Holder:        "jane doe",
		HolderType:    model.Individual,
		Type:          model.Checking,
		RoutingNumber: "121042882",
		Status:        model.DepositoryVerified,
		Metadata:      "jane doe checking",
		Keeper:        keeper,
	}
	receiverDep.ReplaceAccountNumber("2")
	receiver := &model.Receiver{
		ID:                model.ReceiverID(base.ID()),
		Email:             "jane.doe@example.com",
		DefaultDepository: receiverDep.ID,
		Status:            model.ReceiverVerified,
		Metadata:          "jane doe",
	}
	origDep := &model.Depository{
		ID:            id.Depository(base.ID()),
		BankName:      "foo bank",
		Holder:        "john doe",
		HolderType:    model.Individual,
		Type:          model.Savings,
		RoutingNumber: "231380104",
		Status:        model.DepositoryVerified,
		Metadata:      "john doe savings",
		Keeper:        keeper,
	}
	origDep.ReplaceAccountNumber("2")
	orig := &model.Originator{
		ID:                model.OriginatorID(base.ID()),
		DefaultDepository: origDep.ID,
		Identification:    "dddd",
		Metadata:          "john doe",
	}
	amt, _ := model.NewAmount("USD", "100.00")
	transfer := &model.Transfer{
		ID:                     id.Transfer(base.ID()),
		Type:                   model.PushTransfer,
		Amount:                 *amt,
		Originator:             orig.ID,
		OriginatorDepository:   origDep.ID,
		Receiver:               receiver.ID,
		ReceiverDepository:     receiverDep.ID,
		Description:            "sending money",
		StandardEntryClassCode: "WEB",
		Status:                 model.TransferPending,
		UserID:                 userID.String(),
		WEBDetail: &model.WEBDetail{
			PaymentInformation: "test payment",
			PaymentType:        model.WEBSingle,
		},
	}

	batch, err := createWEBBatch(depID, transfer, receiver, receiverDep, orig, origDep)
	if err != nil {
		t.Fatal(err)
	}
	if batch == nil {
		t.Error("nil WEB Batch")
	}

	file, err := ConstructFile(depID, "", transfer, receiver, receiverDep, orig, origDep)
	if err != nil {
		t.Fatal(err)
	}
	if file == nil {
		t.Error("nil WEB ach.File")
	}

	// Make sure WEBReoccurring are rejected
	transfer.WEBDetail.PaymentType = "reoccurring"
	batch, err = createWEBBatch(depID, transfer, receiver, receiverDep, orig, origDep)
	if batch != nil || err == nil {
		t.Errorf("expected error, but got batch: %v", batch)
	} else {
		if !strings.Contains(err.Error(), "createWEBBatch: reoccurring WEB transfers are not supported") {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}
