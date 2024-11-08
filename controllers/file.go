package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	middlewares "ipfs/middlewares"
	"ipfs/models"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
)

// Paste here the local path of your computer where the file will be downloaded
const YourLocalPath = "./"

// Paste here your public key
const YourPublicKey = "publickey"

var UploadFile = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		middlewares.ErrorResponse("error uploading file:"+err.Error(), rw)
		return
	}
	defer file.Close()

	src, err := header.Open()
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), rw)
		return
	}

	defer src.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), rw)
		return
	}

	sh := shell.NewShell("localhost:5001")
	cid, err := addFile(sh, data)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), rw)
		return
	}

	var resp models.FileResp

	resp.FileHash = cid
	resp.Filename = header.Filename
	resp.FileSize = header.Size
	resp.IpfsURL = fmt.Sprintf("https://ipfs.io/ipfs/%v", cid)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(resp)

	// middlewares.SuccessRespond(resp.IpfsURL, rw)
})

func addFile(sh *shell.Shell, data []byte) (string, error) {
	return sh.Add(bytes.NewReader(data))
}
