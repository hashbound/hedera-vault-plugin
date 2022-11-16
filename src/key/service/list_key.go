package service

import "fmt"

func (svc *KeyService) List() ([]string, error) {
	keys, err := svc.storage.List()
	if err != nil {
		return nil, fmt.Errorf("list keys failed: %s", err)
	}

	return []string(keys), nil
}
