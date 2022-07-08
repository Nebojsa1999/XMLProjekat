import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { CompanyDto } from '../dto/company.dto';

@Component({
  selector: 'app-companies',
  templateUrl: './companies.component.html',
  styleUrls: ['./companies.component.css']
})
export class CompaniesComponent implements OnInit {

  companies : CompanyDto[] = [];

  constructor(private companyService: CompanyService, private router: Router) { }

  ngOnInit(): void {
    this.getCompanies();
  }

  getCompanies():void{
    this.companyService.getCompanies().subscribe((response) => {
      this.companies = response;
      console.log(this.companies);
    })
  }

  openCompanyPage(id: string): void {
    this.router.navigate(['company', id]);
  }

}
