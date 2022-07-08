import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CommentDto } from 'src/app/components/dto/comment.dto';
import { Comment } from 'src/app/models/comment';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CommentService {

  private applicationURL = environment.apiUrl;

  constructor(private _http: HttpClient) { }


  getComments(): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/agent-app/job/comment");
  }

  createNewComment(dto: CommentDto) {
    return this._http.post(this.applicationURL + "/agent-app/job/comment/create", dto);
  }
}
