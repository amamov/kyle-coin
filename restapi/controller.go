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
func (u url) MarchalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url
	Method      string `json:"method"` // json 포멧에서는 Method가 아니라 method로 보인다.
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty로 만일, Payload 값이 없다면 json 포멧에서 보여주지 않는다.
}

// func (ud urlDescription) String() string {
// 	return ""
// }

type blockBodyForAdd struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
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
			Description: "Add A Block",
			Payload:     "message:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	// rw.Header().Add("Content-Type", "application/json") // 미들웨어로 뺌
	// json.NewEncoder(rw).Encode(data) // 밑에 3줄을 한 줄로 쓴다면

	b, err := json.Marshal(data) // Marshal : JSON으로 인코딩한 interface를 return한다.
	utils.HandleErr(err)
	// fmt.Println(b)
	// fmt.Printf("%s", b)
	fmt.Fprintf(rw, "%s", b)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
	case "POST":
		var blockBody blockBodyForAdd
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&blockBody))
		// fmt.Println(blockBody)
		blockchain.GetBlockChain().AddBlock(blockBody.Message)
		rw.WriteHeader(http.StatusCreated) // 201
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height, err := strconv.Atoi(vars["height"]) // str -> int 변환
	utils.HandleErr(err)
	block, err := blockchain.GetBlockChain().GetBlock(height)
	jsonEncoder := json.NewEncoder(rw)
	if err == blockchain.NotFoundError {
		jsonEncoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		jsonEncoder.Encode(block)
	}

}
