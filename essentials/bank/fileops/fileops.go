package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WriteFloatToFile(value float64, fileName string) {
	text := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(text), 0644);
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName);

	if (err != nil) {
		return 1000, errors.New("failed to find file");
	}

	valueText := strings.Replace(string(data), "\n", "", -1);
	value, err := strconv.ParseFloat(valueText, 64);

	if (err != nil) {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil;
}
