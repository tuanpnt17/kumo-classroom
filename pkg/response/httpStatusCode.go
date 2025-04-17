package response

const (
	Success      = 200001
	ParamInvalid = 20003
)

var message = map[int]string{
	Success:      "success",
	ParamInvalid: "Param is invalid",
}
