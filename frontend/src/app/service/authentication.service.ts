import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable } from 'rxjs';
import { Credentials } from '../model/credentials';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private loginURL = 'http://localhost:8000/user/login';

  private dislinktAppToken = null;

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
    this.dislinktAppToken = null;

    this.router.navigate(['/login']);
  }

  getDislinktAppToken() {
    return localStorage.getItem('dislinktAppToken');
  }

  isDislinktAppTokenPresent(): boolean {
    return this.dislinktAppToken != undefined && this.dislinktAppToken != null;
  }

  getAuthStatus(): boolean {
    return !!localStorage.getItem('dislinktAppToken');
  }

  hasDislinktAppTokenExpired(): boolean {
    var dislinktAppToken = localStorage.getItem('dislinktAppToken');

    if (dislinktAppToken != null) {
      if (this.jwtHelper.isTokenExpired(dislinktAppToken || '{}')) {
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