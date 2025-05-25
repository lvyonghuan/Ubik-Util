package uplugin

import (
	"github.com/lvyonghuan/Ubik-Util/uconfig"
)

type Plugin struct {
	Name        string          `json:"name"`        //Name of the plugin
	Description string          `json:"description"` //Description of the plugin
	Version     string          `json:"version"`     //Version of the plugin
	Author      string          `json:"author"`      //Author of the plugin
	Addr        string          `json:"addr"`        //Address of the plugin, include port
	Local       bool            `json:"local"`       //if the plugin is local or remote
	Nodes       map[string]Node `json:"nodes"`       //Plugin node metadata. Key is the node name, value is the node metadata
}

type Node struct {
	Info    string `json:"info"`     //Description of the node
	IsBegin bool   `json:"is_begin"` //begin node will be the first batch of nodes to be executed in the task flow
	//In directed task flow, usually, there should not be two begin nodes in a loop,
	//because this will cause dependency disorder
	//However, considering the special situation,
	//if a node's partial output does not depend on the input of the predecessor node,
	//it can still serve as the beginning node in the loop
	//This node should be distinguished from the ordinary begin node, so this mark is used
	//The scenario is usually initialization, informing the successor nodes of initialization information,
	//such as how many devices are upstream, etc., to avoid manual configuration.
	//This mark should not be abused
	//Only effective when is_begin is true
	IsSpecialBegin bool            `json:"is_special_begin"`
	Input          map[string]Port `json:"input"`  //Input port metadata. Key is the port name, value is the port metadata
	Output         map[string]Port `json:"output"` //Output port metadata. Key is the port name, value is the port metadata
	Params         map[string]Port `json:"params"` //Parameter port metadata. Key is the port name, value is the port metadata
}

// Port represents input, output, or parameter port information
type Port struct {
	Attribute string `json:"attribute"` //If the properties are consistent, the two ports can be connected
	Type      string `json:"type"`      //Type of the port. Will be displayed in the front end
	Info      string `json:"info"`      //Description of the port

	//If the port is streaming
	//Only when the output port and input port of the two nodes have the same streaming properties can they be connected
	Streaming bool `json:"streaming"`
}

func ReadPluginInfo(path string) (*Plugin, error) {
	var plugin Plugin
	err := uconfig.Read(path, &plugin)
	if err != nil {
		return nil, err
	}

	return &plugin, nil
}

// Params User-set parameters, this structure is used to pass parameters from the leader to the follower
// Key is the parameter name, value is the parameter value. The value is a byte slice, which can be converted to the corresponding type
type Params map[string][]byte
