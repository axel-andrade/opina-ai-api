package cockroach_models

type VoterModel struct {
	BaseModel
	FullName  string `gorm:"type:varchar(255);not null" json:"full_name"`
	Cellphone string `gorm:"type:varchar(20);unique;not null" json:"cellphone"`
}

func (VoterModel) TableName() string {
	return "voters"
}
