import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CommentDto } from 'src/app/components/dto/comment.dto';
import { NewCommentDto } from 'src/app/components/dto/new-comment.dto';
import { Comment } from 'src/app/models/comment';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CommentService {

  private applicationURL = environment.apiUrl;

  constructor(private _http: HttpClient) { }


  getComments(id:string): Observable<CommentDto[]> {
    return this._http.get<CommentDto[]>(`${this.applicationURL}/agent-app/company/comment/${id}`);
  }

  createNewComment(dto: NewCommentDto): Observable<any> {
    return this._http.post<any>(this.applicationURL + "/agent-app/job/comment/create", dto);
  }
}
