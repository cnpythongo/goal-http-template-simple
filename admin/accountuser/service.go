package accountuser

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"strings"
)

type IUserService interface {
	GetUserList(req *ReqGetUserList) (*RespGetUserList, int, error)
	GetUserDetail(uuid string) (*RespUserDetail, int, error)
	CreateUser(payload *ReqCreateUser) (*RespUserDetail, int, error)
	DeleteUserByUUID(uuid string) (int, error)
	UpdateUserByUUID(uuid string, payload *ReqUpdateUser) (int, error)
	GetUserByPhone(phone string) (*model.User, int, error)
	GetUserByUUID(uuid string) (*model.User, int, error)
	GetUserByEmail(email string) (*model.User, int, error)
	GetUserProfile(userId uint64) (*model.UserProfile, int, error)
	UpdateUserProfile(payload *ReqUpdateUserProfile) (int, error)
}

type userService struct {
}

func NewUserService() IUserService {
	return &userService{}
}

func (s *userService) GetUserList(req *ReqGetUserList) (*RespGetUserList, int, error) {
	var query []string
	var args []interface{}
	if req.Email != "" {
		query = append(query, "email like concat('%', ?, '%')")
		args = append(args, req.Email)
	}
	if req.Phone != "" {
		query = append(query, "phone like concat('%', ?, '%')")
		args = append(args, req.Phone)
	}
	if req.Nickname != "" {
		query = append(query, "nickname like concat('%s', ?, '%')")
		args = append(args, req.Nickname)
	}
	if len(req.Status) > 0 {
		query = append(query, "status IN ?")
		args = append(args, req.Status)
	}
	if req.IsAdmin != nil && (*req.IsAdmin == 1 || *req.IsAdmin == 0) {
		query = append(query, "is_admin = ?")
		args = append(args, req.IsAdmin)
	}
	if req.UUID != "" {
		query = append(query, "uuid = ?")
		args = append(args, req.UUID)
	}
	if req.CreatedAtStart > 0 {
		query = append(query, "created_at >= ?")
		args = append(args, req.CreatedAtStart)
	}
	if req.CreatedAtEnd > 0 {
		query = append(query, "created_at <= ?")
		args = append(args, req.CreatedAtEnd)
	}
	if req.LastLoginAtStart > 0 {
		query = append(query, "last_login_at >= ?")
		args = append(args, req.LastLoginAtStart)
	}
	if req.LastLoginAtEnd > 0 {
		query = append(query, "last_login_at <= ?")
		args = append(args, req.LastLoginAtEnd)
	}
	queryStr := strings.Join(query, " AND ")
	rows, total, err := model.GetUserList(model.GetDB(), req.Page, req.Limit, queryStr, args)
	if err != nil {
		return nil, render.QueryError, err
	}

	result := make([]*RespUserBasic, 0)
	for _, row := range rows {
		item := new(RespUserBasic)
		err = copier.Copy(item, row)
		if err != nil {
			log.GetLogger().Error(err)
			return nil, render.DBAttributesCopyError, err
		}
		result = append(result, item)
	}
	resp := &RespGetUserList{
		Page:   req.Page,
		Limit:  req.Limit,
		Total:  total,
		Result: result,
	}
	return resp, render.OK, nil
}

func (s *userService) GetUserDetail(uuid string) (*RespUserDetail, int, error) {
	user, code, err := s.GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		} else {
			log.GetLogger().Error(err)
			return nil, code, err
		}
	}
	return s.transUserToResponseData(user)
}

func (s *userService) CreateUser(payload *ReqCreateUser) (*RespUserDetail, int, error) {
	user, _, err := s.GetUserByPhone(payload.Phone)
	if user != nil {
		return nil, render.DataExistError, errors.New(render.GetCodeMsg(render.DataExistError, nil))
	}

	user, _, err = s.GetUserByEmail(payload.Email)
	if user != nil {
		return nil, render.AccountEmailExistsError, errors.New(render.GetCodeMsg(render.AccountEmailExistsError, nil))
	}
	user = model.NewUser()
	err = copier.Copy(user, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}

	user, err = model.CreateUser(model.GetDB(), user)
	if err != nil {
		log.GetLogger().Error(fmt.Sprintf("model.CreateUser ==> %v", err))
		return nil, render.CreateError, err
	}

	return s.transUserToResponseData(user)
}

func (s *userService) DeleteUserByUUID(uuid string) (int, error) {
	go func() {
		// todo 异步清理用户的其他数据
	}()
	err := model.DeleteUser(model.GetDB(), uuid)
	if err != nil {
		log.GetLogger().Error(err)
		return render.DeleteError, err
	}
	return render.OK, nil
}

func (s *userService) GetUserByPhone(phone string) (*model.User, int, error) {
	user, err := model.GetUserByPhone(model.GetDB(), phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		log.GetLogger().Error(err)
		return nil, render.QueryError, err
	}
	return user, render.OK, nil
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, int, error) {
	user, err := model.GetUserByUUID(model.GetDB(), uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		log.GetLogger().Error(err)
		return nil, render.QueryError, err
	}
	return user, render.OK, nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, int, error) {
	user, err := model.GetUserByEmail(model.GetDB(), email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		log.GetLogger().Error(err)
		return nil, render.QueryError, err
	}
	return user, render.OK, nil
}

func (s *userService) transUserToResponseData(user *model.User) (*RespUserDetail, int, error) {
	result := new(RespUserDetail)
	err := copier.Copy(result, user)
	if err != nil {
		log.GetLogger().Error(err)
		return nil, render.DBAttributesCopyError, err
	}
	result.Phone = user.PhoneMask()
	return result, render.OK, nil
}

func (s *userService) UpdateUserByUUID(uuid string, payload *ReqUpdateUser) (int, error) {
	user, code, err := s.GetUserByUUID(uuid)
	if err != nil {
		return code, err
	}
	err = copier.Copy(user, payload)
	if err != nil {
		log.GetLogger().Error(err)
		return render.DBAttributesCopyError, err
	}
	err = model.GetDB().Save(&user).Error
	if err != nil {
		log.GetLogger().Error(err)
		return render.DataExistError, err
	}
	return render.OK, nil
}

func (s *userService) GetUserProfile(userId uint64) (*model.UserProfile, int, error) {
	pf, err := model.GetUserProfileByUserId(model.GetDB(), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		log.GetLogger().Errorf("model.GetUserProfileByUserId Error ==> %v", err)
		return nil, render.QueryError, err
	}
	return pf, render.OK, nil
}

func (s *userService) UpdateUserProfile(payload *ReqUpdateUserProfile) (int, error) {
	pf, code, err := s.GetUserProfile(payload.UserId)
	if err != nil {
		return code, err
	}

	_, err = model.UpdateUserProfile(model.GetDB(), pf)
	if err != nil {
		log.GetLogger().Errorf("model.UpdateUserProfile Error ==> %v", err)
		return render.UpdateError, err
	}
	return render.OK, nil
}
