package domain

type OTPSession struct {
	ID        uint   `json:"id" gorm:"primaryKey,index"`
	OtpId     string `json:"otpid" gorm:"not null"`
	MobileNum string `json:"mobilenum" gorm:"not null"`
}
