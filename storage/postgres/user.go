package postgres

import (
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "user-service/genproto/user-service"
)

type userRepo struct {
	db *sql.DB
}

// NewUserRepo
func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

//rpc CreateUser(User) returns (User) {}
//rpc UpdateUser(User) returns (User) {}
//rpc GetUserById(GetUserReqById) returns (User) {}
//rpc GetUserByEmail(GetUserByEmailReq) returns (GetUserByEmailResp) {}
//rpc DeleteUser(DeleteUserReq) returns (google.protobuf.Empty) {}
//rpc CheckEmail(CheckEmailReq) returns (CheckEmailResp) {}
//rpc CheckField(CheckEmailReq) returns (CheckEmailResp) {}
//rpc GetAllUsers(ListUsersReq) returns (ListUsersResp) {}
//rpc Login(LoginReq) returns (LoginResp) {}

func (u *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	query := `INSERT INTO users(first_name, 
                  				last_name, 
                  				birth_date, 
                  				email, password) 
								VALUES($1, $2, $3, $4, $5) 
								RETURNING id,
								first_name, 
                  				last_name, 
                  				birth_date, 
                  				email, 
                  				password,
								created_at`

	row := u.db.QueryRow(query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Email,
		user.Password)
	if err := row.Scan(&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Email,
		&user.Password); err != nil {
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

func (u *userRepo) GetUserByEmail(userEmail *pb.GetUserByEmailReq) (*pb.GetUserByEmailResp, error) {
	query := `SELECT id,
					first_name, 
                  	last_name, 
					birth_date, 
					email, 
					password,
					created_at FROM users WHERE email = $1 AND deleted_at IS NULL`

	row := u.db.QueryRow(query, userEmail.Email)
	response := pb.GetUserByEmailResp{User: &pb.User{}}
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

	response.User = &user

	return &response, nil
}

func (u *userRepo) DeleteUser(userId *pb.DeleteUserReq) (*empty.Empty, error) {
	query := `UPDATE TABLE SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1`

	_, err := u.db.Exec(query, userId.UserId)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (u *userRepo) CheckEmail(emailReq *pb.CheckEmailReq) (*pb.CheckEmailResp, error) {
	query := `SELECT count(1) FROM users WHERE email = $1 AND deleted_at IS NOT NULL`

	var isExists int
	row := u.db.QueryRow(query, emailReq.Email)
	if err := row.Scan(&isExists); err != nil {
		return nil, err
	}

	if isExists == 1 {
		return &pb.CheckEmailResp{
			Status: true,
		}, nil
	}
	return &pb.CheckEmailResp{
		Status: false,
	}, nil
}

func (u *userRepo) CheckField(req *pb.CheckFieldReq) (*pb.CheckFieldResp, error) {
	query := fmt.Sprintf(`SELECT count(1) FROM users WHERE %s = $2 AND deleted_at IS NULL`, req.Field)

	var isExists int

	row := u.db.QueryRow(query, req.Value)
	if err := row.Scan(&isExists); err != nil {
		return nil, err
	}

	if isExists == 1 {
		return &pb.CheckFieldResp{
			Status: false,
		}, nil
	}

	return &pb.CheckFieldResp{
		Status: false,
	}, nil
}

func (u *userRepo) GetAllUsers(req *pb.ListUsersReq) (*pb.ListUsersResp, error) {
	query := `SELECT id
				first_name,
				last_name,
				birth_date,
			    email
 				password
				created_at 	
				FROM users WHERE deleted_at IS NULL`

	rows, err := u.db.Query(query)
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

func (u *userRepo) Login(req *pb.LoginReq) (*pb.LoginResp, error) {
	query := `SELECT id,
					first_name, 
                  	last_name, 
					birth_date, 
					email, 
					password,
					created_at FROM users WHERE email = $1 AND password = $2`

	response := pb.LoginResp{User: &pb.User{}}
	row := u.db.QueryRow(query, req.Email, req.Password)
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
