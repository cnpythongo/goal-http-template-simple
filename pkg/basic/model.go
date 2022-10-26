package basic

type BaseModel struct {
	ID        int64 `gorm:"primary_key;comment:流水ID" json:"-"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime;default:0;comment:数据创建时间" json:"-"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime;default:0;comment:数据更新时间" json:"-"`
	DeletedAt int64 `gorm:"column:deleted_at;default:0;comment:数据删除时间" json:"-"`
}
