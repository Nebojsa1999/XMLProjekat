import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Credentials } from 'src/app/model/credentials';
import { AuthenticationService } from 'src/app/service/authentication.service';
import jwt_decode from 'jwt-decode';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: any;
  isSubmitted: boolean = false;

  dislinktAppToken: string = '';

  constructor(private authService: AuthenticationService, private formBuilder: FormBuilder, 
    private router: Router) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      username: ['', Validators.compose([Validators.required, Validators.minLength(4), 
        Validators.maxLength(32)])],
      password: ['', Validators.compose([Validators.required, Validators.minLength(6), 
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

        localStorage.setItem('dislinktAppToken', data.token);
        let tokenInfo = this.getDecodedAccessToken(data.token);

        localStorage.setItem('id', tokenInfo.id);
        localStorage.setItem('username', tokenInfo.username);
        localStorage.setItem('exp', tokenInfo.exp);

        this.router.navigateByUrl('/').then(() => { window.location.reload(); });
      },
      error => {
        this.isSubmitted = false;

        console.log('Error on login: ', error);
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