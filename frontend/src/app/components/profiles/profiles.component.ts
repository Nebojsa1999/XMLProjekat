import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileService } from 'src/app/service/profile-service/profile.service';
import { User } from 'src/app/model/user'
@Component({
  selector: 'app-profiles',
  templateUrl: './profiles.component.html',
  styleUrls: ['./profiles.component.css']
})
export class ProfilesComponent implements OnInit {

  public profiles: User[] = [];
  public results: number = 0;

  constructor(private _profileService: ProfileService,
    public _router: Router,) { }

  ngOnInit(): void {
    this.getPublicProfiles();
  }


  getPublicProfiles(): void {
    this._profileService.getPublicProfiles().subscribe(
      response => {
        this.profiles = response.users;
        console.log(this.profiles);
        this.results = response.users.length;
        console.log(response.length);
        console.log(typeof(this.profiles))
      }
    )
  }


}
