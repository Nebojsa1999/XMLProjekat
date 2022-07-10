import { Component, OnInit } from '@angular/core';
import { Post } from 'src/app/model/post';
import { User } from 'src/app/model/user';
import { UserDataAndPostWrapper } from 'src/app/model/userDataAndPostWrapper';
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
  public wrappers: UserDataAndPostWrapper[] = [];
  // public currentWrapper: UserDataAndPostWrapper = {
  //   user: {
  //     id: "",
  //     username: "",
  //     password: "",
  //     isPrivate: false,
  //     firstName: "",
  //     lastName: "",
  //     email: "",
  //     phone: "",
  //     gender: "",
  //     dateOfBirth: new Date(1/1/1999),
  //     biography: "",
  //     workExperience: "",
  //     education: "",          
  //     skills: "",          
  //     interests: "",      
  //   },
  //   post: {
  //     id: "",
  //     ownerId: "",
  //     content: "",
  //     image: "",
  //     likesCount: 0,
  //     dislikesCount: 0,
  //     comments: [],
  //     links: [],
  //     whoLiked: [],
  //     whoDisliked: [],
  //     postedAt: new Date(1/1/1999)
  //   }
  // }

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
    link:[""]         
  }
  
  selectedImage: any;
  selectedImageName: string = "undefined";

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
        console.log("Zapraceni prijavljenog korisnika:\n" + JSON.stringify(response));
        for(let i = 0;i<response.connections.length;i++){
          if(response.connections[i].isApproved == true){
            this.getPostsOfFollowingUsers(response.connections[i].subjectId);
          }
        }
      }
    )
  }

  getPostsOfFollowingUsers(userId: string): void {
    this._postService.getPostsByUser(userId).subscribe(
      response => {
        console.log(response)
        for(let i = 0;i<response.posts.length;i++){
          this.posts.push(response.posts[i]);
         // this.posts[i]=findLinksInContentOfPost(this.posts[i]);
          this.setUserToPost(this.posts);

        }
      }
    )
  }

  // findLinksInContentOfPost(post:Post): Post{
  //   //let match=new RegExp(/^\"(http|www)\"$/)
  //   let match=new RegExp(/^(http|ftp|https):\/\/([\w-]+(?:(?:.[\w-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])$/)
  //   if(match.test(post.content)){
  //     post.
  //   }
  // }

  setUserToPost(posts: Post[]): void {
    for(let i = 0; i < posts.length; i++) {
      this._profileService.getProfile(posts[i].ownerId).subscribe(
        response => {
          posts[i].user = response;
        }
      )
    }
  }


  createPost(): void {
    this.newPost.ownerId = this.id;
    this._postService.createPost(this.id,this.newPost).subscribe(
      response => {
        console.log(response);
        alert("Post has been successfully created.");
      },
      error => {
        alert("Error occurred, post has not been created.");
      }
    )
  }

  onImageSelected(event: any) {
    if (!event || !event.target) {
      return;
    }

    if (!event.target.files) {
      this.selectedImage = null;
    }

    this.selectedImage = <File>event.target.files[0];
    this.selectedImageName = event.target.files[0].name;
    let reader = new FileReader();
    reader.readAsDataURL(this.selectedImage);
    var self = this;

    reader.onload = function() {
      self.selectedImage = reader.result;
      self.newPost.image = self.selectedImageName;
    }
    reader.onerror = function() {
      console.log("Error occured while converting image!");
    }
  }

  createComment(ownerId:string,postId: string): void {
    console.log(this.commentDTO)
    console.log(postId)
    this._postService.createComment(ownerId,postId,this.commentDTO).subscribe(
      response => {
        console.log(response);
      }
    )
    window.location.reload();
  }

  insertLikeOrDislike(ownerId:string,postId:string,type:string):void{

    this.reactionDTO.by_user_id=this.id;
    console.log(this.reactionDTO);

    this._postService.insertLikeOrDislike(ownerId,postId,type,this.reactionDTO).subscribe(
      response => {
        console.log(response);
        window.location.reload();
      },
      error => {
        alert("Already liked/disliked")
      }
      
    )
    
  }

  
  ParseComment(comment:any):string{
    return JSON.parse(JSON.stringify(comment)).content;
  }

  ParseFirstName(user:any):string{
    return JSON.parse(JSON.stringify(user)).user.firstName;
  }

  ParseLastName(user:any):string{
    return JSON.parse(JSON.stringify(user)).user.lastName;
  }



}
