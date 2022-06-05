import { Component, OnInit } from '@angular/core';

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

  constructor(private userService: UserService) { }
  
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
        console.log('User data: ' + data);
        this.user = data;
      },
      error => {
        console.log('Error on getUserById!', error)
      }
    );
  }
}
