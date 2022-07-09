import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import { JobTokenDialogComponent } from 'src/app/modal/job-token-dialog/job-token-dialog.component';
import { Post } from 'src/app/model/post';
import { User } from 'src/app/model/user';
import { AuthenticationService } from 'src/app/service/authentication.service';
import { PostService } from 'src/app/service/post-service/post.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  private id: any;
  isProfileOwner = false; 
  isAuthenticated = false;
  profilePrivacy=false;
  profilePassword="";
 
  profile : User = {
    id:"",
    username:"",
    password:"",
    isPrivate:false,
    firstName:"",
    lastName:"",
    email:"",
    phone:"",
    gender:"",
    dateOfBirth: new Date(1/1/1999),
    biography:"",
    workExperience:"",
    education:"",          
    skills:"",          
    interests:"",      
  }

  userForm: FormGroup;

  posts: Post[] = [];

  constructor(public matDialog: MatDialog,private fb: FormBuilder,private _route: ActivatedRoute,private authservice: AuthenticationService,
    private _profileService: ProfileService,private postService:PostService) { 
      this.userForm = this.fb.group({
        username: ["", [Validators.required]],
        firstName: ["", [Validators.required]],
        lastName: ["", [Validators.required]],
        email: ["", [Validators.required]],
        phone: ["", []],
        gender: ["", []],
        dateOfBirth: [new Date(), []],
        biography: ["", []],
        workExperience: ["", []],
        education: ["", []],
        skills: ["", []],
        interests: ["", []],
      });
    }

  ngOnInit(): void {
    this.id = this._route.snapshot.url[1].path;
    this.isLoggedIn();
    this.isProfileOwner = this.checkIfIsOwner(this.id);
    this.getProfile(this.id);
    this.getUserPosts(this.id);
   
   
  }

  getProfile(id: string): void {
    this._profileService.getProfile(id).subscribe(
      response => {
        this.profile = response.user;
        this.profilePrivacy=response.user.isPrivate;
        this.profilePassword=response.user.password;
        console.log(this.profilePrivacy);
        console.log(this.profilePassword);
        this.userForm.patchValue(this.profile);
        console.log(this.profile)
      }
    )
  }

  isLoggedIn() : void{
    if(this.authservice.getDislinktAppToken()){
      this.isAuthenticated=true;
    }
    else{
      this.isAuthenticated=false;
    }
  }

  checkIfIsOwner(id: string): boolean {
    return id === localStorage.getItem("id") ? true : false;
  }

  onSubmit(): void {
    this.profile=this.userForm.value;
    this.profile.id=this._route.snapshot.url[1].path;
    this.profile.password=this.profilePassword;
    this.profile.isPrivate=this.profilePrivacy;

    console.log(this.id);
    this._profileService.updateProfile(this.id,this.profile).subscribe(
      response => {
        this.profile = response.user;

      },
      error => {
        console.log("error on profile update",error);
        window.location.reload();
        alert("Error username taken")
      }
    )


  }

  getUserPosts(ownerId: string): void {
    this.postService.getPostsByUser(ownerId).subscribe(
      response => {
        this.posts = response.posts;
        console.log(this.posts);
      }
    )
  }

  generateTokenDialog():void{
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = "add-interests-modal";
    dialogConfig.height = "32%";
    dialogConfig.width = "22%";
    const modalDialog = this.matDialog.open(JobTokenDialogComponent, dialogConfig);
  }

  pc(comment:any):string{
    return JSON.parse(JSON.stringify(comment)).content;
  }

 
  

}
