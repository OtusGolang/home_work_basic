package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// если limit >= 0, обрезаем
	if limit >= 0 && limit < len(data) {
		data = data[:limit]
	}

	return data, nil
}

/*package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
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

	res := data

	return res, nil
}
*/
