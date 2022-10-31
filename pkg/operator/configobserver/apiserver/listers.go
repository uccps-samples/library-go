package apiserver

import (
	configlistersv1 "github.com/uccps-samples/client-go/config/listers/config/v1"
)

type APIServerLister interface {
	APIServerLister() configlistersv1.APIServerLister
}
