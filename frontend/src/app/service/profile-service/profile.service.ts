import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  private applicationURL = environment.apiUrl;

  constructor(private _http: HttpClient) { }

  getPublicProfiles(): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/user/public");
  }

  searchProfiles(param: string): Observable<any> {
    return this._http.post<any>(this.applicationURL + "/user/search/",param);
  }
}
