import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import { JobService } from 'src/app/service/job-service/job.service';

@Component({
  selector: 'app-job-token-dialog',
  templateUrl: './job-token-dialog.component.html',
  styleUrls: ['./job-token-dialog.component.css']
})
export class JobTokenDialogComponent implements OnInit {

  constructor(private _route: ActivatedRoute,private jobService: JobService,  public dialogRef: MatDialogRef<JobTokenDialogComponent>) { }
  
  userId:any
  token: string="";

  ngOnInit(): void {
    this.userId = localStorage.getItem('id');
    console.log(this.userId);
    this.jobService.getConnectionToken(this.userId).subscribe((response) => {
      console.log(response)
      this.token=response.token;
    })
  }

}
