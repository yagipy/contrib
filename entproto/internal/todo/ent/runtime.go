// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/contrib/entproto/internal/todo/ent/attachment"
	"entgo.io/contrib/entproto/internal/todo/ent/schema"
	"entgo.io/contrib/entproto/internal/todo/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attachmentFields := schema.Attachment{}.Fields()
	_ = attachmentFields
	// attachmentDescID is the schema descriptor for id field.
	attachmentDescID := attachmentFields[0].Descriptor()
	// attachment.DefaultID holds the default value on creation for the id field.
	attachment.DefaultID = attachmentDescID.Default.(func() uuid.UUID)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescBanned is the schema descriptor for banned field.
	userDescBanned := userFields[7].Descriptor()
	// user.DefaultBanned holds the default value on creation for the banned field.
	user.DefaultBanned = userDescBanned.Default.(bool)
}
