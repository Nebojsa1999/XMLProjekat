import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/model/user';
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

  getAllProfiles(): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/user");
  }

  searchProfiles(param: string): Observable<any> {
    return this._http.post<any>(this.applicationURL + "/user/search",JSON.stringify(param));
  }

  getProfile(id: string): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/user/" + id);
  }

  updateProfile(id: string, updatedProfile: User): Observable<any> {
    return this._http.put<any>(this.applicationURL + "/user/" + id, JSON.stringify(updatedProfile));
  }

  isUserPrivate(id:string) : Observable<any>{
    return this._http.get<any>(this.applicationURL + "/user/" + id + "/is-private");
  }
}
