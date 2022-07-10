import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { NewWageDto } from 'src/app/components/dto/new-wage.dto';
import { WageDto } from 'src/app/components/dto/wage.dto';
import { Wage } from 'src/app/models/wage';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class WageService {

  private applicationURL = environment.apiUrl;

  constructor(private _http: HttpClient) { }


  getWages(id:string): Observable<any> {
    return this._http.get<any>(`${this.applicationURL}/agent-app/company/wage/${id}`);
  }

  createNewWage(dto: NewWageDto): Observable<any> {
    return this._http.post<any>(this.applicationURL + "/agent-app/job/wage/create", dto);
  }
}
