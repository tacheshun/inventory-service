package receipt

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

var ReceiptDirectory string = filepath.Join("uploads")

type Receipt struct {
	ReceptName string `json:"name"`
	UploadDate time.Time `json:"uploadDate"`
}

func GetReceipts() ([]Receipt,error)  {
	receipts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptDirectory)
	if err != nil {
		return nil, err
	}
	for _, f := range files{
		receipts = append(receipts, Receipt{ReceptName: f.Name(), UploadDate: f.ModTime()})
	}
	return receipts, nil
}
