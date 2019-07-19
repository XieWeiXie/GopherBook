package vote

import (
	"GopherBook/chapter11/web/make_request"
	"log"
	"net/http"
)

type ControllerVote struct {
}

var Default = ControllerVote{}

func (c ControllerVote) GetAllVote(writer http.ResponseWriter, request *http.Request) {
	search := make_request.Query(request, "search")
	log.Println("Search: ", search)

	returnAll := make_request.QueryAndDefault(request, "return", "all_list")

	log.Println("Return: ", returnAll)

}
func (c ControllerVote) GetOneVote(writer http.ResponseWriter, request *http.Request) {
	voteId := make_request.Query(request, "vote_id")
	log.Println("vote id: ", voteId)
}
func (c ControllerVote) PostOneVote(writer http.ResponseWriter, request *http.Request) {
	var param CreateVoteParam
	if err := make_request.BindJson(request, &param); err != nil {
		log.Println("err: ", err.Error())
		return
	}
	log.Println("Param: ", param)
}
func (c ControllerVote) PatchOneVote(writer http.ResponseWriter, request *http.Request) {
	params, err := make_request.Params(request)
	if err != nil {
		log.Println("err: ", err.Error())
		return
	}
	log.Println("Param: ", params)

}
func (c ControllerVote) DeleteOneVote(writer http.ResponseWriter, request *http.Request) {}

func (c ControllerVote) VoteHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPatch {
		c.PatchOneVote(writer, request)
		return
	}
	if request.Method == http.MethodGet {
		c.GetOneVote(writer, request)
		return
	}
	if request.Method == http.MethodDelete {
		c.DeleteOneVote(writer, request)
		return
	}
	if request.Method == http.MethodPost {
		c.PostOneVote(writer, request)
		return
	}
}
