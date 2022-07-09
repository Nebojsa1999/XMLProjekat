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


  getInterviews(): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/agent-app/job/interview");
  }

  createNewInterview(dto: NewInterviewDto) {
    return this._http.post(this.applicationURL + "/agent-app/job/interview/create", dto);
  }
}
