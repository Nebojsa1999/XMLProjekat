import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable } from 'rxjs';

import { Credentials } from 'src/app/models/credentials';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private loginURL = 'http://localhost:8001/agent-app/user/login';

  private agentAppToken = null;

  constructor(private jwtHelper: JwtHelperService, private httpClient: HttpClient, 
    private router: Router) { }

  loginWith(credentials: Credentials): Observable<any> {
    const loginHeaders = new HttpHeaders({
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    });

    return this.httpClient.post(this.loginURL, JSON.stringify(credentials), { headers: loginHeaders });
  }

  logOut() {
    this.agentAppToken = null;

    this.router.navigate(['/login']);
  }

  getAgentAppToken() {
    return localStorage.getItem('agentAppToken');
  }

  isAgentAppTokenPresent(): boolean {
    return this.agentAppToken != undefined && this.agentAppToken != null;
  }

  getAuthStatus(): boolean {
    return !!localStorage.getItem('agentAppToken');
  }

  hasAgentAppTokenExpired(): boolean {
    var agentAppToken = localStorage.getItem('agentAppToken');

    if (agentAppToken != null) {
      if (this.jwtHelper.isTokenExpired(agentAppToken || '{}')) {
        return true;
      }
    }

    return false;
  }

  isARegisteredUser(): boolean {
    var role = localStorage.getItem('role');

    if (role != undefined && role != null) {
      return true;
    }

    return false;
  }
}
