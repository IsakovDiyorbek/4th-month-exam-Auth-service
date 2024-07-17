package postgres

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/Exam4/4th-month-exam-Auth-service/genproto/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}
func (p *UserRepo) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	query := `
		SELECT id, username, email, full_name, created_at
		FROM users
		WHERE id = $1
	`
	var user pb.GetProfileResponse
	err := p.db.QueryRowContext(ctx, query, req.Id).Scan(
		&user.Id, &user.Username, &user.Email, &user.FullName, &user.CreatedAt,
	)
	if err != nil {
		log.Printf("Error get profile: %v\n", err)
		return nil, err
	}
	return &user, nil
}


func (p *UserRepo) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, full_name = $3, updated_at = NOW()
		WHERE id = $4
		`

	_, err := p.db.ExecContext(ctx, query, req.Username, req.Email, req.FullName, req.Id)
	if err != nil {
		log.Printf("Error update profile: %v\n", err)
		return nil, err
	}
	return &pb.UpdateProfileResponse{Message: "Profile updated successfully"}, nil
	
}

func (p *UserRepo) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	query := `
		UPDATE users
		SET password_hash = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := p.db.ExecContext(ctx, query, req.NewPassword, req.Id)
	if err != nil {
		log.Printf("Error change password: %v\n", err)
		return nil, err
	}
	return &pb.ChangePasswordResponse{Message: "Password changed successfully"}, nil
}
	

