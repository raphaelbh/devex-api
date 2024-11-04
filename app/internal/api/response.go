package api

type ApiResponse struct {
	Data any `json:"data"`
}

type EnvironmentResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func buildResponse(entity any) ApiResponse {
	return ApiResponse{
		Data: entity,
	}
}
