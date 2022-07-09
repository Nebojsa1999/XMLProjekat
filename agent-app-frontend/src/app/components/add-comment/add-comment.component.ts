import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { CommentDto } from 'src/app/components/dto/comment.dto';
import { CommentService } from 'src/app/services/comment/comment.service';
import { NewCommentDto } from '../dto/new-comment.dto';

export interface DialogData {
  companyId: string;
}

@Component({
  selector: 'app-add-comment',
  templateUrl: './add-comment.component.html',
  styleUrls: ['./add-comment.component.css']
})
export class AddCommentComponent implements OnInit {
  public addCommentForm: FormGroup;

  public position: FormControl;
  public engagement: FormControl;
  public experienceLevel: FormControl;
  public content: FormControl;

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData,private fb: FormBuilder,private commentService: CommentService,  public dialogRef: MatDialogRef<AddCommentComponent>) {
      this.position= new FormControl("", [Validators.required]);
      this.engagement=new FormControl("", [Validators.required]);
      this.experienceLevel= new FormControl("", [Validators.required]);
      this.content= new FormControl("", [Validators.required]);

      this.addCommentForm = new FormGroup({
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
      let dto: NewCommentDto = {companyId: this.data.companyId, position: this.position.value, engagement: this.engagement.value, experienceLevel: this.experienceLevel.value, content: this.content.value}
      this.commentService.createNewComment(dto).subscribe((response) => {
        console.log("comment added")
        this.dialogRef.close()
      })

    }
  }


}
