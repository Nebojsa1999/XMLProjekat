package enums

type UserRole string

const (
	CommonUser    UserRole = "CommonUser"
	CompanyOwner           = "CompanyOwner"
	Administrator          = "Administrator"
)

type Gender string

const (
	Undefined Gender = ""
	Male             = "Male"
	Female           = "Female"
)

type CompanyRegistrationRequestStatus string

const (
	Pending  CompanyRegistrationRequestStatus = "Pending"
	Accepted                                  = "Accepted"
	Rejected                                  = "Rejected"
)
