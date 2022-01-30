package block

type Block struct {
	Nr 				int 											`json:"nr" bson:"_id,omitempty"`
	Hash			string 										`json:"hash"`
	PrevHash  string 										`json:"prevHash"`
	Timestamp int64 										`json:"timestamp"`
	Data      map[string]interface{} 		`json:"data"`
	Nonce     int 											`json:"nonce"`
}