package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amamov/kyle-coin/blockchain"
	"github.com/amamov/kyle-coin/utils"
	"github.com/gorilla/mux"
)

type url string

// interface, interceptor
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type apiDescription struct {
	URL         url
	Method      string `json:"method"` // json 포멧에서는 Method가 아니라 method로 보인다.
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty로 만일, Payload 값이 없다면 json 포멧에서 보여주지 않는다.
}

// func (ad apiDescription) String() string {
// 	return "Intercepted when printing an apiDescription object."
// }

type blockBodyForAppend struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func getDocsController(rw http.ResponseWriter, req *http.Request) {
	apiDocs := []apiDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Append A Block",
			Payload:     "message:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	// json.NewEncoder(rw).Encode(data) // 밑에 3줄을 한 줄로 쓴다면
	apiDocsJsonByte, err := json.Marshal(apiDocs) // Marshal : JSON으로 인코딩한 interface를 return한다.
	utils.HandleErr(err)
	fmt.Fprintf(rw, "%s", apiDocsJsonByte)
}

func appendBlockController(rw http.ResponseWriter, req *http.Request) {
	var blockBody blockBodyForAppend
	utils.HandleErr(json.NewDecoder(req.Body).Decode(&blockBody))
	blockchain.GetBlockChain().AppendBlock(blockBody.Message)
	rw.WriteHeader(http.StatusCreated) // 201
}

func getBlocksController(rw http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
}

func getBlockController(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	height, err := strconv.Atoi(vars["height"]) // str -> int 변환
	utils.HandleErr(err)
	block, err := blockchain.GetBlockChain().GetBlock(height)
	jsonEncoder := json.NewEncoder(rw)
	if err == blockchain.NotFoundError {
		jsonEncoder.Encode(errorResponse{ErrorMessage: fmt.Sprint(err)})
	} else {
		jsonEncoder.Encode(block)
	}
}
