// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/thoohv5/template/internal/ent/miniprogramaccount"
	"github.com/thoohv5/template/internal/ent/schema"
	"github.com/thoohv5/template/internal/ent/user"
	"github.com/thoohv5/template/internal/ent/useraccount"
	"github.com/thoohv5/template/internal/ent/userinfo"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	miniprogramaccountFields := schema.MiniProgramAccount{}.Fields()
	_ = miniprogramaccountFields
	// miniprogramaccountDescNickName is the schema descriptor for nick_name field.
	miniprogramaccountDescNickName := miniprogramaccountFields[3].Descriptor()
	// miniprogramaccount.DefaultNickName holds the default value on creation for the nick_name field.
	miniprogramaccount.DefaultNickName = miniprogramaccountDescNickName.Default.(string)
	// miniprogramaccountDescAvatarURL is the schema descriptor for avatar_url field.
	miniprogramaccountDescAvatarURL := miniprogramaccountFields[4].Descriptor()
	// miniprogramaccount.DefaultAvatarURL holds the default value on creation for the avatar_url field.
	miniprogramaccount.DefaultAvatarURL = miniprogramaccountDescAvatarURL.Default.(string)
	// miniprogramaccountDescGender is the schema descriptor for gender field.
	miniprogramaccountDescGender := miniprogramaccountFields[5].Descriptor()
	// miniprogramaccount.DefaultGender holds the default value on creation for the gender field.
	miniprogramaccount.DefaultGender = miniprogramaccountDescGender.Default.(int32)
	// miniprogramaccountDescCountry is the schema descriptor for country field.
	miniprogramaccountDescCountry := miniprogramaccountFields[6].Descriptor()
	// miniprogramaccount.DefaultCountry holds the default value on creation for the country field.
	miniprogramaccount.DefaultCountry = miniprogramaccountDescCountry.Default.(string)
	// miniprogramaccountDescProvince is the schema descriptor for province field.
	miniprogramaccountDescProvince := miniprogramaccountFields[7].Descriptor()
	// miniprogramaccount.DefaultProvince holds the default value on creation for the province field.
	miniprogramaccount.DefaultProvince = miniprogramaccountDescProvince.Default.(string)
	// miniprogramaccountDescCity is the schema descriptor for city field.
	miniprogramaccountDescCity := miniprogramaccountFields[8].Descriptor()
	// miniprogramaccount.DefaultCity holds the default value on creation for the city field.
	miniprogramaccount.DefaultCity = miniprogramaccountDescCity.Default.(string)
	// miniprogramaccountDescLanguage is the schema descriptor for language field.
	miniprogramaccountDescLanguage := miniprogramaccountFields[9].Descriptor()
	// miniprogramaccount.DefaultLanguage holds the default value on creation for the language field.
	miniprogramaccount.DefaultLanguage = miniprogramaccountDescLanguage.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescType is the schema descriptor for user field.
	userDescType := userFields[2].Descriptor()
	// user.DefaultType holds the default value on creation for the user field.
	user.DefaultType = userDescType.Default.(int32)
	// userDescIsDisable is the schema descriptor for is_disable field.
	userDescIsDisable := userFields[3].Descriptor()
	// user.DefaultIsDisable holds the default value on creation for the is_disable field.
	user.DefaultIsDisable = userDescIsDisable.Default.(int32)
	useraccountFields := schema.UserAccount{}.Fields()
	_ = useraccountFields
	// useraccountDescPassword is the schema descriptor for password field.
	useraccountDescPassword := useraccountFields[3].Descriptor()
	// useraccount.DefaultPassword holds the default value on creation for the password field.
	useraccount.DefaultPassword = useraccountDescPassword.Default.(string)
	// useraccountDescSalt is the schema descriptor for salt field.
	useraccountDescSalt := useraccountFields[4].Descriptor()
	// useraccount.DefaultSalt holds the default value on creation for the salt field.
	useraccount.DefaultSalt = useraccountDescSalt.Default.(string)
	userinfoFields := schema.UserInfo{}.Fields()
	_ = userinfoFields
	// userinfoDescChannel is the schema descriptor for channel field.
	userinfoDescChannel := userinfoFields[2].Descriptor()
	// userinfo.DefaultChannel holds the default value on creation for the channel field.
	userinfo.DefaultChannel = userinfoDescChannel.Default.(int32)
	// userinfoDescForm is the schema descriptor for form field.
	userinfoDescForm := userinfoFields[3].Descriptor()
	// userinfo.DefaultForm holds the default value on creation for the form field.
	userinfo.DefaultForm = userinfoDescForm.Default.(int32)
}
