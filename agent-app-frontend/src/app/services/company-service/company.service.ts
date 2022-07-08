import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CompanyRequestDto } from 'src/app/components/dto/company-request.dto';
import { CompanyDto } from 'src/app/components/dto/company.dto';
import { NewCompanyDto } from 'src/app/components/dto/new-company.dto';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {

  constructor(private http: HttpClient) { }

  registerCompany(dto: NewCompanyDto):Observable<any> {
    return this.http.post<any>(environment.apiUrl + "/agent-app/company-registration-request", dto);
  }

  getCompanyRequests() : Observable<CompanyRequestDto[]> {
    return this.http.get<CompanyRequestDto[]>(environment.apiUrl + "/agent-app/company-registration-request/pending");
  }

  approveRequest(id:string,dto:CompanyRequestDto): Observable<CompanyRequestDto>{
    return this.http.put<CompanyRequestDto>(`${environment.apiUrl}/agent-app/company-registration-request/${id}/update-by-administrator`,dto);
  }

  createCompany(dto:CompanyDto): Observable<any> {
    return this.http.post<any>(environment.apiUrl + "/agent-app/company/register", dto);
  }
  
  getCompanies():Observable<CompanyDto[]>{
    return this.http.get<CompanyDto[]>(environment.apiUrl + "/agent-app/company");
  }

  getCompany(id:string):Observable<CompanyDto>{
    return this.http.get<CompanyDto>(environment.apiUrl + "/agent-app/company/" + id);
  }

  updateCompany(id:string,dto:CompanyDto): Observable<CompanyDto>{
    return this.http.put<CompanyDto>(`${environment.apiUrl}/agent-app/company-registration-request/${id}/update-by-owner`,dto);
  }
  
}
