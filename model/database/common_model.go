package database

type DateColumn struct {
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
	DeletedAt *int64 `bson:"deleted_at"`
}
