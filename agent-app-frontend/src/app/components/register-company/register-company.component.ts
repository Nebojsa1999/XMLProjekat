import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { NewCompanyDto } from '../dto/new-company.dto';

@Component({
  selector: 'app-register-company',
  templateUrl: './register-company.component.html',
  styleUrls: ['./register-company.component.css']
})
export class RegisterCompanyComponent implements OnInit {
  id:any;
  public form: FormGroup;
  public name: FormControl;
  public address: FormControl;
  public email: FormControl;
  public phone: FormControl;
  public areaOfWork: FormControl;
  public description: FormControl;
  public workCulture: FormControl;

  constructor(private router: Router, private companyService: CompanyService, 
    private snackBar: MatSnackBar) {
    this.name= new FormControl("",[Validators.required]);
    this.address= new FormControl("",[Validators.required]);
    this.email= new FormControl("",[Validators.required]);
    this.phone= new FormControl("",[Validators.required]);
    this.areaOfWork= new FormControl("",[Validators.required]);
    this.description= new FormControl("",[Validators.required]);
    this.workCulture= new FormControl("",[Validators.required]);

    this.form = new FormGroup({
      'name' : this.name,
      'address' : this.address,
      'email' : this.email,
      'phone' : this.phone,
      'areaOfWork' : this.areaOfWork,
      'description' : this.description,
      'workCulture' : this.workCulture,
    })

   }



  ngOnInit(): void {
    this.id=localStorage.getItem("id");
  }

  register() {
    if (this.form.valid) {
      let dto: NewCompanyDto = { name: this.name.value, address: this.address.value, 
        email: this.email.value, phone: this.phone.value, areaOfWork: this.areaOfWork.value, 
        description: this.description.value, workCulture: this.workCulture.value, status: "Pending", 
        ownerId:this.id };

      this.companyService.registerCompany(dto).subscribe(
        (response) => {
          console.log(response)
          console.log(dto);
          this.snackBar.open(response, "Close", { duration: 5000 });
          
          this.router.navigate(['/companies']).then(()=>{ location.reload() });
        },
        (error: HttpErrorResponse) => {
          console.log("Error on creating a company registration request: ", error.error);
          this.snackBar.open(error.error, "Close", { duration: 5000});
        });
    }
  }
}
