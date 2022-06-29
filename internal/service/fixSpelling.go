package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"nexign/internal/model"
	"strings"
)

const spellerApiUrl = "https://speller.yandex.net/services/spellservice.json/checkTexts"

func FixSpelling(givenText model.TextToCheck) (model.TextToCheck, error) {
	textToCheck := url.Values{}
	result := model.TextToCheck{}

	for _, value := range givenText.Texts {
		textToCheck.Add("text", value)
	}

	response, err := http.PostForm(spellerApiUrl, textToCheck)

	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return model.TextToCheck{}, err
	}

	var apiResp []model.SpellerApiResponse
	err = json.Unmarshal(rawBody, &apiResp)
	if err != nil {
		return model.TextToCheck{}, err
	}

	for index, value := range apiResp {
		for _, respValue := range value {
			givenText.Texts[index] = strings.Replace(givenText.Texts[index], respValue.Word, respValue.S[0], 1)
		}
		result.Texts = append(result.Texts, givenText.Texts[index])
	}
	return result, nil
}
