package apierror

type ApiRuntimeError interface {
	GetError() *ApiError
	GetStatus() int
}

type apiRuntimeErrorImp struct {
	status   int
	apiError ApiError
}

func NewApiRuntimeError(status int, err ApiError) ApiRuntimeError {
	return &apiRuntimeErrorImp{
		status:   status,
		apiError: err,
	}
}

func (a *apiRuntimeErrorImp) GetError() *ApiError {
	return &a.apiError
}

func (a *apiRuntimeErrorImp) GetStatus() int {
	return a.status
}
