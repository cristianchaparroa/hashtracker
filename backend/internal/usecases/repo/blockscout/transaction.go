package blockscout

import (
	"context"
	"encoding/json"
	"fmt"
	"hashtracker/internal/entities"
	"hashtracker/internal/usecases"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://eth.blockscout.com/api/v2"
)

type transaction struct {
}

func NewTransactionRepository() usecases.ETHTransactionRepository {
	return &transaction{}
}

func (t *transaction) GetTransactions(ctx context.Context, address string) (*entities.TransactionList, error) {
	endpoint := fmt.Sprintf("/addresses/%s/transactions", address)
	params := url.Values{}
	params.Add("sort", "desc")
	// params.Add("filter", "to") // Change to "from" for outgoing transactions
	params.Add("limit", "100") // Adjust this value as needed
	fullURL := fmt.Sprintf("%s%s?%s", baseURL, endpoint, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println(string(body))
	var result TransactionResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	var txs []*entities.Transaction
	for _, r := range result.Items {
		tx := &entities.Transaction{
			Hash:      r.Hash,
			Value:     r.Value,
			From:      r.From.Hash,
			To:        r.To.Hash,
			Timestamp: r.Timestamp,
		}
		fmt.Println(tx)
		txs = append(txs, tx)
	}
	return entities.NewTransactionList(txs), nil
}

func (t *transaction) GetOutTransactions(ctx context.Context, address string) (*entities.TransactionList, error) {
	txs, err := t.GetTransactions(ctx, address)
	if err != nil {
		return nil, err
	}

	results := make([]*entities.Transaction, 0)
	for _, tx := range txs.List {
		if tx.To == address {
			continue
		}
		results = append(results, tx)
	}
	return entities.NewTransactionList(results), nil
}