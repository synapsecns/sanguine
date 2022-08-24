package dbcommon

// GetGormFieldName wraps getGormFieldName to export it for testing.
func GetGormFieldName(model interface{}, fieldName string) string {
	return getGormFieldName(model, fieldName)
}
