package account

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strings"
)

type IUserService interface {
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

type userService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func (s *userService) GetUserList(req *types.ReqGetUserList) (*types.RespGetUserList, int, error) {
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
		return nil, response.DBQueryError, err
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

func (s *userService) GetUserDetail(uuid string) (*types.RespUserDetail, int, error) {
	user, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.AccountUserNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return nil, response.AccountQueryUserError, err
		}
	}
	return s.transUserToResponseData(user)
}

func (s *userService) CreateUser(payload *types.ReqCreateUser) (*types.RespUserDetail, int, error) {
	user, err := s.GetUserByPhone(payload.Phone)
	if user != nil {
		return nil, response.AccountUserExistError, errors.New(response.GetCodeMsg(response.AccountUserExistError))
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
		return nil, response.AccountCreateError, err
	}
	return s.transUserToResponseData(user)
}

func (s *userService) DeleteUserByUUID(uuid string) (int, error) {
	go func() {
		// todo 异步清理用户的其他数据
	}()
	err := model.DeleteUser(s.db, uuid)
	if err != nil {
		log.GetLogger().Error(err)
		return response.AccountQueryUserError, err
	}
	return response.SuccessCode, nil
}

func (s *userService) GetUserByPhone(phone string) (*model.User, error) {
	return model.GetUserByPhone(s.db, phone)
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, error) {
	return model.GetUserByUUID(s.db, uuid)
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return model.GetUserByEmail(s.db, email)
}

func (s *userService) UpdateUserLastLogin(uuid string) error {
	return model.UpdateUserLastLoginAt(s.db, uuid)
}

func (s *userService) transUserToResponseData(user *model.User) (*types.RespUserDetail, int, error) {
	result := new(types.RespUserDetail)
	err := copier.Copy(result, user)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, response.DBAttributesCopyError, err
	}
	result.Phone = user.PhoneMask()
	return result, response.SuccessCode, nil
}

func (s *userService) UpdateUserByUUID(uuid string, payload *types.ReqUpdateUser) (int, error) {
	user, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.AccountUserNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return response.AccountQueryUserError, err
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
		return response.AccountUserExistError, err
	}
	return response.SuccessCode, nil
}

func NewUserService(ctx *gin.Context) IUserService {
	db := model.GetDB()
	return &userService{
		ctx: ctx,
		db:  db,
	}
}
