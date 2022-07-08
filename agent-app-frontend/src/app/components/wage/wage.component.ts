import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { AddWageComponent } from 'src/app/components/add-wage/add-wage.component';
//import { Job } from 'src/app/models/job';
//import { JobService } from 'src/app/services/job/job.service';
import { WageService } from 'src/app/services/wage/wage.service';
import { Wage } from 'src/app/models/wage'
import { InterviewDto } from '../dto/interview.dto';
import { WageDto } from '../dto/wage.dto'
//import { NewJobtDto } from '../dto/new-job-dto'

@Component({
  selector: 'app-wage',
  templateUrl: './wage.component.html',
  styleUrls: ['./wage.component.css']
})
export class WageComponent implements OnInit {
  wages: Wage[] = [];
  allWages: Wage[] = [];

  public newWage: WageDto = {
    id: "",
    userId: "",
    jobId: "",
    position: "",
    engagement: "",
    experienceLevel: "",
    netoWage: ""
  }

  constructor(private wageService: WageService, public matDialog: MatDialog) { }

  ngOnInit(): void {
    this.wageService.getWages().subscribe(
      response => {
        this.wages = response.wages;
        console.log(this.wages);
        this.allWages = response.wages;
        console.log(this.allWages);
;
      }
    )
  }

  openNewWageDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    const modalDialog = this.matDialog.open(AddWageComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }
}