import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CommentDTO } from 'src/app/components/dto/comment.dto';
import { PostDTO } from 'src/app/components/dto/post.dto';
import { ReactionDTO } from 'src/app/components/dto/reaction.dto';
import { Post } from 'src/app/model/post';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PostService {

  private applicationURL = environment.apiUrl;
  constructor(private _http: HttpClient) { }

  getPostsByUser(ownerId: string): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/user/"+ ownerId + "/post");
  }

  createPost(ownerId:string,newPost: PostDTO): Observable<any>{
    return this._http.post<any>(this.applicationURL+"/user/"+ownerId+"/post",newPost);
  }

  insertLikeOrDislike(ownerId:string,postId:string,type:string,reactionDTO:ReactionDTO):Observable<any>{
    return this._http.put<any>(`${this.applicationURL}/user/${ownerId}/post/${postId}/liked_or_disliked_by/${type}`,reactionDTO);
  }

  createComment(ownerId:string,postId:string,commentDTO:CommentDTO):Observable<any>{
    return this._http.post<any>(`${this.applicationURL}/user/${ownerId}/post/${postId}/comment`,commentDTO);
  }
}
