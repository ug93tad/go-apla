// Exporting types for querying states
// ug93tad
package api

type EcoSystemsResult struct {
	Number uint32
}

type GUID struct {
	UID 	string
	Token 	string
	Expire	string
	EcosystemID	string
	KeyID	string
	Address	string
}
type LoginResult struct {
	Token       string        `json:"token,omitempty"`
	Refresh     string        `json:"refresh,omitempty"`
	EcosystemID string        `json:"ecosystem_id,omitempty"`
	KeyID       string        `json:"key_id,omitempty"`
	Address     string        `json:"address,omitempty"`
	NotifyKey   string        `json:"notify_key,omitempty"`
	IsNode      bool          `json:"isnode,omitempty"`
	IsOwner     bool          `json:"isowner,omitempty"`
	IsVDE       bool          `json:"vde,omitempty"`
	Timestamp   string        `json:"timestamp,omitempty"`
	Roles       []rolesResult `json:"roles,omitempty"`
}
type rolesResult struct {
	RoleId   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}
type ContractField struct {
	Name string `json:"name"`
	HTML string `json:"htmltype"`
	Type string `json:"txtype"`
	Tags string `json:"tags"`
}

type GetContractResult struct {
	StateID  uint32          `json:"state"`
	Active   bool            `json:"active"`
	TableID  string          `json:"tableid"`
	WalletID string          `json:"walletid"`
	TokenID  string          `json:"tokenid"`
	Address  string          `json:"address"`
	Fields   []ContractField `json:"fields"`
	Name     string          `json:"name"`
}

type ContractsResult struct {
	Count string              `json:"count"`
	List  []map[string]string `json:"list"`
}

type ColumnInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Perm string `json:"perm"`
}

type TableResult struct {
	Name       string       `json:"name"`
	Insert     string       `json:"insert"`
	NewColumn  string       `json:"new_column"`
	Update     string       `json:"update"`
	Read       string       `json:"read,omitempty"`
	Filter     string       `json:"filter,omitempty"`
	Conditions string       `json:"conditions"`
	AppID      string       `json:"app_id"`
	Columns    []ColumnInfo `json:"columns"`
}

type TableInfo struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}

type TablesResult struct {
	Count int64       `json:"count"`
	List  []TableInfo `json:"list"`
}

type ListResult struct {
	Count string              `json:"count"`
	List  []map[string]string `json:"list"`
}

