import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { InterviewDto } from 'src/app/components/dto/interview.dto';
import { InterviewService } from 'src/app/services/interview/interview.service';

@Component({
  selector: 'app-add-interview',
  templateUrl: './add-interview.component.html',
  styleUrls: ['./add-interview.component.css']
})
export class AddInterviewComponent implements OnInit {

  public addInterviewForm: FormGroup;
  public id: FormControl;
  public userId: FormControl;
  public jobId: FormControl;
  public position: FormControl;
  public title: FormControl;
  public yearOfInterview: FormControl;
  public hrInterview: FormControl;
  public technicalInterview: FormControl;

  constructor(private fb: FormBuilder,private interviewService: InterviewService,  public dialogRef: MatDialogRef<AddInterviewComponent>) {
      this.id= new FormControl("", [Validators.required]);
      this.userId= new FormControl("", [Validators.required]);
      this.jobId= new FormControl("", [Validators.required]);
      this.position= new FormControl("", [Validators.required]);
      this.title=new FormControl("", [Validators.required]);
      this.yearOfInterview= new FormControl("", [Validators.required]);
      this.hrInterview= new FormControl("", [Validators.required]);
      this.technicalInterview= new FormControl("", [Validators.required]);

      this.addInterviewForm = new FormGroup({
        'id': this.id,
        'userId': this.userId,
        'jobId': this.jobId,
        'position': this.position,
        'title': this.title,
        'yearOfInterview': this.yearOfInterview,
        'hrInterview': this.hrInterview,
        'techniaclInterview': this.technicalInterview,
      })
   }

  ngOnInit(): void {
  }

  
  onAdd() {
    if (this.addInterviewForm.valid) {
      let dto: InterviewDto = {id: this.id.value, userId: this.userId.value, jobId: this.jobId.value ,position: this.position.value, title: this.title.value, yearOfInterview: this.yearOfInterview.value, hrInterview: this.hrInterview.value, technicalInterview: this.technicalInterview.value}
      this.interviewService.createNewInterview(dto).subscribe((response) => {
        console.log("interview added")
        this.dialogRef.close()
      })

    }
  }


}