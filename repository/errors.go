package repository

type RelationshipErrDoesNotBelong struct{}

func (e *RelationshipErrDoesNotBelong) Error() string {
	return "this relationship does not belong to this user"
}

type RelationshipErrNotAFriend struct{}

func (e *RelationshipErrNotAFriend) Error() string {
	return "these two users does not have relationship"
}
