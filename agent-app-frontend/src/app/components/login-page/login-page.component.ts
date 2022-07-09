import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import jwt_decode from 'jwt-decode';

import { MatSnackBar } from '@angular/material/snack-bar';

import { AuthService } from 'src/app/services/auth/auth.service';

import { Credentials } from 'src/app/models/credentials';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {
  form: any;
  isSubmitted: boolean = false;

  agentAppToken: string = '';

  constructor(private authService: AuthService, private formBuilder: FormBuilder, 
    private router: Router, private snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      username: ['', Validators.compose([Validators.required, Validators.minLength(4), 
        Validators.maxLength(32)])],
      password: ['', Validators.compose([Validators.required, Validators.minLength(5), 
        Validators.maxLength(32)])]
    });
  }

  submitCredentials() {
    this.isSubmitted = true;

    const credentials: Credentials = {
      username: this.form.value.username,
      password: this.form.value.password
    }

    this.authService.loginWith(credentials).subscribe(
      data => {
        console.log('Login response: ', data);
        this.snackBar.open('Login succeeded.', 'Close', { duration: 5000 });

        localStorage.setItem('agentAppToken', data.token);
        let tokenInfo = this.getDecodedAccessToken(data.token);

        localStorage.setItem('id', tokenInfo.id);
        localStorage.setItem('username', tokenInfo.username);
        localStorage.setItem('role', tokenInfo.role);
        localStorage.setItem('ownedCompanyId', tokenInfo.ownedCompanyId);
        localStorage.setItem('issuedCompanyRequestId', tokenInfo.issuedCompanyRequestId);
        localStorage.setItem('exp', tokenInfo.exp);

        this.router.navigateByUrl('/').then(() => { window.location.reload(); });
      },
      (error: HttpErrorResponse) => {
        this.isSubmitted = false;

        console.log('Error on login: ', error.error);
        this.snackBar.open(error.error, 'Close', { duration: 5000 });
      }
    );
  }

  getDecodedAccessToken(token: string): any {
    try {
      return jwt_decode(token);
    } catch (error) {
      return '';
    }
  }
}
