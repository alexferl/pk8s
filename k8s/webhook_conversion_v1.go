// Code generated; DO NOT EDIT.

package k8s

// WebhookConversionV1 WebhookConversion describes how to call a conversion webhook
type WebhookConversionV1 struct {
	// WebhookClientConfig contains the information to make a TLS connection with the webhook.
	ClientConfig *WebhookClientConfigV1 `json:"clientConfig,omitempty"`
	// conversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail.
	ConversionReviewVersions []string `json:"conversionReviewVersions"`
}
