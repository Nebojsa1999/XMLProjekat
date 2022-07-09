import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileService } from 'src/app/service/profile-service/profile.service';
import { User } from 'src/app/model/user'
import { AuthenticationService } from 'src/app/service/authentication.service';
@Component({
  selector: 'app-profiles',
  templateUrl: './profiles.component.html',
  styleUrls: ['./profiles.component.css']
})
export class ProfilesComponent implements OnInit {

  public profiles: User[] = [];
  public results: number = 0;
  public searchText: string = "";
  isAuthenticated = false;
  private id: any;


  constructor(private _profileService: ProfileService,private authservice: AuthenticationService,
    public _router: Router,) { }

  ngOnInit(): void {
    this.getPublicProfiles();
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


  getPublicProfiles(): void {
    this._profileService.getPublicProfiles().subscribe(
      response => {
        this.profiles = response.users;
     //   console.log(this.profiles);
        this.results = response.users.length;
    //    console.log(response.length);
    //    console.log(typeof(this.profiles))
      }
    )
  }

  searchProfiles(): void {
    if (this.searchText === "") {
      this.getPublicProfiles();
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
    this.getPublicProfiles();
  }

  viewFullProfile(id: string): void {
    this._router.navigate(['profile/' + id])
  }



}
