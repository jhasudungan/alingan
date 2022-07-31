package controller

import (
	"alingan/middleware"
	"alingan/model"
	"alingan/service"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type AgentManagamentController struct {
	AgentService   service.AgentService
	StoreService   service.StoreService
	AuthMiddleware middleware.AuthMiddleware
}

func (a *AgentManagamentController) ShowAgentData(w http.ResponseWriter, r *http.Request) {

	// isAuthenticated, err := a.AuthMiddleware.AuthenticateOwner(r)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, "Something Went Wrong - Exceute Render", 500)
	// 	return
	// }

	// if isAuthenticated == false {
	// 	http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
	// 	return
	// }

	// ownerId will get from session when authentication is integrated
	ownerId := "owner-001"

	agents, err := a.AgentService.GetOwnerAgentList(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/agent_list.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})

	templateData["agents"] = agents

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (a *AgentManagamentController) ShowAgentInformation(w http.ResponseWriter, r *http.Request) {

	// storeId
	params := mux.Vars(r)
	agentId := params["agentId"]

	agent, err := a.AgentService.GetAgentInformation(agentId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/view_agent.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})

	templateData["agent"] = agent

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (a *AgentManagamentController) ShowCreateAgentForm(w http.ResponseWriter, r *http.Request) {

	ownerId := "owner-001"

	stores, err := a.StoreService.FindStoreByOwnerId(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/register_agent.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})
	templateData["stores"] = stores

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (a *AgentManagamentController) HandleCreateAgentFormRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	request := model.RegisterNewAgentRequest{}
	request.StoreId = r.Form.Get("agent-store")
	request.AgentEmail = r.Form.Get("agent-email")
	request.AgentPassword = r.Form.Get("agent-password")
	request.AgentName = r.Form.Get("agent-name")

	err = a.AgentService.RegisterNewAgent(request)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	http.Redirect(w, r, "/owner/agent", http.StatusSeeOther)

}

func (a *AgentManagamentController) HandleSetAgentInactiveRequest(w http.ResponseWriter, r *http.Request) {

	// storeId
	params := mux.Vars(r)
	agentId := params["agentId"]

	err := a.AgentService.SetAgentInactive(agentId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	http.Redirect(w, r, "/owner/agent", http.StatusSeeOther)

}
