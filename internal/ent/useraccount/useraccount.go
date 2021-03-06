// Code generated by entc, DO NOT EDIT.

package useraccount

const (
	// Label holds the string label denoting the useraccount user in the database.
	Label = "user_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserIdentity holds the string denoting the user_identity field in the database.
	FieldUserIdentity = "user_identity"
	// FieldAccount holds the string denoting the user field in the database.
	FieldAccount = "user"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// Table holds the table name of the useraccount in the database.
	Table = "user_account"
)

// Columns holds all SQL columns for useraccount fields.
var Columns = []string{
	FieldID,
	FieldUserIdentity,
	FieldAccount,
	FieldPassword,
	FieldSalt,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPassword holds the default value on creation for the password field.
	DefaultPassword string
	// DefaultSalt holds the default value on creation for the salt field.
	DefaultSalt string
)
