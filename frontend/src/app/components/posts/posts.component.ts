import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';
import { User } from 'src/app/model/user';
import { ConnectionService } from 'src/app/service/connection-service/connection.service';
import { PostService } from 'src/app/service/post-service/post.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';
import { CommentDTO } from '../dto/comment.dto';
import { PostDTO } from '../dto/post.dto';
import { ReactionDTO } from '../dto/reaction.dto';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {

  public posts: Post[] = [];
  public users: User[]=[];
  private id: any;
  like:string="like";
  dislike:string="dislike";

  public newPost: PostDTO={
    id:"",           
    ownerId:"",      
    content:"",      
    image:"",         
    links:[]          
  }

  private reactionDTO: ReactionDTO={
    by_user_id:"",
  }

  public commentDTO: CommentDTO={
    code:"",
    content:"",
  }

  constructor(private _postService: PostService,
    private _profileService: ProfileService,private _connectionService:ConnectionService) { }

  ngOnInit(): void {
    this.id = localStorage.getItem("id");
    this.getConnections(this.id);
  }


  getConnections(userId: string): void {
    this._connectionService.getConnections(userId).subscribe(
      response => {
        console.log(response);
        for(let i = 0;i<response.connections.length;i++){
          if(response.connections[i].isApproved == true){
            this.getConnectionPosts(response.connections[i].subjectId);
          }
        }    
      }
    )
  }

  getConnectionPosts(id: string): void {
    this._postService.getPostsByUser(id).subscribe(
      response => {
        for(let i = 0;i<response.posts.length;i++){
          this.posts.push(response.posts[i]);
          // this.getProfilesPosts(response.posts[i].ownerId);
        }
        console.log(this.posts);
      
      }
    )
  }

  // getProfilesPosts(id:string):void{
  //   this._profileService.getProfile(id).subscribe(
  //     response => {
  //       console.log(response)
  //       for(let i=0;i<response.users.length;i++){
  //         this.users.push(response.users[i]);
  //       }
    
  //     }
  //   )
  // }

  createPost(): void {
    this.newPost.ownerId = this.id;
    this.newPost.links.push("https://wwww.google.com");
    console.log(this.newPost)
    this._postService.createPost(this.id,this.newPost).subscribe(
      response => {
        console.log(response);
      }
    )
  }

  createComment(ownerId:string,postId: string): void {
    console.log(this.commentDTO)
    console.log(postId)
    this._postService.createComment(ownerId,postId,this.commentDTO).subscribe(
      response => {
        console.log(response);
      }
    )
  }

  insertLikeOrDislike(ownerId:string,postId:string,type:string):void{

    this.reactionDTO.by_user_id=this.id;
    console.log(this.reactionDTO);

    this._postService.insertLikeOrDislike(ownerId,postId,type,this.reactionDTO).subscribe(
      response => {
        console.log(response);
      },
      error => {
        alert("Already liked/disliked")
      }
      
    )
  }



}
