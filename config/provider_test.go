package config

import (
	"encoding/json"
	"testing"
)

func TestProviderResourcesConfigured(t *testing.T) {
	var s struct {
		ProviderSchemas map[string]struct {
			ResourceSchemas map[string]any `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal([]byte(providerSchema), &s); err != nil {
		t.Fatal(err)
	}
	tfResources := s.ProviderSchemas["registry.terraform.io/stackitcloud/stackit"].ResourceSchemas
	if len(tfResources) == 0 {
		t.Fatal("no resources found in schema.json")
	}
	for name := range tfResources {
		if _, ok := ExternalNameConfigs[name]; !ok {
			t.Errorf("resource %s from schema.json missing in ExternalNameConfigs", name)
		}
		if _, ok := resourceSpecs[name]; !ok {
			t.Errorf("resource %s from schema.json missing in resourceSpecs", name)
		}
	}

	for _, p := range []*struct {
		name  string
		count int
	}{
		{"cluster", len(GetProvider().Resources)},
		{"namespaced", len(GetProviderNamespaced().Resources)},
	} {
		if p.count != len(tfResources) {
			t.Fatalf("%s provider configured %d resources, want %d", p.name, p.count, len(tfResources))
		}
	}
	for name, r := range GetProvider().Resources {
		if r.ShortGroup == "" || r.Kind == "" {
			t.Fatalf("resource %s missing ShortGroup (%q) or Kind (%q)", name, r.ShortGroup, r.Kind)
		}
	}
}
