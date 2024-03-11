package postgres

import (
	"database/sql"
	"fmt"
	pb "user-service/genproto/user-service"

	"github.com/golang/protobuf/ptypes/empty"
)

type userRepo struct {
	db *sql.DB
}

// NewUserRepo
func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	query := `INSERT INTO users(id,
								first_name, 
                  				last_name, 
                  				birth_date, 
                  				email, 
								password,
								refresh_token) 
								VALUES($1, $2, $3, $4, $5, $6, $7) 
								RETURNING id,
								first_name, 
                  				last_name, 
                  				birth_date, 
                  				email, 
                  				password,
								refresh_token,
								created_at`

	row := u.db.QueryRow(query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Email,
		user.Password,
		user.RefreshToken)
	if err := row.Scan(&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Email,
		&user.Password,
		&user.RefreshToken,
		&user.CreatedAt); err != nil {
			fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func (u *userRepo) UpdateUser(user *pb.User) (*pb.User, error) {
	query := `UPDATE users SET first_name = $1, last_name = $2, birth_date = $3 WHERE id = $4 AND deleted_at IS NULL
					RETURNING id,
					first_name, 
                  	last_name, 
					birth_date, 
					email, 
					password,
					created_at,
					`

	row := u.db.QueryRow(query, user.FirstName, user.LastName, user.BirthDate, user.Id)
	if err := row.Scan(&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Email,
		&user.Password,
		&user.CreatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepo) GetUserById(userId *pb.GetUserReqById) (*pb.User, error) {
	query := `SELECT id,
					first_name, 
                  	last_name, 
					birth_date, 
					email, 
					password,
					created_at FROM users WHERE id = $1 AND deleted_at IS NULL`
	row := u.db.QueryRow(query, userId.UserId)
	user := pb.User{}
	if err := row.Scan(&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Email,
		&user.Password,
		&user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) DeleteUser(userId *pb.DeleteUserReq) (*empty.Empty, error) {
	query := `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL`

	_, err := u.db.Exec(query, userId.UserId)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (u *userRepo) CheckField(req *pb.CheckFieldReq) (*pb.CheckFieldResp, error) {
	query := fmt.Sprintf(`SELECT count(1) FROM users WHERE %s = $1 AND deleted_at IS NULL`, req.Field)

	var isExists int

	row := u.db.QueryRow(query, req.Value)
	if err := row.Scan(&isExists); err != nil {
		return nil, err
	}

	if isExists == 1 {
		return &pb.CheckFieldResp{
			Status: true,
		}, nil
	}

	return &pb.CheckFieldResp{
		Status: false,
	}, nil
}

func (u *userRepo) GetAllUsers(req *pb.ListUsersReq) (*pb.ListUsersResp, error) {
	offset := (req.Page - 1) * req.Limit
	validColumns := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"birth_date": true,
	}
	orderby := "created_at"
	if _, ok := validColumns[req.Filter]; ok {
		orderby = req.Filter
	}

	query := fmt.Sprintf(`SELECT id,
				first_name,
				last_name,
				birth_date,
			    email,
 				password,
				created_at	
				FROM users WHERE deleted_at IS NULL ORDER BY %s LIMIT $1 OFFSET $2`, orderby)

	rows, err := u.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	var cnt int64
	users := &pb.ListUsersResp{}
	for rows.Next() {
		user := pb.User{}
		if err := rows.Scan(&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.BirthDate,
			&user.Email,
			&user.Password,
			&user.CreatedAt); err != nil {
			return nil, err
		}

		users.Users = append(users.Users, &user)
		cnt += 1
	}

	users.Count = cnt
	return users, nil
}

func (u *userRepo) IfExists(req *pb.IfExistsReq) (*pb.IfExistsResp, error) {
	query := `SELECT id,
					first_name, 
                  	last_name, 
					birth_date, 
					email, 
					password,
					created_at FROM users WHERE email = $1 AND deleted_at IS NULL`

	response := pb.IfExistsResp{User: &pb.User{}}
	row := u.db.QueryRow(query, req.Email)
	if err := row.Scan(&response.User.Id,
		&response.User.FirstName,
		&response.User.LastName,
		&response.User.BirthDate,
		&response.User.Email,
		&response.User.Password,
		&response.User.CreatedAt); err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *userRepo) ChangePassword(req *pb.ChangeUserPasswordReq) (*pb.ChangeUserPasswordResp, error) {
	query := `UPDATE users SET password = $1 WHERE email = $2 AND deleted_at IS NULL`

	_, err := u.db.Exec(query, req.Password, req.Email)
	if err != nil {
		return &pb.ChangeUserPasswordResp{Status: false}, err
	}

	return &pb.ChangeUserPasswordResp{Status: true}, nil
}

func (u *userRepo) UpdateRefreshToken(req *pb.UpdateRefreshTokenReq) (*pb.UpdateRefreshTokenResp, error) {
	query := `UPDATE users SET refresh_token = $1 WHERE id = $2 AND deleted_at IS NULL`

	_, err := u.db.Exec(query, req.RefreshToken, req.UserId)
	if err != nil {
		return &pb.UpdateRefreshTokenResp{Status: false}, err
	}

	return &pb.UpdateRefreshTokenResp{Status: true}, err
}
