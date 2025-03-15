package uplugin_test

import (
	"testing"

	"github.com/lvyonghuan/Ubik-Util/uconfig"
	"github.com/lvyonghuan/Ubik-Util/uplugin"
)

func TestReadModel(t *testing.T) {
	var plugin uplugin.Plugin
	err := uconfig.Read("./example-info.json", &plugin)
	if err != nil {
		t.Fatalf("Read function failed: %v", err)
	}

	// Verify basic plugin information
	if plugin.Name != "ExamplePlugin" {
		t.Errorf("Expected plugin name to be 'ExamplePlugin', got '%s'", plugin.Name)
	}
	if plugin.Version != "1.0.0" {
		t.Errorf("Expected plugin version to be '1.0.0', got '%s'", plugin.Version)
	}
	if plugin.Author != "Ubik" {
		t.Errorf("Expected plugin author to be 'Ubik', got '%s'", plugin.Author)
	}

	// Check nodes count
	if len(plugin.Nodes) != 3 {
		t.Errorf("Expected 3 nodes, got %d", len(plugin.Nodes))
	}

	// Verify StartNode properties
	startNode, exists := plugin.Nodes["StartNode"]
	if !exists {
		t.Fatal("Expected 'StartNode' to exist in nodes")
	}
	if !startNode.IsBegin {
		t.Error("Expected StartNode to be a begin node")
	}
	if len(startNode.Output) != 2 {
		t.Errorf("Expected StartNode to have 2 outputs, got %d", len(startNode.Output))
	}

	// Verify ProcessNode properties
	processNode, exists := plugin.Nodes["ProcessNode"]
	if !exists {
		t.Fatal("Expected 'ProcessNode' to exist in nodes")
	}
	if processNode.IsBegin {
		t.Error("Expected ProcessNode to not be a begin node")
	}

	// Check streaming property
	if !processNode.Output["processedData"].Streaming {
		t.Error("Expected ProcessNode's processedData output to be streaming")
	}

	// Verify OutputNode parameters
	outputNode, exists := plugin.Nodes["OutputNode"]
	if !exists {
		t.Fatal("Expected 'OutputNode' to exist in nodes")
	}
	if len(outputNode.Params) != 2 {
		t.Errorf("Expected OutputNode to have 2 params, got %d", len(outputNode.Params))
	}

	t.Log(plugin)
}

func TestReadModelByEncapsulatedFunc(t *testing.T) {
	plugin, err := uplugin.ReadPluginInfo("./example-info.json")
	if err != nil {
		t.Fatalf("Read function failed: %v", err)
	}

	// Verify basic plugin information
	if plugin.Name != "ExamplePlugin" {
		t.Errorf("Expected plugin name to be 'ExamplePlugin', got '%s'", plugin.Name)
	}
	if plugin.Version != "1.0.0" {
		t.Errorf("Expected plugin version to be '1.0.0', got '%s'", plugin.Version)
	}
	if plugin.Author != "Ubik" {
		t.Errorf("Expected plugin author to be 'Ubik', got '%s'", plugin.Author)
	}

	// Check nodes count
	if len(plugin.Nodes) != 3 {
		t.Errorf("Expected 3 nodes, got %d", len(plugin.Nodes))
	}

	// Verify StartNode properties
	startNode, exists := plugin.Nodes["StartNode"]
	if !exists {
		t.Fatal("Expected 'StartNode' to exist in nodes")
	}
	if !startNode.IsBegin {
		t.Error("Expected StartNode to be a begin node")
	}
	if len(startNode.Output) != 2 {
		t.Errorf("Expected StartNode to have 2 outputs, got %d", len(startNode.Output))
	}

	// Verify ProcessNode properties
	processNode, exists := plugin.Nodes["ProcessNode"]
	if !exists {
		t.Fatal("Expected 'ProcessNode' to exist in nodes")
	}
	if processNode.IsBegin {
		t.Error("Expected ProcessNode to not be a begin node")
	}

	// Check streaming property
	if !processNode.Output["processedData"].Streaming {
		t.Error("Expected ProcessNode's processedData output to be streaming")
	}

	// Verify OutputNode parameters
	outputNode, exists := plugin.Nodes["OutputNode"]
	if !exists {
		t.Fatal("Expected 'OutputNode' to exist in nodes")
	}
	if len(outputNode.Params) != 2 {
		t.Errorf("Expected OutputNode to have 2 params, got %d", len(outputNode.Params))
	}

	t.Log(plugin)
}
