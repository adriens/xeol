package db

import xeolDB "github.com/noqcks/xeol/xeol/db/v1"

type mockStore struct {
	data map[string][]xeolDB.Cycle
}

func newMockStore() *mockStore {
	d := mockStore{
		data: make(map[string][]xeolDB.Cycle),
	}
	d.stub()
	return &d
}

func (d *mockStore) stub() {
	d.data["pkg:deb/debian/mongodb-org-server"] = []xeolDB.Cycle{
		{
			ProductName: "debian:distro:debian:8",
		},
	}
}

func (s *mockStore) GetCyclesByPurl(purl string) ([]xeolDB.Cycle, error) {
	return s.data[purl], nil
}

func (s *mockStore) GetAllProducts() (*[]xeolDB.Product, error) {
	return nil, nil
}
