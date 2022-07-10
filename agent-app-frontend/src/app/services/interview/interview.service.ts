import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { InterviewDto } from 'src/app/components/dto/interview.dto';
import { NewInterviewDto } from 'src/app/components/dto/new-interview.dto';
import { Interview } from 'src/app/models/interview';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class InterviewService {

  private applicationURL = environment.apiUrl;

  constructor(private _http: HttpClient) { }


  getInterviews(id:string): Observable<any> {
    return this._http.get<any>(`${this.applicationURL}/agent-app/company/interview/${id}`);
  }

  createNewInterview(dto: NewInterviewDto): Observable<any> {
    return this._http.post<any>(this.applicationURL + "/agent-app/job/interview/create", dto);
  }
}
