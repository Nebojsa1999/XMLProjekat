import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { AddInterviewComponent } from 'src/app/components/add-interview/add-interview.component';
//import { Job } from 'src/app/models/job';
//import { JobService } from 'src/app/services/job/job.service';
import { InterviewService } from 'src/app/services/interview/interview.service';
import { InterviewDto } from '../dto/interview.dto';
import { Interview } from 'src/app/models/interview'
import { WageDto } from '../dto/wage.dto'
//import { NewJobtDto } from '../dto/new-job-dto'

@Component({
  selector: 'app-interview',
  templateUrl: './interview.component.html',
  styleUrls: ['./interview.component.css']
})
export class InterviewComponent implements OnInit {
  interviews: Interview[] = [];
  allInterviews: Interview[] = [];

  public newInterview: InterviewDto = {
    id: "",
    companyId: "",
    position: "",
    title: "",
    yearOfInterview: "",
    hrInterview: "",
    technicalInterview: "",
  }

  constructor(private interviewService: InterviewService, public matDialog: MatDialog) { }

  ngOnInit(): void {
    this.interviewService.getInterviews().subscribe(
      response => {
        this.interviews = response.interviews;
        console.log(this.interviews);
        this.allInterviews = response.interviews;
        console.log(this.allInterviews);
;
      }
    )
  }

  openNewInterviewDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    const modalDialog = this.matDialog.open(AddInterviewComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }
}
