package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	Uf          string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
}

func main() {
	http.HandleFunc("/", BuscaCEP)    //Cria um endpoint.
	http.ListenAndServe(":8080", nil) //Sobe o servidor na porta desejada
}

// Request => Request body do Webservice
// Response => Resposta do endpoint
func BuscaCEP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, err := buscaCep(cepParam)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(cep)
	//Mesma coisa que:
	// jsonCep, err := json.Marshal(cep)
	// if err != nil {
	// 	writer.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// writer.Write(jsonCep)
}

func buscaCep(cepString string) (*ViaCEP, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cepString + "/json/")
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var Cep ViaCEP
	err = json.Unmarshal(body, &Cep)
	if err != nil {
		return nil, err
	}
	return &Cep, nil
}
