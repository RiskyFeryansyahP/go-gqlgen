// Code generated by entc, DO NOT EDIT.

package users

const (
	// Label holds the string label denoting the users type in the database.
	Label = "users"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username vertex property in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email vertex property in the database.
	FieldEmail = "email"

	// Table holds the table name of the users in the database.
	Table = "users"
)

// Columns holds all SQL columns for users fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Users type.
var ForeignKeys = []string{
	"meetups_user_id",
}