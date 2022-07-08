import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-register-company',
  templateUrl: './register-company.component.html',
  styleUrls: ['./register-company.component.css']
})
export class RegisterCompanyComponent implements OnInit {
  public form: FormGroup;
  public name: FormControl;
  public address: FormControl;
  public email: FormControl;
  public phone: FormControl;
  public areaOfWork: FormControl;
  public description: FormControl;
  public workCulture: FormControl;
  constructor(private router: Router, private companyService: CompanyService) { }

  ngOnInit(): void {
  }

}
