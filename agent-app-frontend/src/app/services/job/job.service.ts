import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { NewJobDto } from 'src/app/components/dto/new-job-dto';
import { Job } from 'src/app/models/job';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class JobService {

  private applicationURL = environment.apiUrl;

  constructor(private _http: HttpClient) { }


  getJobs(): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/agent-app/job");
  }

  createNewJob(dto: NewJobDto) {
    return this._http.post(this.applicationURL + "/agent-app/job/create", dto);
  }
}
