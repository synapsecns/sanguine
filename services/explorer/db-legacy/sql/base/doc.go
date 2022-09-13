// Package base provides the sql implementation of the database driver
//
// When contributing to this package, note: field names should not be defined as
// string literals in compiled code. This is error prone and likely to cause
// hard to debug issues. Ideally we could follow c# nameOf syntax
// (see: https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/operators/nameof)
// to do nameOf(BridgeDepositModel.BlockNumber) and get BlockNumber. from there
// we could figure out the gorm field name using ParseGormField(reflect.TypeOf(user).Elem().FieldByName("Name"))
// to get block_number. This is being proposed for go 2 (see: https://git.io/JW6QW for gorm specific and
// https://git.io/JW6Q0 or https://git.io/JW6Qz for more generalized proposals for solving this).
// In the meantime, we do the following:
//
// define the field name:
//
// var blockNumberField string
//
// then define an init() function for getting the field name, panicking if it cannot be found
//
//	func init(){
//		blockNumberField = getGormFieldName("BlockNumber")
//	}
//
// getGormFieldName will panic if the field is not available
package base
