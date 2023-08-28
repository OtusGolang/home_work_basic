package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/l1vesey/home_work_basic/tree/master/hw02_fix_app/types"
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(b, &data)

	res := data

	return res, nil
}
