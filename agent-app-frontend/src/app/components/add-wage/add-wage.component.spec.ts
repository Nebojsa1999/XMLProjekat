import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddWageComponent } from './add-wage.component';

describe('AddWageComponent', () => {
  let component: AddWageComponent;
  let fixture: ComponentFixture<AddWageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddWageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddWageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
