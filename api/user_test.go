package api

import (
	"context"
	"net/http"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func Test_GetAccountInfo(t *testing.T) {
	type args struct {
		mockMethod      string
		mockURL         string
		mockBaseAddress string
		mockEndpoint    string
		mockStatus      int
		payload         string
		ctx             context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *GetAccountInfoResponse
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx:             context.Background(),
				mockMethod:      "GET",
				mockURL:         "https://sandbox.xfers.io/api/v3/user",
				mockBaseAddress: "https://sandbox.xfers.io",
				mockEndpoint:    "/api/v3/user",
				mockStatus:      200,
				payload:         `{"id":"user_osdzwma400iq","available_balance":"93739.21","ledger_balance":"93739.21","bank_transfer_rates":"0.0","bank_transfer_fees":"0.45","full_name":"","first_name":"docs.xfers.io","last_name":"testingaccount","date_of_birth":"","gender":"male","email":"docs@xfers.com","unconfirmed_email":"","is_trial":true,"country":"SG","state":"Singapore","city":"Singapore","nationality":"Singaporean","address_line_1":"Blk 71 Ayer Rajah Cresent","address_line_2":"#02-52","postal_code":"541121","identity_no":"S8117102G","phone_no":"+6589564339","bank_accounts":[{"id":400,"account_no":"0393123432","account_holder_name":"Tian Wei","bank_abbrev":"DBS","disabled":"false","hidden":"false","verification_status":"pending","verified":true,"usage":"all","reject_reason":null},{"id":808,"account_no":"0393123432","account_holder_name":null,"bank_abbrev":"CITI","disabled":"false","hidden":"false","verification_status":"pending","verified":true,"usage":"all","reject_reason":null}],"annual_income":"","id_front_url":"","id_back_url":"","selfie_2id_url":"","proof_of_address_url":"","multi_bank_account_detected":false,"account_locked":false,"kyc_limit_remaining":200000.0,"meta_data":"","wallet_name":"Your General Wallet Account","wallet_id":1,"gauth_enabled":false,"kyc_verified":true,"account_fully_verified":true,"kyc_rejected_reason":null,"kyc_information_edit_allowed":false,"kyc_information_verifying":false,"storage_limit_exceeded":false,"country_of_birth":"SG","nric_issue_date":"","nric_expiry_date":"","nric_type":"","is_myinfo_flow":false}`,
			},
			want:    &GetAccountInfoResponse{}, // TODO: check marshall correct here
			wantErr: false,
		},
		{
			name: "test error invalid request",
			args: args{
				ctx:             context.Background(),
				mockMethod:      "GET",
				mockURL:         "https://sandbox.xfers.io/api/v3/user",
				mockBaseAddress: "https://sandbox.xfers.io",
				mockEndpoint:    "/api/v3/user",
				mockStatus:      200,
				payload:         `{"error":"Invalid API KEY"}`,
			},
			want:    &GetAccountInfoResponse{Error: "Invalid API KEY"},
			wantErr: true,
		},

		//TODO: add other test negative cases
	}
	for _, tt := range tests {
		if tt.args.mockMethod != "" && tt.args.mockBaseAddress != "" {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder(tt.args.mockMethod, tt.args.mockURL, httpmock.NewStringResponder(tt.args.mockStatus, tt.args.payload))
			clientReq = &http.Client{}
			mainConfig.API.BaseAddress = tt.args.mockBaseAddress
			mainConfig.API.GetAccountInfo = tt.args.mockEndpoint
		}

		t.Run(tt.name, func(t *testing.T) {
			//TODO: need to check returned response here
			_, err := GetAccountInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountInfo(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetTransferInfo(t *testing.T) {
	type args struct {
		mockMethod      string
		mockURL         string
		mockBaseAddress string
		mockEndpoint    string
		mockStatus      int
		payload         string
		ctx             context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *GetTransferInfoResponse
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx:             context.Background(),
				mockMethod:      "GET",
				mockURL:         "https://sandbox.xfers.io/api/v3/user/transfer_info",
				mockBaseAddress: "https://sandbox.xfers.io",
				mockEndpoint:    "/api/v3/user/transfer_info",
				mockStatus:      200,
				payload:         `{"id":"user_osdzwma400iq","available_balance":"93739.21","ledger_balance":"93739.21","bank_transfer_rates":"0.0","bank_transfer_fees":"0.45","full_name":"","first_name":"docs.xfers.io","last_name":"testingaccount","date_of_birth":"","gender":"male","email":"docs@xfers.com","unconfirmed_email":"","is_trial":true,"country":"SG","state":"Singapore","city":"Singapore","nationality":"Singaporean","address_line_1":"Blk 71 Ayer Rajah Cresent","address_line_2":"#02-52","postal_code":"541121","identity_no":"S8117102G","phone_no":"+6589564339","bank_accounts":[{"id":400,"account_no":"0393123432","account_holder_name":"Tian Wei","bank_abbrev":"DBS","disabled":"false","hidden":"false","verification_status":"pending","verified":true,"usage":"all","reject_reason":null},{"id":808,"account_no":"0393123432","account_holder_name":null,"bank_abbrev":"CITI","disabled":"false","hidden":"false","verification_status":"pending","verified":true,"usage":"all","reject_reason":null}],"annual_income":"","id_front_url":"","id_back_url":"","selfie_2id_url":"","proof_of_address_url":"","multi_bank_account_detected":false,"account_locked":false,"kyc_limit_remaining":200000.0,"meta_data":"","wallet_name":"Your General Wallet Account","wallet_id":1,"gauth_enabled":false,"kyc_verified":true,"account_fully_verified":true,"kyc_rejected_reason":null,"kyc_information_edit_allowed":false,"kyc_information_verifying":false,"storage_limit_exceeded":false,"country_of_birth":"SG","nric_issue_date":"","nric_expiry_date":"","nric_type":"","is_myinfo_flow":false}`,
			},
			want:    &GetTransferInfoResponse{}, // TODO: check marshall correct here
			wantErr: false,
		},
		{
			name: "test error invalid request",
			args: args{
				ctx:             context.Background(),
				mockMethod:      "GET",
				mockURL:         "https://sandbox.xfers.io/api/v3/user/transfer_info",
				mockBaseAddress: "https://sandbox.xfers.io",
				mockEndpoint:    "/api/v3/user/transfer_info",
				mockStatus:      200,
				payload:         `{"error":"Invalid API KEY"}`,
			},
			want:    &GetTransferInfoResponse{Error: "Invalid API KEY"},
			wantErr: true,
		},

		//TODO: add other test negative cases
	}
	for _, tt := range tests {
		if tt.args.mockMethod != "" && tt.args.mockBaseAddress != "" {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder(tt.args.mockMethod, tt.args.mockURL, httpmock.NewStringResponder(tt.args.mockStatus, tt.args.payload))
			clientReq = &http.Client{}
			mainConfig.API.BaseAddress = tt.args.mockBaseAddress
			mainConfig.API.GetTransferInfo = tt.args.mockEndpoint
		}

		t.Run(tt.name, func(t *testing.T) {
			//TODO: need to check returned response here
			_, err := GetTransferInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransferInfo(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}
