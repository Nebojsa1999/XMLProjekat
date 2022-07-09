import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  isAuthenticated = false;
  role : any;
  user:any;

  constructor(private router: Router,private service:AuthService) { }



  ngOnInit(): void { 
    this.isLoggedIn();
    this.user=localStorage.getItem('username');
    this.role=localStorage.getItem('role');
    console.log(this.role);
  }

  logOut() {
    localStorage.removeItem('agentAppToken');
    localStorage.removeItem('id');
    localStorage.removeItem('username');
    localStorage.removeItem('role');
    localStorage.removeItem('ownedCompanyId');
    localStorage.removeItem('issuedCompanyRequestId');
    localStorage.removeItem('exp');

    this.router.navigate(['/login']);
     if(this.service.getAgentAppToken()){
      this.isAuthenticated=true;
    }
    else{
      this.isAuthenticated=false;
    }

  }

  isLoggedIn() : void{
    if(this.service.getAgentAppToken()){
      this.isAuthenticated=true;
    }
    else{
      this.isAuthenticated=false;
    }
  }
}
