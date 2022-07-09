import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { InterviewDto } from 'src/app/components/dto/interview.dto';
import { InterviewService } from 'src/app/services/interview/interview.service';
import { NewInterviewDto } from '../dto/new-interview.dto';


export interface DialogData {
  companyId: string;
}
@Component({
  selector: 'app-add-interview',
  templateUrl: './add-interview.component.html',
  styleUrls: ['./add-interview.component.css']
})
export class AddInterviewComponent implements OnInit {

  public addInterviewForm: FormGroup;

  public position: FormControl;
  public title: FormControl;
  public yearOfInterview: FormControl;
  public hrInterview: FormControl;
  public technicalInterview: FormControl;

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData,private fb: FormBuilder,private interviewService: InterviewService,  public dialogRef: MatDialogRef<AddInterviewComponent>) {
      this.position= new FormControl("", [Validators.required]);
      this.title=new FormControl("", [Validators.required]);
      this.yearOfInterview= new FormControl("", [Validators.required]);
      this.hrInterview= new FormControl("", [Validators.required]);
      this.technicalInterview= new FormControl("", [Validators.required]);

      this.addInterviewForm = new FormGroup({
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
      let dto: NewInterviewDto = {companyId: this.data.companyId ,position: this.position.value, title: this.title.value, yearOfInterview: this.yearOfInterview.value, hrInterview: this.hrInterview.value, technicalInterview: this.technicalInterview.value}
      this.interviewService.createNewInterview(dto).subscribe((response) => {
        console.log("interview added")
        this.dialogRef.close()
      })

    }
  }


}