package weatherAPIClient

import "net/http"

type apiGetClient interface {
	call() (*http.Response, error)
}
