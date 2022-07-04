import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ConnectionDTO } from 'src/app/components/dto/connection.dto';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ConnectionService {

  private applicationURL = environment.apiUrl;
  constructor(private _http: HttpClient) { }

  getConnections(userId: string): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/connection/" + userId + "/following");
  }

  getConnectionsRequest(userId: string): Observable<any> {
    return this._http.get<any>(this.applicationURL + "/connection/" + userId + "/followers");
  }

  makeConnection(request: ConnectionDTO): Observable<any> {
    return this._http.post<any>(this.applicationURL + "/connection", request);
  }

  editRequest(requestId:string,request:ConnectionDTO): Observable<any>{
    return this._http.put<any>(this.applicationURL + "/connection/"+ requestId, request);
  }

  deleteConnection(id: string): Observable<any> {
    return this._http.delete<any>(this.applicationURL + "/connection/" + id);
  }
}
