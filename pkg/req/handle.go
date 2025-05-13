package req

import (
	"net/http"
	"url/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JsonResponse(*w, err.Error(), 402)
		return nil, err
	}

	err = IsValid[T](body)
	if err != nil {
		res.JsonResponse(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil

}
