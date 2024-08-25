import { Component, Input } from '@angular/core';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-form-field-email',
  standalone: true,
  imports: [ MatFormFieldModule, MatInputModule, FormsModule, ReactiveFormsModule ],
  templateUrl: './form-field-email.component.html',
  styleUrl: './form-field-email.component.css'
})
export class FormFieldEmailComponent {

  @Input() formControl!: FormControl;

}
