package explorer

import (
	"net/http"

	"github.com/amamov/kyle-coin/blockchain"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func HomeController(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func AddBlockController(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		blockData := r.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(blockData)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}

}
