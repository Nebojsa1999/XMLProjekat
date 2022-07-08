import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { User } from 'src/app/models/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private userEntityURL = 'http://localhost:8001/agent-app/user'

  constructor(private httpClient: HttpClient) { }

  getUserWith(id: string): Observable<any> {
    return this.httpClient.get<any>(this.userEntityURL + '/' + id);
  }

  registerAsACommonUser(newCommonUser: User): Observable<any> {
    const headers = new HttpHeaders({
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    });

    return this.httpClient.post<any>(this.userEntityURL + '/register', JSON.stringify(newCommonUser), 
      { headers: headers });
  }
}
