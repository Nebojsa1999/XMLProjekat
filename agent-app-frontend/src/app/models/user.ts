export interface User {
    id: string;
    role: string;
    ownedCompanyId: string;
    issuedCompanyRequestId: string;
    username: string;
    password: string;
    firstName: string;
    lastName: string;
    email: string;
    phone: string;
    gender: string;
    dateOfBirth: Date;
    biography: string;
    workExperience: string;
    education: string;
    skills: string;
    interests: string;
}
