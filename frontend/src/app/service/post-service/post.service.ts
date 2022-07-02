import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PostService {

  private applicationURL = environment.apiUrl;
  constructor(private _http: HttpClient) { }

  getPostsByUser(ownerId: string): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/user/"+ ownerId + "/post");
  }
}
