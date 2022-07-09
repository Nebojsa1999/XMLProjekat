import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { AddJobComponent } from 'src/app/components/add-job/add-job.component';
import { Job } from 'src/app/models/job';
import { JobService } from 'src/app/services/job/job.service';
import { CommentDto } from '../dto/comment.dto';
import { InterviewDto } from '../dto/interview.dto';
import { WageDto } from '../dto/wage.dto'
import { NewJobDto } from '../dto/new-job-dto'

@Component({
  selector: 'app-job',
  templateUrl: './job.component.html',
  styleUrls: ['./job.component.css']
})
export class JobComponent implements OnInit {
  jobs: Job[] = [];
  allJobs: Job[] = [];

  public newJob: NewJobDto = {
    companyId: "",
    position: "",
    description: "",
    requirements: "",
  }

  constructor(private jobService: JobService, public matDialog: MatDialog) { }

  ngOnInit(): void {
    this.jobService.getJobs().subscribe(
      response => {
        this.jobs = response.jobs;
        console.log(this.jobs);
        this.allJobs = response.jobs;
        console.log(this.allJobs);
      }
    )
  }

  openNewJobDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    const modalDialog = this.matDialog.open(AddJobComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }
}
