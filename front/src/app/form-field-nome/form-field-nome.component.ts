import { Component, Input } from '@angular/core';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-form-field-nome',
  standalone: true,
  imports: [ MatFormFieldModule, MatInputModule, FormsModule, ReactiveFormsModule ],
  templateUrl: './form-field-nome.component.html',
  styleUrl: './form-field-nome.component.css'
})
export class FormFieldNomeComponent {

  @Input() formControl!: FormControl;

}
