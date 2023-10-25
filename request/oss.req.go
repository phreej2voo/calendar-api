package request

type OssAliyunPolicy struct {
	ModelName string `query:"modelName" validate:"required"`
	ModelId   string `query:"modelId" validate:"required"`
}
