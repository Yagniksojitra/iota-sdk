package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"github.com/iota-agency/iota-erp/internal/domain/user"
	"github.com/iota-agency/iota-erp/sdk/utils"
	"net/http"
	"time"

	model "github.com/iota-agency/iota-erp/graph/gqlmodels"
	"github.com/iota-agency/iota-erp/sdk/composables"
	"github.com/iota-agency/iota-erp/sdk/mapper"
)

// Authenticate is the resolver for the authenticate field.
func (r *mutationResolver) Authenticate(ctx context.Context, email string, password string) (*model.Session, error) {
	writer, ok := composables.UseWriter(ctx)
	if !ok {
		return nil, fmt.Errorf("request params not found")
	}
	_, session, err := r.app.AuthService.Authenticate(ctx, email, password)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    session.Token,
		Expires:  time.Now().Add(utils.SessionDuration()),
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
		Domain:   utils.GetEnv("DOMAIN", "localhost"),
	}
	http.SetCookie(writer, cookie)
	return session.ToGraph(), nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	u := &user.User{}
	if err := mapper.LenientMapping(&input, u); err != nil {
		return nil, err
	}
	if input.Password != nil {
		if err := u.SetPassword(*input.Password); err != nil {
			return nil, err
		}
	}
	if err := r.app.UserService.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return u.ToGraph(), nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, input model.UpdateUser) (*model.User, error) {
	user, err := r.app.UserService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := mapper.LenientMapping(&input, user); err != nil {
		return nil, err
	}
	if input.Password != nil {
		if err := user.SetPassword(*input.Password); err != nil {
			return nil, err
		}
	}
	if err := r.app.UserService.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	return user.ToGraph(), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	if err := r.app.UserService.DeleteUser(ctx, id); err != nil {
		return false, err
	}
	return true, nil
}

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input model.CreateRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: CreateRole - createRole"))
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id int64, input model.UpdateRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: UpdateRole - updateRole"))
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteRole - deleteRole"))
}

// CreateRolePermission is the resolver for the createRolePermission field.
func (r *mutationResolver) CreateRolePermission(ctx context.Context, input model.CreateRolePermission) (*model.RolePermissions, error) {
	panic(fmt.Errorf("not implemented: CreateRolePermission - createRolePermission"))
}

// CreateExpenseCategory is the resolver for the createExpenseCategory field.
func (r *mutationResolver) CreateExpenseCategory(ctx context.Context, input model.CreateExpenseCategory) (*model.ExpenseCategory, error) {
	panic(fmt.Errorf("not implemented: CreateExpenseCategory - createExpenseCategory"))
}

// UpdateExpenseCategory is the resolver for the updateExpenseCategory field.
func (r *mutationResolver) UpdateExpenseCategory(ctx context.Context, id int64, input model.UpdateExpenseCategory) (*model.ExpenseCategory, error) {
	panic(fmt.Errorf("not implemented: UpdateExpenseCategory - updateExpenseCategory"))
}

// DeleteExpenseCategory is the resolver for the deleteExpenseCategory field.
func (r *mutationResolver) DeleteExpenseCategory(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteExpenseCategory - deleteExpenseCategory"))
}

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.CreateExpense) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: CreateExpense - createExpense"))
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, id int64, input model.UpdateExpense) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: UpdateExpense - updateExpense"))
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteExpense - deleteExpense"))
}

// CreatePosition is the resolver for the createPosition field.
func (r *mutationResolver) CreatePosition(ctx context.Context, input model.CreatePosition) (*model.Position, error) {
	panic(fmt.Errorf("not implemented: CreatePosition - createPosition"))
}

// UpdatePosition is the resolver for the updatePosition field.
func (r *mutationResolver) UpdatePosition(ctx context.Context, id int64, input model.UpdatePosition) (*model.Position, error) {
	panic(fmt.Errorf("not implemented: UpdatePosition - updatePosition"))
}

// DeletePosition is the resolver for the deletePosition field.
func (r *mutationResolver) DeletePosition(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented: DeletePosition - deletePosition"))
}

// DeleteSession is the resolver for the deleteSession field.
func (r *mutationResolver) DeleteSession(ctx context.Context, token string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteSession - deleteSession"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	user, err := r.app.UserService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.ToGraph(), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedUsers, error) {
	users, err := r.app.UserService.GetUsersPaginated(ctx, limit, offset, sortBy)
	if err != nil {
		return nil, err
	}
	result := make([]*model.User, len(users))
	for _, user := range users {
		result = append(result, user.ToGraph())
	}
	total, err := r.app.UserService.GetUsersCount(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PaginatedUsers{
		Data:  result,
		Total: total,
	}, nil
}

// Upload is the resolver for the upload field.
func (r *queryResolver) Upload(ctx context.Context, id int64) (*model.Upload, error) {
	upload, err := r.app.UploadService.GetUploadByID(ctx, id)
	return upload.ToGraph(), err
}

// Uploads is the resolver for the uploads field.
func (r *queryResolver) Uploads(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedUploads, error) {
	uploads, err := r.app.UploadService.GetUploadsPaginated(ctx, limit, offset, sortBy)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Upload, len(uploads))
	for _, upload := range uploads {
		result = append(result, upload.ToGraph())
	}
	total, err := r.app.UploadService.GetUploadsCount(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PaginatedUploads{
		Data:  result,
		Total: total,
	}, nil
}

// Employee is the resolver for the employee field.
func (r *queryResolver) Employee(ctx context.Context, id int64) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented: Employee - employee"))
}

// Employees is the resolver for the employees field.
func (r *queryResolver) Employees(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedEmployees, error) {
	panic(fmt.Errorf("not implemented: Employees - employees"))
}

// Position is the resolver for the position field.
func (r *queryResolver) Position(ctx context.Context, id int64) (*model.Position, error) {
	panic(fmt.Errorf("not implemented: Position - position"))
}

// Positions is the resolver for the positions field.
func (r *queryResolver) Positions(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedPositions, error) {
	panic(fmt.Errorf("not implemented: Positions - positions"))
}

// Role is the resolver for the role field.
func (r *queryResolver) Role(ctx context.Context, id int64) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedRoles, error) {
	panic(fmt.Errorf("not implemented: Roles - roles"))
}

// Permission is the resolver for the permission field.
func (r *queryResolver) Permission(ctx context.Context, id int64) (*model.Permission, error) {
	panic(fmt.Errorf("not implemented: Permission - permission"))
}

// Permissions is the resolver for the permissions field.
func (r *queryResolver) Permissions(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedPermissions, error) {
	panic(fmt.Errorf("not implemented: Permissions - permissions"))
}

// RolePermission is the resolver for the rolePermission field.
func (r *queryResolver) RolePermission(ctx context.Context, roleID int64, permissionID int64) (*model.RolePermissions, error) {
	panic(fmt.Errorf("not implemented: RolePermission - rolePermission"))
}

// RolePermissions is the resolver for the rolePermissions field.
func (r *queryResolver) RolePermissions(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedRolePermissions, error) {
	panic(fmt.Errorf("not implemented: RolePermissions - rolePermissions"))
}

// ExpenseCategory is the resolver for the expenseCategory field.
func (r *queryResolver) ExpenseCategory(ctx context.Context, id int64) (*model.ExpenseCategory, error) {
	panic(fmt.Errorf("not implemented: ExpenseCategory - expenseCategory"))
}

// ExpenseCategories is the resolver for the expenseCategories field.
func (r *queryResolver) ExpenseCategories(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedExpenseCategories, error) {
	panic(fmt.Errorf("not implemented: ExpenseCategories - expenseCategories"))
}

// Expense is the resolver for the expense field.
func (r *queryResolver) Expense(ctx context.Context, id int64) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: Expense - expense"))
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedExpenses, error) {
	panic(fmt.Errorf("not implemented: Expenses - expenses"))
}

// AuthenticationLog is the resolver for the authenticationLog field.
func (r *queryResolver) AuthenticationLog(ctx context.Context, id int64) (*model.AuthenticationLog, error) {
	panic(fmt.Errorf("not implemented: AuthenticationLog - authenticationLog"))
}

// AuthenticationLogs is the resolver for the authenticationLogs field.
func (r *queryResolver) AuthenticationLogs(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedAuthenticationLogs, error) {
	panic(fmt.Errorf("not implemented: AuthenticationLogs - authenticationLogs"))
}

// Session is the resolver for the session field.
func (r *queryResolver) Session(ctx context.Context, token string) (*model.Session, error) {
	panic(fmt.Errorf("not implemented: Session - session"))
}

// Sessions is the resolver for the sessions field.
func (r *queryResolver) Sessions(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedSessions, error) {
	panic(fmt.Errorf("not implemented: Sessions - sessions"))
}

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan *model.User, error) {
	//ch := make(chan *model.Time)
	//
	//// You can (and probably should) handle your channels in a central place outside of `schema.resolvers.go`.
	//// For this example we'll simply use a Goroutine with a simple loop.
	//go func() {
	//	// Handle deregistration of the channel here. Note the `defer`
	//	defer close(ch)
	//
	//	for {
	//		// In our example we'll send the current time every second.
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("Tick")
	//
	//		// Prepare your object.
	//		currentTime := time.Now()
	//		t := &model.Time{
	//			UnixTime:  int(currentTime.Unix()),
	//			TimeStamp: currentTime.Format(time.RFC3339),
	//		}
	//
	//		// The subscription may have got closed due to the client disconnecting.
	//		// Hence we do send in a select block with a check for context cancellation.
	//		// This avoids goroutine getting blocked forever or panicking,
	//		select {
	//		case <-ctx.Done(): // This runs when context gets cancelled. Subscription closes.
	//			fmt.Println("Subscription Closed")
	//			// Handle deregistration of the channel here. `close(ch)`
	//			return // Remember to return to end the routine.
	//
	//		case ch <- t: // This is the actual send.
	//			// Our message went through, do nothing
	//		}
	//	}
	//}()
	//
	//// We return the channel and no error.
	//return ch, nil
	panic(fmt.Errorf("not implemented: UserCreated - userCreated"))
}

// UserUpdated is the resolver for the userUpdated field.
func (r *subscriptionResolver) UserUpdated(ctx context.Context) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented: UserUpdated - userUpdated"))
}

// UserDeleted is the resolver for the userDeleted field.
func (r *subscriptionResolver) UserDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: UserDeleted - userDeleted"))
}

// RoleCreated is the resolver for the roleCreated field.
func (r *subscriptionResolver) RoleCreated(ctx context.Context) (<-chan *model.Role, error) {
	panic(fmt.Errorf("not implemented: RoleCreated - roleCreated"))
}

// RoleUpdated is the resolver for the roleUpdated field.
func (r *subscriptionResolver) RoleUpdated(ctx context.Context) (<-chan *model.Role, error) {
	panic(fmt.Errorf("not implemented: RoleUpdated - roleUpdated"))
}

// RoleDeleted is the resolver for the roleDeleted field.
func (r *subscriptionResolver) RoleDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: RoleDeleted - roleDeleted"))
}

// RolePermissionCreated is the resolver for the rolePermissionCreated field.
func (r *subscriptionResolver) RolePermissionCreated(ctx context.Context) (<-chan *model.RolePermissions, error) {
	panic(fmt.Errorf("not implemented: RolePermissionCreated - rolePermissionCreated"))
}

// RolePermissionDeleted is the resolver for the rolePermissionDeleted field.
func (r *subscriptionResolver) RolePermissionDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: RolePermissionDeleted - rolePermissionDeleted"))
}

// ExpenseCategoryCreated is the resolver for the expenseCategoryCreated field.
func (r *subscriptionResolver) ExpenseCategoryCreated(ctx context.Context) (<-chan *model.ExpenseCategory, error) {
	panic(fmt.Errorf("not implemented: ExpenseCategoryCreated - expenseCategoryCreated"))
}

// ExpenseCategoryUpdated is the resolver for the expenseCategoryUpdated field.
func (r *subscriptionResolver) ExpenseCategoryUpdated(ctx context.Context) (<-chan *model.ExpenseCategory, error) {
	panic(fmt.Errorf("not implemented: ExpenseCategoryUpdated - expenseCategoryUpdated"))
}

// ExpenseCategoryDeleted is the resolver for the expenseCategoryDeleted field.
func (r *subscriptionResolver) ExpenseCategoryDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: ExpenseCategoryDeleted - expenseCategoryDeleted"))
}

// ExpenseCreated is the resolver for the expenseCreated field.
func (r *subscriptionResolver) ExpenseCreated(ctx context.Context) (<-chan *model.Expense, error) {
	panic(fmt.Errorf("not implemented: ExpenseCreated - expenseCreated"))
}

// ExpenseUpdated is the resolver for the expenseUpdated field.
func (r *subscriptionResolver) ExpenseUpdated(ctx context.Context) (<-chan *model.Expense, error) {
	panic(fmt.Errorf("not implemented: ExpenseUpdated - expenseUpdated"))
}

// ExpenseDeleted is the resolver for the expenseDeleted field.
func (r *subscriptionResolver) ExpenseDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: ExpenseDeleted - expenseDeleted"))
}

// PositionCreated is the resolver for the positionCreated field.
func (r *subscriptionResolver) PositionCreated(ctx context.Context) (<-chan *model.Position, error) {
	panic(fmt.Errorf("not implemented: PositionCreated - positionCreated"))
}

// PositionUpdated is the resolver for the positionUpdated field.
func (r *subscriptionResolver) PositionUpdated(ctx context.Context) (<-chan *model.Position, error) {
	panic(fmt.Errorf("not implemented: PositionUpdated - positionUpdated"))
}

// PositionDeleted is the resolver for the positionDeleted field.
func (r *subscriptionResolver) PositionDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: PositionDeleted - positionDeleted"))
}

// SessionDeleted is the resolver for the sessionDeleted field.
func (r *subscriptionResolver) SessionDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: SessionDeleted - sessionDeleted"))
}

// Avatar is the resolver for the avatar field.
func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*model.Upload, error) {
	if obj.AvatarID == nil {
		return nil, nil
	}
	upload, err := r.app.UploadService.GetUploadByID(ctx, *obj.AvatarID)
	return upload.ToGraph(), err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
