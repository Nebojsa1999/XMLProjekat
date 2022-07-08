import { TestBed } from '@angular/core/testing';

import { WageService } from './wage.service';

describe('WageService', () => {
  let service: WageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WageService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});