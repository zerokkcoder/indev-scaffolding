package repository

type UserRepository interface {
	// Create(ctx context.Context, user *model.User) error
	// Update(ctx context.Context, user *model.User) error
	// GetByID(ctx context.Context, id string) (*model.User, error)
	// GetByEmail(ctx context.Context, email string) (*model.User, error)
}

func NewUserRepository(
	r *Repository,
) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}
