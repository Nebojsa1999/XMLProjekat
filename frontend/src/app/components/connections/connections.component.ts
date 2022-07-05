import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { ConnectionService } from 'src/app/service/connection-service/connection.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';
import { ConnectionDTO } from '../dto/connection.dto';

@Component({
  selector: 'app-connections',
  templateUrl: './connections.component.html',
  styleUrls: ['./connections.component.css']
})
export class ConnectionsComponent implements OnInit {
  private id: any;
  public connectionsId: string[]=[];
  users: User[] = [];
  numberOfConnections: number = 0;

  constructor(private _connectionService: ConnectionService,
              private _profileService: ProfileService) { }

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
            this.connectionsId.push(response.connections[i].id);
            console.log(this.connectionsId)
            this.getConnectionProfiles(response.connections[i].subjectId);
          }
        }
   
        
      }
    )
  }

  getConnectionProfiles(id: string): void {
      this._profileService.getProfile(id).subscribe(
        response => {
          console.log(response.user)
          this.users.push(response.user);
          console.log(  this.users)
         
        }
      )
    }

  deleteConnection(subjectId:string):void{
    this._connectionService.deleteConnection(this.id,subjectId).subscribe(
      response=>{
        console.log("deleted connection");
      }
    )
    window.location.reload();
  }
  }


