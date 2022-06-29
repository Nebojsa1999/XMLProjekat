import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { AddJobComponent } from 'src/app/modal/add-job/add-job.component';
import { Job } from 'src/app/model/job.model';
import { JobService } from 'src/app/service/job-service/job.service';

@Component({
  selector: 'app-jobs',
  templateUrl: './jobs.component.html',
  styleUrls: ['./jobs.component.css']
})
export class JobsComponent implements OnInit {
  jobs: Job[] = [];
  allJobs: Job[] = [];
  public searchText: string = "";

  constructor(private jobService: JobService, public matDialog: MatDialog) { }

  ngOnInit(): void {
    this.jobService.getJobs().subscribe(
      response => {
        this.jobs = response.jobs;
        console.log(this.jobs);
        this.allJobs = response;
      }
    )
  }

  openNewJobDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = "add-comment-modal";
    dialogConfig.height = "550px";
    dialogConfig.width = "32%";
    const modalDialog = this.matDialog.open(AddJobComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }

  search() {
    var searchWords = this.searchText.toLowerCase().split(" ");
    this.jobs = [];
    const result = this.allJobs.filter(job => {
      var isFound = true;
      for (let sw of searchWords) {
        if (!(job.description.toLowerCase().includes(sw) || job.position.toLowerCase().includes(sw)   ) ) {
          isFound = false;
        }
      }
      if (isFound == true) {
        this.jobs.push(job);
      }
    });
  }

}
