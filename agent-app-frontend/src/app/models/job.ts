export interface Job {
    id: string;
    companyId: string;
    createdAt: Date;
    position: string;
    description: string;
    requirements: string;
}

export interface JobWithNameOfCompany {
    nameOfCompany: string;
    job: Job;
}
