import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';

import { Credentials } from 'src/app/models/credentials';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {
  form: any;
  isSubmitted: boolean = false;

  token: string = '';

  constructor(private formBuilder: FormBuilder) { }

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
  }
}
