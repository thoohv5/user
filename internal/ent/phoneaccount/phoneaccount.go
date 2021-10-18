// Code generated by entc, DO NOT EDIT.

package phoneaccount

const (
	// Label holds the string label denoting the phoneaccount user in the database.
	Label = "phone_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserIdentity holds the string denoting the user_identity field in the database.
	FieldUserIdentity = "user_identity"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// Table holds the table name of the phoneaccount in the database.
	Table = "phone_account"
)

// Columns holds all SQL columns for phoneaccount fields.
var Columns = []string{
	FieldID,
	FieldUserIdentity,
	FieldPhone,
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
