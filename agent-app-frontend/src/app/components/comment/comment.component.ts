import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { AddCommentComponent } from 'src/app/components/add-comment/add-comment.component';
//import { Job } from 'src/app/models/job';
//import { JobService } from 'src/app/services/job/job.service';
import { CommentService } from 'src/app/services/comment/comment.service';
import { CommentDto } from '../dto/comment.dto';
import { Comment } from 'src/app/models/comment'
import { InterviewDto } from '../dto/interview.dto';
import { WageDto } from '../dto/wage.dto'
//import { NewJobtDto } from '../dto/new-job-dto'

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.css']
})
export class CommentComponent implements OnInit {
  comments: Comment[] = [];
  allComments: Comment[] = [];

  public newComment: CommentDto = {
    id: "",
    companyId: "",
    position: "",
    engagement: "",
    experienceLevel: "",
    content: ""
  }

  constructor(private commentService: CommentService, public matDialog: MatDialog) { }

  ngOnInit(): void {
//     this.commentService.getComments().subscribe(
//       response => {
//         this.comments = response.comments;
//         console.log(this.comments);
//         this.allComments = response.comments;
//         console.log(this.allComments);
// ;
//       }
//     )
  }

  openNewCommentDialog(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.height = "450px";
    dialogConfig.width = "35%";
    const modalDialog = this.matDialog.open(AddCommentComponent, dialogConfig);
    modalDialog.afterClosed().subscribe(result => {
      location.reload()
    })
    
  }
}
