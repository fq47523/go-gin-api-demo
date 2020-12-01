package utils

type Response struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
	Data interface{} `json:"data"`
}


func ResponseFmt(Code int,Message interface{},Data interface{},DataCount int)Response{
	DataFmt := make(map[string]interface{})
	DataFmt["items"] = Data
	DataFmt["total"] = DataCount

	return Response{Code,Message,DataFmt}
}





