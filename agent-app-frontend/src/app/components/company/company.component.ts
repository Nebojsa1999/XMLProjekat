import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { CompanyDto } from '../dto/company.dto';

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

  companyForm: FormGroup;

  constructor(private fb: FormBuilder,public matDialog: MatDialog, private companyService: CompanyService, private authService: AuthService, private route: ActivatedRoute) {
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

  
}
