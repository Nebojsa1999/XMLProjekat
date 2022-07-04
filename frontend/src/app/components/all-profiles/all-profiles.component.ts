import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user';
import { AuthenticationService } from 'src/app/service/authentication.service';
import { ConnectionService } from 'src/app/service/connection-service/connection.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';
import { ConnectionDTO } from '../dto/connection.dto';

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
  isProfileOwner = false; 
  // isFollowing = false;

  private connectionDTO: ConnectionDTO = {
    issuerId : "",
    subjectId : "",
    isApproved : false,
  }

  constructor(private _profileService: ProfileService,private authservice: AuthenticationService,
    public _router: Router, private _connectionService: ConnectionService) { }

  ngOnInit(): void {
    this.getAllProfiles();
    this.isLoggedIn();
    this.id = localStorage.getItem("id");
    // this.isFollowingUser(this.id);
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

  connect(subjectId: string, isPrivate: boolean): void {
    this.connectionDTO.issuerId = this.id;
    this.connectionDTO.subjectId = subjectId;
    if (isPrivate) {
      this.connectionDTO.isApproved=false;
      this._connectionService.makeConnection(this.connectionDTO).subscribe(
        response => {
          console.log(response);
          alert("Done");
        },
        error =>{ alert("Already sent request to this profile")

        }
      )
    } else {
      this.connectionDTO.isApproved=true;
      this._connectionService.makeConnection(this.connectionDTO).subscribe(
        response => {
          console.log(response);
          alert("Done");
        },
        error =>{ alert("Already sent request to this profile")}
      )
    }
  }

  // isFollowingUser(userId: string): void {
  //   this._connectionService.getConnections(userId).subscribe(
  //     response => {
  //       console.log(response);
  //       for(let i = 0;i<response.connections.length;i++){
  //         if(response.connections[i].isApproved == true){
  //           this.isFollowing=true;
  //           console.log(this.isFollowing)
  //         }
  //       } 
  //     }
  //   )
  // }

}
