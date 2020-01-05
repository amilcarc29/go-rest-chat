package mocks

import (
	"crypto/md5"
	"encoding/json"
)

// Mock struct for Repository mock
type Mock struct {
	patchCheckMap map[hash][]outputForCheck
}

// NewMock Repository Mock
func NewMock() *Mock {
	patchCheckMap := make(map[hash][]outputForCheck)

	return &Mock{
		patchCheckMap: patchCheckMap,
	}
}

type hash [16]byte

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}

// PatchGet patch for Get function
func (mock *Mock) PatchCheck(outputCheck bool, outputErr error) {
	inputHash := toHash("check")
	output := getOutputForCheck(outputCheck, outputErr)

	if _, exists := mock.patchCheckMap[inputHash]; !exists {
		arrOutputForGet := make([]outputForCheck, 0)
		mock.patchCheckMap[inputHash] = arrOutputForGet
	}
	mock.patchCheckMap[inputHash] = append(mock.patchCheckMap[inputHash], output)
}

// Get mock for Get function
func (mock *Mock) Check() (bool, error) {
	inputHash := toHash("check")
	arrOutputForGet, exists := mock.patchCheckMap[inputHash]
	if !exists || len(arrOutputForGet) == 0 {
		panic("Mock not available for Check")
	}

	output := arrOutputForGet[0]
	arrOutputForGet = arrOutputForGet[1:]
	mock.patchCheckMap[inputHash] = arrOutputForGet

	return output.outputCheck, output.outputError
}

type outputForCheck struct {
	outputCheck bool
	outputError error
}

func getOutputForCheck(outputCheck bool, outputError error) outputForCheck {

	return outputForCheck{
		outputCheck: outputCheck,
		outputError: outputError,
	}
}
