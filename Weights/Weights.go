package Weights

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Get_Weights() map[string]float64  {

	jsonFile, _ := os.Open("weights.json")
	bytes, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var f interface{}

	json.Unmarshal(bytes,&f)

	m := f.(map[string][]float64)
	weights := make(map[string]float64)
	for r , v := range m {
		weights[r]=v[0]

	}

	return weights
}
func Get_Threshold() map[string]float64  {

	jsonFile, _ := os.Open("weights.json")
	bytes, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var f interface{}

	json.Unmarshal(bytes,&f)

	m := f.(map[string][]float64)
	Threshold := make(map[string]float64)
	for r , v := range m {
		Threshold[r]=v[0]

	}

	return Threshold
}
