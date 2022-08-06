package controller

import (
	"alingan/middleware"
	"alingan/model"
	"alingan/service"
	"html/template"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type AgentManagamentController struct {
	AgentService   service.AgentService
	StoreService   service.StoreService
	AuthMiddleware middleware.AuthMiddleware
	ErrorHandler   middleware.ErrorHandler
}

func (a *AgentManagamentController) ShowAgentData(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// ownerId will get from session when authentication is integrated
	ownerId := session.Id

	agents, err := a.AgentService.GetOwnerAgentList(ownerId)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/agent_list.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	templateData := make(map[string]interface{})

	templateData["agents"] = agents

	err = template.Execute(w, templateData)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

}

func (a *AgentManagamentController) ShowAgentInformation(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	agentId := params["agentId"]

	agent, err := a.AgentService.GetAgentInformation(agentId)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/view_agent.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	templateData := make(map[string]interface{})

	templateData["agent"] = agent

	err = template.Execute(w, templateData)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

}

func (a *AgentManagamentController) ShowCreateAgentForm(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	ownerId := session.Id

	stores, err := a.StoreService.FindStoreByOwnerId(ownerId)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/register_agent.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	templateData := make(map[string]interface{})
	templateData["stores"] = stores

	err = template.Execute(w, templateData)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

}

func (a *AgentManagamentController) HandleCreateAgentFormRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	request := model.RegisterNewAgentRequest{}
	request.StoreId = r.Form.Get("agent-store")
	request.AgentEmail = r.Form.Get("agent-email")
	request.AgentPassword = r.Form.Get("agent-password")
	request.AgentName = r.Form.Get("agent-name")

	err = a.AgentService.RegisterNewAgent(request)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	http.Redirect(w, r, "/owner/agent", http.StatusSeeOther)

}

func (a *AgentManagamentController) HandleSetAgentInactiveRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	agentId := params["agentId"]

	err = a.AgentService.SetAgentInactive(agentId)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	http.Redirect(w, r, "/owner/agent", http.StatusSeeOther)

}

func (a *AgentManagamentController) HandleReactiveActiveRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	agentId := params["agentId"]

	err = a.AgentService.SetAgentActive(agentId)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/agent")
		return
	}

	http.Redirect(w, r, "/owner/agent", http.StatusSeeOther)

}
