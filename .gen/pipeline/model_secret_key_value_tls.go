/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.17.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type SecretKeyValueTls struct {
	Hosts string `json:"hosts"`
	Validity string `json:"validity,omitempty"`
	CaCert string `json:"caCert,omitempty"`
	CaKey string `json:"caKey,omitempty"`
	ServerCert string `json:"serverCert,omitempty"`
	ServerKey string `json:"serverKey,omitempty"`
	ClientCert string `json:"clientCert,omitempty"`
	ClientKey string `json:"clientKey,omitempty"`
}
