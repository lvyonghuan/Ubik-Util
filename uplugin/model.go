package uplugin

type Plugin struct {
	Name        string `json:"name"`        //Name of the plugin
	Description string `json:"description"` //Description of the plugin
	Version     string `json:"version"`     //Version of the plugin
	Author      string `json:"author"`      //Author of the plugin
	Nodes       []Node `json:"nodes"`       //Plugin node metadata
}

type Node struct {
	Name    string `json:"name"`     //Name of the node
	Info    string `json:"info"`     //Description of the node
	IsBegin bool   `json:"is_begin"` //begin node will be the first node to run
	Input   []Port `json:"input"`
	Output  []Port `json:"output"`
	Params  []Port `json:"params"`
}

// Port represents input, output, or parameter port information
type Port struct {
	Attribute string `json:"attribute"` //If the properties are consistent, the two ports can be connected
	Name      string `json:"name"`      //Name of the port. Will be displayed in the front end
	Type      string `json:"type"`      //Type of the port. Will be displayed in the front end
	Info      string `json:"info"`      //Description of the port

	//If the port is streaming
	//Only when the output port and input port of the two nodes have the same streaming properties can they be connected
	Streaming bool `json:"streaming"`
}
