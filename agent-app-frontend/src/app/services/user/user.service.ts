import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { catchError, tap} from 'rxjs/operators';

import { User } from 'src/app/models/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private userEntityURL = 'http://localhost:8001/agent-app/user'

  constructor(private httpClient: HttpClient) { }

  getUserWith(id: string): Observable<User> {
    return this.httpClient.get<User>(this.userEntityURL + '/' + id).pipe(
      tap(data => 
        console.log('User with id = "' + id + '": ' + JSON.stringify(data))), 
        catchError(this.handleError));
  }

  private handleError(err: HttpErrorResponse) {
    console.log(err.message);

    return Observable.throw(err.message);
  }
}
