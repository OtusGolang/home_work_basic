package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/aasdhajkshd/home_work_basic/tree/master/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading JSON content: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, nil
	}

	res := data
	return res, err
}
