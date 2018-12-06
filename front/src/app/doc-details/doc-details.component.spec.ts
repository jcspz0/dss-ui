import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DocDetailsComponent } from './doc-details.component';

describe('DocDetailsComponent', () => {
  let component: DocDetailsComponent;
  let fixture: ComponentFixture<DocDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DocDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DocDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
