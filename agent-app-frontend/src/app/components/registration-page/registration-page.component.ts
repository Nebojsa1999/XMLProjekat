import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { MatSnackBar } from '@angular/material/snack-bar'

import { UserService } from 'src/app/services/user/user.service';

import { User } from 'src/app/models/user';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.css']
})
export class RegistrationPageComponent implements OnInit {
  form: any;
  isSubmitted: boolean = false;

  constructor(private userService: UserService, private formBuilder: FormBuilder, 
    private router: Router, private snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      username: ['', Validators.compose([Validators.required, Validators.minLength(4), 
        Validators.maxLength(32)])],
      password: ['', Validators.compose([Validators.required, Validators.minLength(6), 
        Validators.maxLength(15)])],
      firstName: ['', Validators.compose([Validators.required, Validators.minLength(2), 
        Validators.pattern(/^[A-Z\p{L}][a-z\p{L}]+([ -][A-Z\p{L}][a-z\p{L}]+)*$/u)])],
      lastName: ['', Validators.compose([Validators.required, Validators.minLength(2), 
        Validators.pattern(/^[A-Z\p{L}][a-z\p{L}]+([ -][A-Z\p{L}][a-z\p{L}]+)*$/u)])],
      email: ['', Validators.compose([Validators.required, 
        Validators.pattern(/^[a-z0-9\_\-\.]+@[a-z]+\.[a-z\.]+$/)])],
      phone: ['', Validators.compose([Validators.required, Validators.minLength(11), 
        Validators.maxLength(12), 
        Validators.pattern(/^[0-9]{3}\/[0-9]{3,4}-[0-9]{3,4}$/)])],
      gender: ['Male', Validators.compose([Validators.required])],
      dateOfBirth: ['', Validators.compose([Validators.required, Validators.minLength(10), 
        Validators.maxLength(10), 
        Validators.pattern(/^(19|20)[0-9]{2}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$/)])],
      biography: ['', Validators.compose([Validators.required])],
      workExperience: ['', Validators.compose([Validators.required])],
      education: ['', Validators.compose([Validators.required])],
      skills: ['', Validators.compose([Validators.required])],
      interests: ['', Validators.compose([Validators.required])]
    });
  }

  submitData() {
    this.isSubmitted = true;

    let dateOfBirthAsString: string = this.form.value.dateOfBirth;
    dateOfBirthAsString.concat('T00:00:00Z');

    const newCommonUser: User = {
      id: "",
      role: "CommonUser",
      ownedCompanyId: "000000000000000000000000",
      issuedCompanyRequestId: "000000000000000000000000",
      username: this.form.value.username,
      password: this.form.value.password,
      firstName: this.form.value.firstName,
      lastName: this.form.value.lastName,
      email: this.form.value.email,
      phone: this.form.value.phone,
      gender: this.form.value.gender,
      dateOfBirth: new Date(dateOfBirthAsString),
      biography: this.form.value.biography,
      workExperience: this.form.value.workExperience,
      education: this.form.value.education,
      skills: this.form.value.skills,
      interests: this.form.value.interests
    }

    this.userService.registerAsACommonUser(newCommonUser).subscribe(
      data => {
        console.log('Registration response: ', data);
        this.snackBar.open('Registration succeeded.', 'Close', { duration: 5000 });
        
        this.router.navigateByUrl('/login').then(() => { window.location.reload(); });
      },
      (error: HttpErrorResponse) => {
        this.isSubmitted = false;

        console.log('Error on registration: ', error.error);
        this.snackBar.open(error.error, 'Close', { duration: 5000 });
      }
    );
  }
}
