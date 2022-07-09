import { HttpErrorResponse } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { NewJobDto } from 'src/app/components/dto/new-job-dto';
import { JobService } from 'src/app/services/job/job.service';


export interface DialogData {
  companyId: string;
}
@Component({
  selector: 'app-add-job',
  templateUrl: './add-job.component.html',
  styleUrls: ['./add-job.component.css']
})
export class AddJobComponent implements OnInit {

  public addJobForm: FormGroup;

  public position: FormControl;
  public description: FormControl;
  public requirements: FormControl;

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData, private fb: FormBuilder, 
    private jobService: JobService, public dialogRef: MatDialogRef<AddJobComponent>, 
    private snackBar: MatSnackBar) {

      this.position= new FormControl("", [Validators.required]);
      this.description=new FormControl("", [Validators.required]);
      this.requirements= new FormControl("", [Validators.required]);

      this.addJobForm = new FormGroup({
        'position': this.position,
        'description': this.description,
        'requirements': this.requirements,
      })
   }

  ngOnInit(): void {
  }

  
  onAdd(): void {
    if (this.addJobForm.valid) {
      let dto: NewJobDto = { companyId: this.data.companyId, createdAt: new Date(), 
        position: this.position.value, description: this.description.value, 
        requirements: this.requirements.value };
      
      this.jobService.createNewJob(dto).subscribe(
        (response) => {
          console.log(response);
          this.snackBar.open(response, "Close", { duration: 5000 });

          this.dialogRef.close();
        },
        (error: HttpErrorResponse) => {
          console.log("Error on creating a job: ", error.error);
          this.snackBar.open(error.error, "Close", { duration: 5000} );
        }
      );
    }
  }
}
