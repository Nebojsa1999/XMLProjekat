import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthenticationService } from 'src/app/service/authentication.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  isAuthenticated = false;

  constructor(private _router: Router,private service:AuthenticationService) { }

  ngOnInit(): void {
    this.isLoggedIn();
    console.log(this.isAuthenticated);
  }

  isLoggedIn() : void{
    if(this.service.getDislinktAppToken()){
      this.isAuthenticated=true;
    }
    else{
      this.isAuthenticated=false;
    }
  }

  signOut() : void{
    this.service.logOut();
    if(this.service.getDislinktAppToken()){
      this.isAuthenticated=true;
    }
    else{
      this.isAuthenticated=false;
    }
    console.log(this.isAuthenticated);
  }

  openProfile(): void {
    this._router.navigate(['profile/' + localStorage.getItem("id")]);
    setTimeout(() => {
      window.location.reload();
    }, 200);
  }

}
