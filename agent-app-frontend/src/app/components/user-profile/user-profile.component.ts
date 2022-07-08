import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

import { MatSnackBar } from '@angular/material/snack-bar';

import { UserService } from 'src/app/services/user/user.service';

import { User } from 'src/app/models/user';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {
  user: User = {
    id: '',
    role: '',
    ownedCompanyId: '',
    issuedCompanyRequestId: '',
    username: '',
    password: '',
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    gender: '',
    dateOfBirth: new Date(),
    biography: '',
    workExperience: '',
    education: '',
    skills: '',
    interests: ''
  }

  constructor(private userService: UserService, private snackBar: MatSnackBar) { }
  
  ngOnInit(): void {
    this.getUser();
  }

  getUser() {
    let id = localStorage.getItem('id');

    if (id == undefined || id == null) {
      return;
    }

    this.userService.getUserWith(id).subscribe(
      data => {
        console.log('User data: ' + JSON.stringify(data));
        this.snackBar.open('Retrieving logged user\'s data succeeded.', 'Close', { duration: 5000 });

        this.user = data;
      },
      (error: HttpErrorResponse) => {
        console.log('Error on get user by id: ', error.error);
        this.snackBar.open(error.error, 'Close', { duration: 5000 });
      }
    );
  }
}
