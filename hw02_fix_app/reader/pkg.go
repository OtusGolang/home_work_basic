package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kidnation-pd/hw02_fix_app/types"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bt, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bt, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	if limit > 0 {
		return data[:limit], nil
	}

	return data, nil
}
