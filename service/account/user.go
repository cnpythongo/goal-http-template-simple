package account

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserList(req *types.ReqGetUserList) (*types.RespGetUserList, int)
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(payload *model.User) (*model.User, error)
	DeleteUserByUUID(uuid string) error

	UpdateUserLastLogin(id int64) error
}

type userService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func (s *userService) UpdateUserLastLogin(id int64) error {
	return model.UpdateUserLastLoginAt(s.db, id)
}

func (s *userService) DeleteUserByUUID(uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserByPhone(phone string) (*model.User, error) {
	return model.GetUserByPhone(s.db, phone)
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) CreateUser(payload *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserList(req *types.ReqGetUserList) (*types.RespGetUserList, int) {
	page := req.Page
	size := req.Size
	conditions := make(map[string]interface{})
	rows, count, err := model.GetUserList(s.db, page, size, conditions)
	if err != nil {
		return nil, response.DBQueryError
	}

	result := make([]*types.RespUser, 0)
	for _, row := range rows {
		item := new(types.RespUser)
		err = copier.Copy(item, row)
		if err != nil {
			return nil, response.DBAttributesCopyError
		}
		result = append(result, item)
	}
	resp := &types.RespGetUserList{
		Page:   page,
		Total:  utils.TotalPage(size, count),
		Count:  count,
		Result: result,
	}
	return resp, response.SuccessCode
}

func NewUserService(ctx *gin.Context) IUserService {
	db := model.GetDB()
	return &userService{
		ctx: ctx,
		db:  db,
	}
}
