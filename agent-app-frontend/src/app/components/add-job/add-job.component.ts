import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { NewJobDto } from 'src/app/components/dto/new-job-dto';
import { JobService } from 'src/app/services/job/job.service';

@Component({
  selector: 'app-add-job',
  templateUrl: './add-job.component.html',
  styleUrls: ['./add-job.component.css']
})
export class AddJobComponent implements OnInit {

  public addJobForm: FormGroup;
  public id: FormControl;
  public userId: FormControl;
  public position: FormControl;
  public description: FormControl;
  public requirements: FormControl;

  constructor(private fb: FormBuilder,private jobService: JobService,  public dialogRef: MatDialogRef<AddJobComponent>) {
      this.id= new FormControl("", [Validators.required]);
      this.userId= new FormControl("", [Validators.required]);
      this.position= new FormControl("", [Validators.required]);
      this.description=new FormControl("", [Validators.required]);
      this.requirements= new FormControl("", [Validators.required]);

      this.addJobForm = new FormGroup({
        'id': this.id,
        "userId": this.userId,
        'position': this.position,
        'description': this.description,
        'requirements': this.requirements,
      })
   }

  ngOnInit(): void {
  }

  
  onAdd() {
    if (this.addJobForm.valid) {
      let dto: NewJobDto = {id: this.id.value, userId: this.userId.value, position: this.position.value, description: this.description.value, requirements: this.requirements.value}
      this.jobService.createNewJob(dto).subscribe((response) => {
        console.log("job added")
        this.dialogRef.close()
      })

    }
  }


}
