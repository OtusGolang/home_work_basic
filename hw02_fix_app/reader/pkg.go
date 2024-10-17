package reader

import (
	"encoding/json"
	"fmt"
	"hw02_fix_app/types"
	"io"
	"os"
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}
	res := data

	return res, nil
}
