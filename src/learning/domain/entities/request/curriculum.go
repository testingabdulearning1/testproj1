package request

type SampleReq struct {
	SampleField1 string `json:"sample_field_1" validate:"required,min=3,max=100,alpha"`
	SampleField2 string `json:"sample_field_2" validate:"required,min=3,max=100,alphanum"`
}

type AddSyllabusReq struct {
	SyllabusName string `json:"syllabus_name" validate:"required,min=3,max=100"`
}
