import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { RegisterDTO } from '../components/dto/register.dto';
import { Credentials } from '../model/credentials';
import { User } from '../model/user';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private loginURL = 'http://localhost:8000/user/login';
  private applicationURL = environment.apiUrl;

  private dislinktAppToken = null;
  private accessToken = localStorage.getItem('dislinktAppToken');

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

    localStorage.removeItem('dislinktAppToken');
    localStorage.removeItem('id');
    localStorage.removeItem('username');
    localStorage.removeItem('exp');
    this.dislinktAppToken = null;
    this.router.navigate(['/login']);
  }

  getDislinktAppToken() {
    return localStorage.getItem('dislinktAppToken');
  }
  getToken() {
    return this.accessToken;
  }

  isLoggedIn() {
    return this.accessToken !== undefined && this.accessToken !== null;
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

  register(user:RegisterDTO): Observable<any>{
    return this.httpClient.post<any>(this.applicationURL + "/user/register", user);
  }
}