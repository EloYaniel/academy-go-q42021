package apiclient

type ApiClient interface {
	Get(url string, params map[string]interface{}, response interface{}) error
}
