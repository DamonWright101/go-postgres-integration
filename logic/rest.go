package logic

import (
	"io/ioutil"
	"net/http"
)

func Query(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetRowFromTableByExactMatch(table string, col string, match string) (string, error) {

	url := "http://localhost:3000/" + table + "?" + col + "=eq." + match
	data, err := Query(url)
	if err != nil {
		return "", err
	}
	return data, nil
}

func GetRowsFromTableByBetweenTwoValues(table string, col string, lowermatch string, uppermatch string) (string, error) {

	url := "http://localhost:3000/" + table + "?" + col + "=gte." + lowermatch + "&" + col + "=lte." + uppermatch
	data, err := Query(url)
	if err != nil {
		return "", err
	}
	return data, nil
}
