import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { User } from 'src/app/models/user';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { UserService } from 'src/app/services/user/user.service';
import { CompanyRequestDto } from '../dto/company-request.dto';
import { CompanyDto } from '../dto/company.dto';

@Component({
  selector: 'app-register-company-requests',
  templateUrl: './register-company-requests.component.html',
  styleUrls: ['./register-company-requests.component.css']
})
export class RegisterCompanyRequestsComponent implements OnInit {
  requests: CompanyRequestDto[] = []
  user: User={
    id: '',
    role: '',
    ownedCompanyId: '',
    issuedCompanyRequestId: '',
    username: '',
    password: '',
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    gender: '',
    dateOfBirth: new Date,
    biography: '',
    workExperience: '',
    education: '',
    skills: '',
    interests: ''
  }
  userId:string="";
  

  constructor(private companyService: CompanyService, private snackBar: MatSnackBar,private userService:UserService) { }

  ngOnInit(): void {
    let userIdInLocalStorage=localStorage.getItem("id");
    if(userIdInLocalStorage != null){
      this.userId=userIdInLocalStorage;
    }
    this.companyService.getCompanyRequests().subscribe((response) => {
      this.requests = response;
      console.log(response);
    })
  }

  approveRequest(id: string,request:CompanyRequestDto){
    request.status="Approved";
    this.companyService.approveRequest(id,request).subscribe(
      (response) => {
        console.log(response);
        let companyDto:CompanyDto = {
          id: request.id,
          ownerId: request.ownerId,
          name: request.name,
          address: request.address,
          email: request.email,
          phone: request.phone,
          areaOfWork: request.areaOfWork,
          description: request.description,
          workCulture: request.workCulture
        }
        
        this.createCompany(companyDto);
        this.getUser(request.ownerId);
     

    },
    (error: HttpErrorResponse) => {
      console.log("Error on approving company registration request: ", error.error);
      this.snackBar.open(error.error, "Close", { duration: 5000 });
    });
  }
  
  createCompany(dto:CompanyDto): void {
    this.companyService.createCompany(dto).subscribe(
      (response) => {
        console.log(response);
        this.snackBar.open(response, "Close", { duration: 5000 });
        
   
      },
      (error: HttpErrorResponse) => {
        console.log("Error on creating company: ", error.error);
        this.snackBar.open(error.error, "Close", { duration: 5000 });
      }
    );
  }

  getUser(id:string):void{
    this.userService.getUserWith(id).subscribe(
      (response) => {
        console.log(response);
        this.user=response;
        this.updateUser(this.user.id);
      },
      (error: HttpErrorResponse) => {
        console.log("Error getting user", error.error);
        this.snackBar.open(error.error, "Close", { duration: 5000 });
      }
    )
  }

  updateUser(id:string){
    this.user.role="CompanyOwner";
    console.log(this.user);
    this.userService.updateUser(id,this.user).subscribe(
      (response) => {
        console.log(response);
        
      },
      (error: HttpErrorResponse) => {
        console.log("Error updating user", error.error);
        this.snackBar.open(error.error, "Close", { duration: 5000 });
      }
    )
  }
}
