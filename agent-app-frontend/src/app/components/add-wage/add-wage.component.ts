import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { WageDto } from 'src/app/components/dto/wage.dto';
import { WageService } from 'src/app/services/wage/wage.service';

@Component({
  selector: 'app-add-wage',
  templateUrl: './add-wage.component.html',
  styleUrls: ['./add-wage.component.css']
})
export class AddWageComponent implements OnInit {

  public id: FormControl;
  public userId: FormControl;
  public jobId: FormControl;
  public addWageForm: FormGroup;
  public position: FormControl;
  public engagement: FormControl;
  public experienceLevel: FormControl;
  public netoWage: FormControl;

  constructor(private fb: FormBuilder,private wageService: WageService,  public dialogRef: MatDialogRef<AddWageComponent>) {

      this.id= new FormControl("", [Validators.required]);
      this.userId= new FormControl("", [Validators.required]);
      this.jobId= new FormControl("", [Validators.required]);
      this.position= new FormControl("", [Validators.required]);
      this.engagement=new FormControl("", [Validators.required]);
      this.experienceLevel= new FormControl("", [Validators.required]);
      this.netoWage = new FormControl("", [Validators.required]);

      this.addWageForm = new FormGroup({
        'id': this.id,
        'userId': this.userId,
        'jobId': this.jobId,
        'position': this.position,
        'engagement': this.engagement,
        'experienceLevel': this.experienceLevel,
        'netoWage': this.netoWage
      })
   }

  ngOnInit(): void {
  }

  
  onAdd() {
    if (this.addWageForm.valid) {
      let dto: WageDto = {id: this.id.value, userId: this.userId.value, jobId: this.jobId.value, position: this.position.value, engagement: this.engagement.value, experienceLevel: this.experienceLevel.value, netoWage: this.netoWage.value}
      this.wageService.createNewWage(dto).subscribe((response) => {
        console.log("wage added")
        this.dialogRef.close()
      })

    }
  }


}
