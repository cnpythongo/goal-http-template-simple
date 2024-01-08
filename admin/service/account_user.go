package service

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/jinzhu/copier"
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
	page := req.Page
	size := req.Size
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
	rows, count, err := model.GetUserList(s.db, page, size, queryStr, args)
	if err != nil {
		return nil, response.QueryError, err
	}

	result := make([]*types.RespUserBasic, 0)
	for _, row := range rows {
		item := new(types.RespUserBasic)
		err = copier.Copy(item, row)
		if err != nil {
			log.GetLogger().Error(err)
			return nil, response.DBAttributesCopyError, err
		}
		result = append(result, item)
	}
	resp := &types.RespGetUserList{
		Page:   page,
		Total:  utils.TotalPage(size, count),
		Count:  count,
		Result: result,
	}
	return resp, response.SuccessCode, nil
}

func (s *accountUserService) GetUserDetail(uuid string) (*types.RespUserDetail, int, error) {
	user, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.DataNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return nil, response.QueryError, err
		}
	}
	return s.transUserToResponseData(user)
}

func (s *accountUserService) CreateUser(payload *types.ReqCreateUser) (*types.RespUserDetail, int, error) {
	user, err := s.GetUserByPhone(payload.Phone)
	if user != nil {
		return nil, response.DataExistError, errors.New(response.GetCodeMsg(response.DataExistError))
	}
	user, err = s.GetUserByEmail(payload.Email)
	if user != nil {
		return nil, response.AccountEmailExistsError, errors.New(response.GetCodeMsg(response.AccountEmailExistsError))
	}
	user = model.NewUser()
	err = copier.Copy(user, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, response.DBAttributesCopyError, err
	}
	user, err = model.CreateUser(s.db, user)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, response.CreateError, err
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
		return response.QueryError, err
	}
	return response.SuccessCode, nil
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
		return nil, response.DBAttributesCopyError, err
	}
	result.Phone = user.PhoneMask()
	return result, response.SuccessCode, nil
}

func (s *accountUserService) UpdateUserByUUID(uuid string, payload *types.ReqUpdateUser) (int, error) {
	user, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.DataNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return response.QueryError, err
		}
	}
	err = copier.Copy(user, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return response.DBAttributesCopyError, err
	}
	err = s.db.Save(&user).Error
	if err != nil {
		log.GetLogger().Error(err)
		return response.DataExistError, err
	}
	return response.SuccessCode, nil
}

func NewAccountUserService() IAccountUserService {
	db := model.GetDB()
	return &accountUserService{
		db: db,
	}
}
