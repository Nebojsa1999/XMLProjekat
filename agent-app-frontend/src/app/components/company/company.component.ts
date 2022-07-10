import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';

import { Interview } from 'src/app/models/interview';
import { AuthService } from 'src/app/services/auth/auth.service';
import { CommentService } from 'src/app/services/comment/comment.service';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { InterviewService } from 'src/app/services/interview/interview.service';
import { WageService } from 'src/app/services/wage/wage.service';
import { AddCommentComponent } from '../add-comment/add-comment.component';
import { AddInterviewComponent } from '../add-interview/add-interview.component';
import { AddWageComponent } from '../add-wage/add-wage.component';
import { CommentDto } from '../dto/comment.dto';
import { CompanyDto } from '../dto/company.dto';
import { InterviewDto } from '../dto/interview.dto';
import { WageDto } from '../dto/wage.dto';

@Component({
  selector: 'app-company',
  templateUrl: './company.component.html',
  styleUrls: ['./company.component.css']
})
export class CompanyComponent implements OnInit {

  company:CompanyDto = {
    id: '',
    ownerId: '',
    name: '',
    address: '',
    email: '',
    phone: '',
    areaOfWork: '',
    description: '',
    workCulture: ''
  }

  cid:any;

  isCompanyOwner:boolean = false;

  companyOwnerId:any;

  currentUserId:any;

  comments: CommentDto[] = []
  interviews: InterviewDto[] = []
  wages: WageDto[]=[]

  companyForm: FormGroup;

  constructor(private fb: FormBuilder,public matDialog: MatDialog, private companyService: CompanyService,private wageService:WageService,private interviewService: InterviewService,private commentService:CommentService, private route: ActivatedRoute) {
    this.companyForm = this.fb.group({
      name: ["", [Validators.required]],
      address: ["", []],
      email: ["", []],
      phone: ["", []],
      areaOfWork: ["", []],
      description: ["", []],
      workCulture: ["", []],
    });

   }

  ngOnInit(): void {
    this.cid =  this.route.snapshot.url[1].path;
    this.companyService.getCompany(this.cid).subscribe((response) => {
      this.company = response;
      this.companyForm.patchValue(this.company);
      this.companyOwnerId=response.ownerId;
      console.log(this.company)
      this.isCompanyOwnerCheck();
    })
    this.currentUserId=localStorage.getItem('id');

    console.log(this.cid)
    this.commentService.getComments(this.cid).subscribe((response) => {
      this.comments = response;
      if(this.comments==null){
        this.comments=[];
      }
      console.log(response)
    })
  

    this.interviewService.getInterviews(this.cid).subscribe((response) => {
      this.interviews = response;
      if(this.interviews==null){
        this.interviews=[];
      }
      console.log(response)
    })


    this.wageService.getWages(this.cid).subscribe((response) => {
      this.wages = response;
      if(this.wages==null){
        this.wages=[];
      }
      console.log(response)
    })
  

 
  }

  isCompanyOwnerCheck():void{
    if(JSON.stringify(this.currentUserId) == JSON.stringify(this.companyOwnerId)){
      this.isCompanyOwner=true;
    }
    else{
      this.isCompanyOwner=false;
    }
    console.log("User is owner:",this.isCompanyOwner);
  }

  onSubmit():void{
    this.company=this.companyForm.value;
    this.company.id=this.route.snapshot.url[1].path;
    this.company.ownerId=this.companyOwnerId;
    console.log(this.company);

    this.companyService.updateCompany(this.cid,this.company).subscribe(
      response => {
        console.log(response)

      }
    )
  }

  openNewCommentDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    dialogConfig.data = { companyId: this.cid }
    const modalDialog = this.matDialog.open(AddCommentComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }

  openNewInterviewDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    dialogConfig.data = { companyId: this.cid }
    const modalDialog = this.matDialog.open(AddInterviewComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }


  openNewWageDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    dialogConfig.data = { companyId: this.cid }
    const modalDialog = this.matDialog.open(AddWageComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }

  
}
