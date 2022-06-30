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

type Position string

const (
	UndefinedPosition Position = ""
	SoftwareDeveloper          = "Software Developer"
	Animator                   = "Animator"
	Administration             = "Administration"
	DataScience                = "Data Science"
	DevOps                     = "DevOps"
	ProjectManager             = "Project Manager"
)

type Engagement string

const (
	UndefinedEngagement Engagement = ""
	FullTime                       = "Full time"
	PartTime                       = "Part time"
	Praksa                         = "Praksa"
)

type ExperienceLevel string

const (
	UndefinedExperienceLevel ExperienceLevel = ""
	Junior                                   = "Junior"
	Medior                                   = "Medior"
	Senior                                   = "Senior"
)
