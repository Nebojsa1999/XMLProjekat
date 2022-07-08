import { Component, OnInit } from '@angular/core';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { CompanyRequestDto } from '../dto/company-request.dto';
import { CompanyDto } from '../dto/company.dto';

@Component({
  selector: 'app-register-company-requests',
  templateUrl: './register-company-requests.component.html',
  styleUrls: ['./register-company-requests.component.css']
})
export class RegisterCompanyRequestsComponent implements OnInit {
  requests: CompanyRequestDto[] = []
  

  constructor(private companyService: CompanyService) { }

  ngOnInit(): void {
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
    })
  }

createCompany(dto:CompanyDto):void{
  this.companyService.createCompany(dto).subscribe(
    (response)=>{
      console.log(response);
  })
  window.location.reload();
}


}
