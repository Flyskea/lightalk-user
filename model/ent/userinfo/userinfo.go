// Code generated by ent, DO NOT EDIT.

package userinfo

import (
	"time"
)

const (
	// Label holds the string label denoting the userinfo type in the database.
	Label = "user_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldMotto holds the string denoting the motto field in the database.
	FieldMotto = "motto"
	// Table holds the table name of the userinfo in the database.
	Table = "user_info"
)

// Columns holds all SQL columns for userinfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserID,
	FieldSex,
	FieldAge,
	FieldBirthday,
	FieldMotto,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultSex holds the default value on creation for the "sex" field.
	DefaultSex int8
)
