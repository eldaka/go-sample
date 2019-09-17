package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetAccountInfoResponse represent Get Account Info API Response
type GetAccountInfoResponse struct {
	ID                        string        `json:"id"`
	AvailableBalance          string        `json:"available_balance"`
	LedgerBalance             string        `json:"ledger_balance"`
	BankTransferRates         string        `json:"bank_transfer_rates"`
	BankTransferFees          string        `json:"bank_transfer_fees"`
	FullName                  string        `json:"full_name"`
	FirstName                 string        `json:"first_name"`
	LastName                  string        `json:"last_name"`
	DateOfBirth               string        `json:"date_of_birth"`
	Gender                    string        `json:"gender"`
	Email                     string        `json:"email"`
	UnconfirmedEmail          string        `json:"unconfirmed_email"`
	IsTrial                   bool          `json:"is_trial"`
	Country                   string        `json:"country"`
	State                     string        `json:"state"`
	City                      string        `json:"city"`
	Nationality               string        `json:"nationality"`
	AddressLine1              string        `json:"address_line_1"`
	AddressLine2              string        `json:"address_line_2"`
	PostalCode                string        `json:"postal_code"`
	IdentityNo                string        `json:"identity_no"`
	PhoneNo                   string        `json:"phone_no"`
	BankAccounts              []BankAccount `json:"bank_accounts"`
	AnnualIncome              string        `json:"annual_income"`
	IDFrontURL                string        `json:"id_front_url"`
	IDBackURL                 string        `json:"id_back_url"`
	Selfie2IDURL              string        `json:"selfie_2id_url"`
	ProofOfAddressURL         string        `json:"proof_of_address_url"`
	MultiBankAccountDetected  bool          `json:"multi_bank_account_detected"`
	AccountLocked             bool          `json:"account_locked"`
	KycLimitRemaining         float64       `json:"kyc_limit_remaining"`
	MetaData                  string        `json:"meta_data"`
	WalletName                string        `json:"wallet_name"`
	WalletID                  int           `json:"wallet_id"`
	GauthEnabled              bool          `json:"gauth_enabled"`
	KycVerified               bool          `json:"kyc_verified"`
	AccountFullyVerified      bool          `json:"account_fully_verified"`
	KycRejectedReason         interface{}   `json:"kyc_rejected_reason"`
	KycInformationEditAllowed bool          `json:"kyc_information_edit_allowed"`
	KycInformationVerifying   bool          `json:"kyc_information_verifying"`
	StorageLimitExceeded      bool          `json:"storage_limit_exceeded"`
	CountryOfBirth            string        `json:"country_of_birth"`
	NricIssueDate             string        `json:"nric_issue_date"`
	NricExpiryDate            string        `json:"nric_expiry_date"`
	NricType                  string        `json:"nric_type"`
	IsMyinfoFlow              bool          `json:"is_myinfo_flow"`
	Error                     string        `json:"error"`
}

// BankAccount represents single bank account detail
type BankAccount struct {
	ID                 int         `json:"id"`
	AccountNo          string      `json:"account_no"`
	AccountHolderName  string      `json:"account_holder_name"`
	BankAbbrev         string      `json:"bank_abbrev"`
	Disabled           string      `json:"disabled"`
	Hidden             string      `json:"hidden"`
	VerificationStatus string      `json:"verification_status"`
	Verified           bool        `json:"verified"`
	Usage              string      `json:"usage"`
	RejectReason       interface{} `json:"reject_reason"`
}

// GetAccountInfo will get account information from get account info API
func GetAccountInfo(ctx context.Context) (*GetAccountInfoResponse, error) {
	completeURL := fmt.Sprintf("%s%s", mainConfig.API.BaseAddress, mainConfig.API.GetAccountInfo)

	requestURL, err := url.Parse(completeURL)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo() request URL (%s) got : %v", completeURL, err)
	}

	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo() request URL (%s) got : %v", requestURL.String(), err)
	}

	//TODO: don't hard token
	req.Header.Set("X-XFERS-USER-API-KEY", "2zsujd47H3-UmsxDL784beVnYbxCYCzL4psSbwZ_Ngk")

	resp, err := clientReq.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo() request URL (%s) got : %v", requestURL.String(), err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GetAccountInfo() read body got :%v", err)
	}

	var result GetAccountInfoResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return &result, fmt.Errorf("GetAccountInfo() Unmarshall body got :%v", err)
	}

	if len(result.Error) > 0 {
		return &result, fmt.Errorf(result.Error)
	}

	return &result, nil
}

// GetTransferInfoResponse represents get transfer info API response
type GetTransferInfoResponse struct {
	UniqueID             string `json:"unique_id"`
	WalletName           string `json:"wallet_name"`
	BankNameFull         string `json:"bank_name_full"`
	BankAbbrev           string `json:"bank_abbrev"`
	BankNameAbbreviation string `json:"bank_name_abbreviation"`
	BankAccountNo        string `json:"bank_account_no"`
	BankPayeeName        string `json:"bank_payee_name"`
	BankCode             string `json:"bank_code"`
	BranchCode           string `json:"branch_code"`
	BranchArea           string `json:"branch_area"`
	ImgSrc               string `json:"img_src"`
	Error                string `json:"error"`
}

// GetTransferInfo will get transfer information from get transfer info API
func GetTransferInfo(ctx context.Context) (*GetTransferInfoResponse, error) {
	completeURL := fmt.Sprintf("%s%s", mainConfig.API.BaseAddress, mainConfig.API.GetTransferInfo)

	requestURL, err := url.Parse(completeURL)
	if err != nil {
		return nil, fmt.Errorf("GetTransferInfo() request URL (%s) got : %v", completeURL, err)
	}

	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("GetTransferInfo() request URL (%s) got : %v", requestURL.String(), err)
	}

	//TODO: don't hard token
	req.Header.Set("X-XFERS-USER-API-KEY", "2zsujd47H3-UmsxDL784beVnYbxCYCzL4psSbwZ_Ngk")

	resp, err := clientReq.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetTransferInfo() request URL (%s) got : %v", requestURL.String(), err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GetTransferInfo() read body got :%v", err)
	}

	var result GetTransferInfoResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return &result, fmt.Errorf("GetTransferInfo() Unmarshall body got :%v", err)
	}

	if len(result.Error) > 0 {
		return &result, fmt.Errorf(result.Error)
	}

	return &result, nil
}
