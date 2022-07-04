import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { ConnectionService } from 'src/app/service/connection-service/connection.service';
import { ProfileService } from 'src/app/service/profile-service/profile.service';

@Component({
  selector: 'app-connections',
  templateUrl: './connections.component.html',
  styleUrls: ['./connections.component.css']
})
export class ConnectionsComponent implements OnInit {
  private id: any;
  public connectionsId: string="";
  connections: User[] = [];
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
            this.connectionsId = response.connections[i].id;
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
          this.connections.push(response.user);
          console.log(  this.connections)
         
        }
      )
    }

  deleteConnection(id:string):void{
    this._connectionService.deleteConnection(id).subscribe(
      response=>{
        console.log("deleted connection");
      }
    )
    window.location.reload();
  }
  }


