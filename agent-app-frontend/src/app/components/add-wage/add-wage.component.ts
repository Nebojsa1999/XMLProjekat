import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { WageDto } from 'src/app/components/dto/wage.dto';
import { WageService } from 'src/app/services/wage/wage.service';
import { NewWageDto } from '../dto/new-wage.dto';

export interface DialogData {
  companyId: string;
}

@Component({
  selector: 'app-add-wage',
  templateUrl: './add-wage.component.html',
  styleUrls: ['./add-wage.component.css']
})
export class AddWageComponent implements OnInit {

  public addWageForm: FormGroup;


  public position: FormControl;
  public engagement: FormControl;
  public experienceLevel: FormControl;
  public netoWage: FormControl;

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData,private fb: FormBuilder,private wageService: WageService,  public dialogRef: MatDialogRef<AddWageComponent>) {

      this.position= new FormControl("", [Validators.required]);
      this.engagement=new FormControl("", [Validators.required]);
      this.experienceLevel= new FormControl("", [Validators.required]);
      this.netoWage = new FormControl("", [Validators.required]);

      this.addWageForm = new FormGroup({
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
      let dto: NewWageDto = { companyId: this.data.companyId, position: this.position.value, engagement: this.engagement.value, experienceLevel: this.experienceLevel.value, netoWage: this.netoWage.value}
      this.wageService.createNewWage(dto).subscribe((response) => {
        console.log("wage added")
        this.dialogRef.close()
      })

    }
  }


}
