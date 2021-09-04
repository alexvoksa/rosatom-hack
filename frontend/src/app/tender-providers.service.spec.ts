import { TestBed } from '@angular/core/testing';

import { TenderProvidersService } from './tender-providers.service';

describe('TenderProvidersService', () => {
  let service: TenderProvidersService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TenderProvidersService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
