package enums

type UserRole string

const (
	UndefinedRole UserRole = ""
	CommonUser             = "CommonUser"
	CompanyOwner           = "CompanyOwner"
	Administrator          = "Administrator"
)

type Gender string

const (
	UndefinedGender Gender = ""
	Male                   = "Male"
	Female                 = "Female"
)

type CompanyRegistrationRequestStatus string

const (
	UndefinedCompanyRegistrationRequestStatus CompanyRegistrationRequestStatus = ""
	Pending                                                                    = "Pending"
	Accepted                                                                   = "Accepted"
	Rejected                                                                   = "Rejected"
)
