import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { AuthenticationService } from 'src/app/service/authentication.service';
import { RegisterDTO } from '../dto/register.dto';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  public user: RegisterDTO = {
    username:"",
    password:"",
    isPrivate:false,
    firstName:"",
    lastName:"",
    email:"",     
  }

  public confirmPassword: string = "";

  constructor(private _authenticationService: AuthenticationService,
    private _snackBar: MatSnackBar,
    private _router: Router) { }

  ngOnInit(): void {
  }

  register(): void {
    if(this.isPasswordValid()) {
      this._authenticationService.register(this.user).subscribe(
        response => {
          console.log(response);
          this._snackBar.open(response.message, "Close");
        },
        error => {
          this._snackBar.open("An error has ocurred. Please try again.", "Close");
        }
      )
    } else {
      this._snackBar.open("Passwords do not match! Please try again.", "Close");
    }
    this._router.navigate(['login']);
  }

  isPasswordValid(): boolean {
    return this.user.password === this.confirmPassword ? true : false;
  }

}
