package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type CreatorObject struct {
	Name        string `json:"name"        toml:"name"`
	Type        string `json:"type"        toml:"type"`
	Description string `json:"description" toml:"description"`
}

type CreatorRole struct {
	Name        string `json:"name"        toml:"name"`
	Description string `json:"description" toml:"description"`
}

type CreatorPolicy struct {
	Object string   `json:"object" toml:"object"`
	Role   string   `json:"role"   toml:"role"`
	Action []string `json:"action" toml:"action"`
}

type DefaultPolicy struct {
	CreatorObject []*CreatorObject `toml:"creator_object"`
	CreatorRole   []*CreatorRole   `toml:"creator_role"`
	CreatorPolicy []*CreatorPolicy `toml:"creator_policy"`
}

func (f *DefaultPolicy) GetCreatorObject() ([]*CreatorObject, error) {
	return f.CreatorObject, nil
}

func (f *DefaultPolicy) GetCreatorRole() ([]*CreatorRole, error) {
	return f.CreatorRole, nil
}

func (f *DefaultPolicy) GetCreatorPolicy() ([]*CreatorPolicy, error) {
	return f.CreatorPolicy, nil
}

func NewDefaultPolicyByFile(policypath string) (*DefaultPolicy, error) {
	b, err := os.ReadFile(policypath)
	if err != nil {
		return nil, err
	}

	dictionary := &DefaultPolicy{}
	extension := filepath.Ext(policypath)
	switch extension {
	case ".toml":
		if err = toml.Unmarshal(b, dictionary); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("not supported extension %v", extension)
	}
	return dictionary, nil
}
