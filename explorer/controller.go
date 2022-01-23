package explorer

import (
	"net/http"

	"github.com/amamov/kyle-coin/blockchain"
	"github.com/amamov/kyle-coin/utils"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func HomeController(rw http.ResponseWriter, req *http.Request) {
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	utils.HandleErr(templates.ExecuteTemplate(rw, "home", data))
}

func BlockController(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		utils.HandleErr(templates.ExecuteTemplate(rw, "block", nil))
	case "POST":
		utils.HandleErr(req.ParseForm())
		blockData := req.Form.Get("blockData")
		blockchain.GetBlockChain().AppendBlock(blockData)
		http.Redirect(rw, req, "/", http.StatusPermanentRedirect)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
