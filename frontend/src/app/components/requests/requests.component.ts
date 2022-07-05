import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { ConnectionService } from 'src/app/service/connection-service/connection.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';
import { ConnectionDTO } from '../dto/connection.dto';

@Component({
  selector: 'app-requests',
  templateUrl: './requests.component.html',
  styleUrls: ['./requests.component.css']
})
export class RequestsComponent implements OnInit {
  private id: any;
  public connectionsId: string="";
  connections: User[] = [];
  numberOfConnections: number = 0;
  private requestDTO: ConnectionDTO = {
    issuerId : "",
    subjectId : "",
    isApproved : false
  }


  constructor(private _connectionService: ConnectionService,
              private _profileService: ProfileService) { }

  ngOnInit(): void {
    this.id = localStorage.getItem("id");
    this.getConnections(this.id);
  }

  getConnections(userId: string): void {
    this._connectionService.getConnectionsRequest(userId).subscribe(
      response => {
        console.log(response);
        for(let i = 0;i<response.connections.length;i++){
          if(response.connections[i].isApproved == false){
            this.connectionsId = response.connections[i].id;
            console.log(this.connectionsId)
            this.getConnectionProfiles(response.connections[i].issuerId);
          }
        }
   
        
      }
    )
  }

  getConnectionProfiles(id: string): void {
      this._profileService.getProfile(id).subscribe(
        response => {
          console.log(response.user)
          this.connections.push(response.user);
          console.log(  this.connections)
         
        }
      )
    }

    acceptRequest(senderId:string): void {
      this.requestDTO.isApproved=true;
      this.requestDTO.issuerId=senderId;
      this.requestDTO.subjectId=this.id;
      this._connectionService.editRequest(this.requestDTO).subscribe(
        response => {
          console.log(response);
        }
      )
      window.location.reload();
  }

  declineRequest(senderId:string): void {
    this._connectionService.deleteConnection(senderId,this.id).subscribe(
      response => {
        console.log(response);
      }
    )
    window.location.reload();
}

}

