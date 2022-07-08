import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { CommentDto } from 'src/app/components/dto/comment.dto';
import { CommentService } from 'src/app/services/comment/comment.service';

@Component({
  selector: 'app-add-comment',
  templateUrl: './add-comment.component.html',
  styleUrls: ['./add-comment.component.css']
})
export class AddCommentComponent implements OnInit {

  public addCommentForm: FormGroup;
  public id: FormControl;
  public userId: FormControl;
  public jobId: FormControl;
  public position: FormControl;
  public engagement: FormControl;
  public experienceLevel: FormControl;
  public content: FormControl;

  constructor(private fb: FormBuilder,private commentService: CommentService,  public dialogRef: MatDialogRef<AddCommentComponent>) {
      this.id= new FormControl("", [Validators.required]);
      this.userId= new FormControl("", [Validators.required]);
      this.jobId= new FormControl("", [Validators.required]);
      this.position= new FormControl("", [Validators.required]);
      this.engagement=new FormControl("", [Validators.required]);
      this.experienceLevel= new FormControl("", [Validators.required]);
      this.content= new FormControl("", [Validators.required]);

      this.addCommentForm = new FormGroup({
        'id': this.id,
        'userId': this.userId,
        'jobId': this.jobId,
        'position': this.position,
        'engagement': this.engagement,
        'experienceLevel': this.experienceLevel,
        'content': this.content,
      })
   }

  ngOnInit(): void {
  }

  
  onAdd() {
    if (this.addCommentForm.valid) {
      let dto: CommentDto = {id: this.id.value, userId: this.userId.value, jobId: this.jobId.value, position: this.position.value, engagement: this.engagement.value, experienceLevel: this.experienceLevel.value, content: this.content.value}
      this.commentService.createNewComment(dto).subscribe((response) => {
        console.log("comment added")
        this.dialogRef.close()
      })

    }
  }


}
