import { ComponentFixture, TestBed } from '@angular/core/testing';

import { JobTokenDialogComponent } from './job-token-dialog.component';

describe('JobTokenDialogComponent', () => {
  let component: JobTokenDialogComponent;
  let fixture: ComponentFixture<JobTokenDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ JobTokenDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(JobTokenDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
