package weatherClients

import "net/http"

type GetClient interface {
	call() (*http.Response, error)
}
