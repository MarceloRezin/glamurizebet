import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormFieldSenhaComponent } from './form-field-senha.component';

describe('FormFieldSenhaComponent', () => {
  let component: FormFieldSenhaComponent;
  let fixture: ComponentFixture<FormFieldSenhaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormFieldSenhaComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FormFieldSenhaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
