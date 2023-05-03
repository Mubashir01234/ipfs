package models

type FileResp struct {
	FileHash string `json:"file_hash,omitempty" bson:"file_hash,omitempty"`
	Filename string `json:"filename,omitempty" bson:"filename,omitempty"`
	FileSize int64  `json:"file_size,omitempty" bson:"file_size,omitempty"`
	IpfsURL  string `json:"ipfs_url,omitempty" bson:"ipfs_url,omitempty"`
}
