package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilites ServerCapabilities `json:"capabilities"`
	Serverinfo  Serverinfo         `json:"serverinfo"`
}

type ServerCapabilities struct{}

type Serverinfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilites: ServerCapabilities{},
			Serverinfo: Serverinfo{
				Name:    "dpt-lsp",
				Version: "0.0.0.0-beta",
			},
		},
	}
}
