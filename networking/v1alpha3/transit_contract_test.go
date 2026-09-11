package v1alpha3

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestTransitServiceDescriptorAndJSONContract(t *testing.T) {
	service := &TransitService{}
	descriptor := service.ProtoReflect().Descriptor()
	if descriptor.FullName() != "dubbo.networking.v1alpha3.TransitService" || descriptor.ParentFile().Path() != "networking/v1alpha3/transit_service.proto" {
		t.Fatalf("incorrect generated descriptor: %s (%s)", descriptor.FullName(), descriptor.ParentFile().Path())
	}
	for name, input := range map[string]string{
		"LLM": `{"ai":{"provider":{"openai":{}},"modelRewrites":{"public":"upstream"}}}`,
		"MCP": `{"mcp":{"targets":[{"name":"tools","static":{"backendRef":{"name":"tools"},"port":8080,"path":"/mcp"}}]}}`,
		"A2A": `{"a2a":{"backendRef":{"name":"agent"},"port":8080,"agent":"assistant"}}`,
	} {
		t.Run(name, func(t *testing.T) {
			value := &TransitService{}
			if err := protojson.Unmarshal([]byte(input), value); err != nil {
				t.Fatal(err)
			}
			if value.GetService() == nil {
				t.Fatal("protocol variant was discarded")
			}
			if _, err := protojson.Marshal(value); err != nil {
				t.Fatal(err)
			}
		})
	}
}
