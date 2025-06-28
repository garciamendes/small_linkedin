import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PreviewMeComponent } from './preview-me.component';

describe('PreviewMeComponent', () => {
  let component: PreviewMeComponent;
  let fixture: ComponentFixture<PreviewMeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PreviewMeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PreviewMeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
