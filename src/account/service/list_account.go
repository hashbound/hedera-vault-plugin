package service

import "fmt"

func (svc *AccountService) List() ([]string, error) {
	accounts, err := svc.storage.List()
	if err != nil {
		return nil, fmt.Errorf("list accounts failed: %s", err)
	}

	return []string(accounts), nil
}
