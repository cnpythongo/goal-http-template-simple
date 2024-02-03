package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/admin/types"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"strings"
)

type IAccountUserService interface {
	GetUserList(req *types.ReqGetUserList) (*types.RespGetUserList, int, error)
	GetUserDetail(uuid string) (*types.RespUserDetail, int, error)
	CreateUser(payload *types.ReqCreateUser) (*types.RespUserDetail, int, error)
	DeleteUserByUUID(uuid string) (int, error)
	UpdateUserByUUID(uuid string, payload *types.ReqUpdateUser) (int, error)

	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUserLastLogin(uuid string) error
}

type accountUserService struct {
	db *gorm.DB
}

func (s *accountUserService) GetUserList(req *types.ReqGetUserList) (*types.RespGetUserList, int, error) {
	var query []string
	var args []interface{}
	if req.Email != "" {
		query = append(query, "email like concat('%', ?, '%s')")
		args = append(args, req.Email)
	}
	if req.Phone != "" {
		query = append(query, "phone like concat('%', ?, '%s')")
		args = append(args, req.Phone)
	}
	if req.Nickname != "" {
		query = append(query, "nickname like concat('%s', ?, '%s')")
		args = append(args, req.Nickname)
	}
	if len(req.Status) > 0 {
		status := strings.Split(string(req.Status[0]), ",")
		query = append(query, "status IN ?")
		args = append(args, status)
	}
	if req.LastLoginAtStart != "" {
		query = append(query, "last_login_at >= ?")
		args = append(args, req.LastLoginAtStart)
	}
	if req.LastLoginAtEnd != "" {
		query = append(query, "last_login_at <= ?")
		args = append(args, req.LastLoginAtEnd)
	}
	queryStr := strings.Join(query, " AND ")
	rows, total, err := model.GetUserList(s.db, req.Page, req.Limit, queryStr, args)
	if err != nil {
		return nil, render.QueryError, err
	}

	result := make([]*types.RespUserBasic, 0)
	for _, row := range rows {
		item := new(types.RespUserBasic)
		err = copier.Copy(item, row)
		if err != nil {
			log.GetLogger().Error(err)
			return nil, render.DBAttributesCopyError, err
		}
		result = append(result, item)
	}
	resp := &types.RespGetUserList{
		Page:   req.Page,
		Limit:  req.Limit,
		Total:  total,
		Result: result,
	}
	return resp, render.OK, nil
}

func (s *accountUserService) GetUserDetail(uuid string) (*types.RespUserDetail, int, error) {
	user, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return nil, render.QueryError, err
		}
	}
	return s.transUserToResponseData(user)
}

func (s *accountUserService) CreateUser(payload *types.ReqCreateUser) (*types.RespUserDetail, int, error) {
	user, err := s.GetUserByPhone(payload.Phone)
	if user != nil {
		return nil, render.DataExistError, errors.New(render.GetCodeMsg(render.DataExistError))
	}
	user, err = s.GetUserByEmail(payload.Email)
	if user != nil {
		return nil, render.AccountEmailExistsError, errors.New(render.GetCodeMsg(render.AccountEmailExistsError))
	}
	user = model.NewUser()
	err = copier.Copy(user, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	user, err = model.CreateUser(s.db, user)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.CreateError, err
	}
	return s.transUserToResponseData(user)
}

func (s *accountUserService) DeleteUserByUUID(uuid string) (int, error) {
	go func() {
		// todo 异步清理用户的其他数据
	}()
	err := model.DeleteUser(s.db, uuid)
	if err != nil {
		log.GetLogger().Error(err)
		return render.QueryError, err
	}
	return render.OK, nil
}

func (s *accountUserService) GetUserByPhone(phone string) (*model.User, error) {
	return model.GetUserByPhone(s.db, phone)
}

func (s *accountUserService) GetUserByUUID(uuid string) (*model.User, error) {
	return model.GetUserByUUID(s.db, uuid)
}

func (s *accountUserService) GetUserByEmail(email string) (*model.User, error) {
	return model.GetUserByEmail(s.db, email)
}

func (s *accountUserService) UpdateUserLastLogin(uuid string) error {
	return model.UpdateUserLastLoginAt(s.db, uuid)
}

func (s *accountUserService) transUserToResponseData(user *model.User) (*types.RespUserDetail, int, error) {
	result := new(types.RespUserDetail)
	err := copier.Copy(result, user)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	result.Phone = user.PhoneMask()
	return result, render.OK, nil
}

func (s *accountUserService) UpdateUserByUUID(uuid string, payload *types.ReqUpdateUser) (int, error) {
	user, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return render.DataNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return render.QueryError, err
		}
	}
	err = copier.Copy(user, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return render.DBAttributesCopyError, err
	}
	err = s.db.Save(&user).Error
	if err != nil {
		log.GetLogger().Error(err)
		return render.DataExistError, err
	}
	return render.OK, nil
}

func NewAccountUserService() IAccountUserService {
	db := model.GetDB()
	return &accountUserService{
		db: db,
	}
}
