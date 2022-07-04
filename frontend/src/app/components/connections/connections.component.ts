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
  private connectionsId: string[]=[];
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
        this.connectionsId = response.connections.id;
        console.log(response.connections[0].id)
        this.getConnectionProfiles(response);
        console.log(response)
        
      }
    )
  }

  getConnectionProfiles(ids: string[]): void {
    for (let i = 0; i < ids.length; i++) {
      this._profileService.getProfile(ids[i]).subscribe(
        response => {
          this.connections.push(response);
          this.numberOfConnections = this.connections.length;
          console.log(response);
        }
      )
    }
  }

}
