import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { AuthenticationService } from 'src/app/service/authentication.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';

@Component({
  selector: 'app-all-profiles',
  templateUrl: './all-profiles.component.html',
  styleUrls: ['./all-profiles.component.css']
})
export class AllProfilesComponent implements OnInit {

  public profiles: User[] = [];
  public results: number = 0;
  public searchText: string = "";
  isAuthenticated = false;
  private id: any;

  constructor(private _profileService: ProfileService,private authservice: AuthenticationService,
    public _router: Router,) { }

  ngOnInit(): void {
    this.getAllProfiles();
    this.isLoggedIn();
    this.id = localStorage.getItem("id");
  }
  
  isLoggedIn() : void{
    if(this.authservice.getDislinktAppToken()){
      this.isAuthenticated=true;
    }
    else{
      this.isAuthenticated=false;
    }
  }


  getAllProfiles(): void {
    this._profileService.getAllProfiles().subscribe(
      response => {
        this.profiles = response.users;
        console.log(this.profiles);
        this.results = response.users.length;
        console.log(response.length);
        console.log(typeof(this.profiles))
      }
    )
  }

  searchProfiles(): void {
    if (this.searchText === "") {
      this.getAllProfiles();
    } else {
      this._profileService.searchProfiles(this.searchText).subscribe(
        response => {
          this.profiles = response.users;
          this.results = response.users.length;
        }
      )
    }
  }

  undoSearch(): void {
    this.searchText = "";
    this.getAllProfiles();
  }

  viewFullProfile(id: string): void {
    this._router.navigate(['profile/' + id])
  }
}
