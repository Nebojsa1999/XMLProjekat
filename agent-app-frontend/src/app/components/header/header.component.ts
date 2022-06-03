import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  constructor(private router: Router) { }

  ngOnInit(): void {}

  logOut() {
    localStorage.removeItem('agentAppToken');
    localStorage.removeItem('id');
    localStorage.removeItem('username');
    localStorage.removeItem('role');
    localStorage.removeItem('ownedCompanyId');
    localStorage.removeItem('issuedCompanyRequestId');
    localStorage.removeItem('exp');

    this.router.navigate(['/login']);
  }
}
