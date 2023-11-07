package apierror

type ApiError struct {
	Code                 string            `json:"code,omitempty"`
	Description          string            `json:"description,omitempty"`
	OriginSystem         string            `json:"origin_system,omitempty"`
	Details              map[string]string `json:"details,omitempty"`
	StackTrace           any               `json:"stack_trace,omitempty"`
	EnableStackTraceBody bool              `json:"-"`
}

func NewError(statusCode int, enum GenericErrorCode, errorMessage string) ApiRuntimeError {
	getApiError := GetApiError(enum, errorMessage)
	return NewApiRuntimeError(statusCode, getApiError)
}

func GetApiError(errorCode GenericErrorCode, errorMessage string) ApiError {
	return Builder().
		Code(errorCode.GetCodeByEnv()).
		Description(errorCode.GetMessageByEnv()).
		OriginSystem(errorCode.GetOrigimSystemByEnv()).
		AddDetails("error_message", errorMessage).
		Build()
}

type apiErrorBuilder struct {
	apiError ApiError
}

func Builder() *apiErrorBuilder {
	return &apiErrorBuilder{
		apiError: ApiError{},
	}
}

func (a *apiErrorBuilder) Code(code string) *apiErrorBuilder {
	a.apiError.Code = code
	return a
}

func (a *apiErrorBuilder) Description(description string) *apiErrorBuilder {
	a.apiError.Description = description
	return a
}
func (a *apiErrorBuilder) OriginSystem(originSystem string) *apiErrorBuilder {
	a.apiError.OriginSystem = originSystem
	return a
}
func (a *apiErrorBuilder) Details(details map[string]string) *apiErrorBuilder {
	a.apiError.Details = details
	return a
}

func (a *apiErrorBuilder) AddDetails(key string, value string) *apiErrorBuilder {
	if a.apiError.Details == nil {
		a.apiError.Details = map[string]string{}
	}
	a.apiError.Details[key] = value
	return a
}

func (a *apiErrorBuilder) StackTrace(stackTrace any) *apiErrorBuilder {
	a.apiError.StackTrace = stackTrace
	return a
}
func (a *apiErrorBuilder) EnableStackTraceBody(enable bool) *apiErrorBuilder {
	a.apiError.EnableStackTraceBody = enable
	return a
}

func (a *apiErrorBuilder) Build() ApiError {
	if a.apiError.EnableStackTraceBody == false {
		a.StackTrace(nil)
	}
	return a.apiError
}
