import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { AddJobComponent } from 'src/app/components/add-job/add-job.component';
import { Job, JobWithNameOfCompany } from 'src/app/models/job';
import { JobService } from 'src/app/services/job/job.service';
import { CommentDto } from '../dto/comment.dto';
import { InterviewDto } from '../dto/interview.dto';
import { WageDto } from '../dto/wage.dto'
import { NewJobDto } from '../dto/new-job-dto'
import { Router } from '@angular/router';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { HttpErrorResponse } from '@angular/common/http';
import { iif } from 'rxjs';

@Component({
  selector: 'app-job',
  templateUrl: './job.component.html',
  styleUrls: ['./job.component.css']
})
export class JobComponent implements OnInit {
  jobs: Job[] = [];

  public newJob: NewJobDto = {
    companyId: "",
    createdAt: new Date(),
    position: "",
    description: "",
    requirements: "",
  }

  isCompanyOwner: boolean = false;
  currentUserRole: string = "";
  ownedCompanyId: string = "";

  constructor(private jobService: JobService, private companyService: CompanyService, 
    public matDialog: MatDialog, private router: Router) { }

  ngOnInit(): void {
    let role = localStorage.getItem("role");
    if (role != null) {
      this.currentUserRole = role;
    }

    let ownedCompanyIdFromLocalStorage = localStorage.getItem("ownedCompanyId");
    if (ownedCompanyIdFromLocalStorage != null) {
      this.ownedCompanyId = ownedCompanyIdFromLocalStorage;
    }

    this.isCompanyOwnerCheck();

    this.jobService.getJobs().subscribe(
      response => {
        this.jobs = response;
        console.log(this.jobs);
      },
      (error: HttpErrorResponse) => {
        console.log("Error while getting jobs:\n" + error.error);
      }
    )
  }

  isCompanyOwnerCheck(): void {
    if (this.currentUserRole === 'CompanyOwner') {
      this.isCompanyOwner = true;
    }
    else {
      this.isCompanyOwner = false;
    }

    console.log("User is owner: ", this.isCompanyOwner);
  }

  openNewJobDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    dialogConfig.data = { companyId: this.ownedCompanyId };
    const modalDialog = this.matDialog.open(AddJobComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(
      result => {
        location.reload();
      },
      error => {
        console.log("Job was not posted on Dislinkt application due to an error!\n\n", error);
      });
  }
}
