package webhook

import "net/http"

// [ ] Implement tenant validation webhook

func (ws *WebhookServer) HandleTenantValidation(w http.ResponseWriter, r *http.Request) {
	// [ ]: Check if tenant has the required labels
	// [ ]: If not, add default labels
	// [ ] If it has the required labels validate the tenant
}
