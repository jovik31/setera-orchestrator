package webhook

const (
	validatePodEndpoint    string = "/validate/pods"
	validateTenantEndpoint string = "/validate/tenants"
	tlsCertFile            string = "tls.crt"
	tlsCertKey             string = "tls.key"
	tlsDir                 string = "/run/secrets/tls"
	serverPort             string = ":8443"

	contentTypeHeader string = "content-type"
	contentTypeJSON   string = "application/json"
)
