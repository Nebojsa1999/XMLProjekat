import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
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

  constructor(private fb: FormBuilder,private _route: ActivatedRoute,private authservice: AuthenticationService,
    private _profileService: ProfileService,private postService:PostService) { 
      this.userForm = this.fb.group({
        username: ["", [Validators.required]],
        firstName: ["", [Validators.required]],
        lastName: ["", [Validators.required]],
        email: ["", [Validators.required]],
        phone: ["", [Validators.required]],
        gender: ["", [Validators.required]],
        dateOfBirth: [new Date(), [Validators.required]],
        biography: ["", [Validators.required]],
        workExperience: ["", [Validators.required]],
        education: ["", [Validators.required]],
        skills: ["", [Validators.required]],
        interests: ["", [Validators.required]],
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
    console.log(this.id);
    this._profileService.updateProfile(this.id,this.profile).subscribe(
      response => {
        this.profile = response.user;

      },
      error => {
        console.log("error on profile update",error);
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

}
