import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private userEntityURL = 'http://localhost:8001/agent-app/user'

  constructor(private httpClient: HttpClient) { }

  getUserWith(id: string): Observable<any> {
    return this.httpClient.get<any>(this.userEntityURL + '/' + id);
  }
}
