// Code generated; DO NOT EDIT.

package k8s

// CustomResourceConversionV1 CustomResourceConversion describes how to convert different versions of a CR.
type CustomResourceConversionV1 struct {
	// strategy specifies how custom resources are converted between versions. Allowed values are: - `"None"`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `"Webhook"`: API Server will call to an external webhook to do the conversion. Additional information   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
	Strategy string `json:"strategy"`
	// WebhookConversion describes how to call a conversion webhook
	Webhook *WebhookConversionV1 `json:"webhook,omitempty"`
}
