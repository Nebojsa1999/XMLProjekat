export interface NewJobDto {
    companyId: string;
    createdAt: Date;
    position: string;
    description: string;
    requirements: string;
}

export interface NewJobForSendingToDislinktAppDto {
    jobOffersAPIToken: string;
    job: NewJobDto;
}
