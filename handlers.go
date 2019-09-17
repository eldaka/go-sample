package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/eldaka/go-sample/api"
)

type mockAData struct {
	BankNameFull     string
	BankAccountNo    string
	UniqueID         string
	AvailableBalance string
	LedgerBalance    string
}

func serveMockA(w http.ResponseWriter, r *http.Request) {
	//TODO: configure better pathing here
	t, err := template.ParseFiles("files/var/www/client/mockA.html")
	if err != nil {
		fmt.Fprintf(w, "Template not found")
	}

	account, err := api.GetAccountInfo(context.Background())
	if err != nil || account == nil {
		fmt.Fprintf(w, "Account not found")
	}

	transfer, err := api.GetTransferInfo(context.Background())
	if err != nil || transfer == nil {
		fmt.Fprintf(w, "Transfer not found")
	}

	data := mockAData{
		BankNameFull:     transfer.BankNameFull,
		BankAccountNo:    transfer.BankAccountNo,
		UniqueID:         transfer.UniqueID,
		AvailableBalance: account.AvailableBalance,
		LedgerBalance:    account.LedgerBalance,
	}

	t.Execute(w, data)
}
