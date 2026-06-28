package client

//codec:gen
type GameRuleValues struct {
	Values []GameRuleValue
}

//codec:gen
type GameRuleValue struct {
	Key   string `mc:"Identifier"`
	Value string
}
