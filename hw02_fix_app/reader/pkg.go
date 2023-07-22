package reader

import	"encoding/json"
import	"fmt"
import	"io"
import	"os"

import	"github.com/fixme_my_friend/hw02_fix_app/types"


func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	byte, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	res := data

	return res, nil
}
