import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormFieldNomeComponent } from './form-field-nome.component';

describe('FormFieldNomeComponent', () => {
  let component: FormFieldNomeComponent;
  let fixture: ComponentFixture<FormFieldNomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormFieldNomeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FormFieldNomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
