import { Component, Input, signal } from '@angular/core';
import { FormControl, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-form-field-senha',
  standalone: true,
  imports: [ MatFormFieldModule, MatInputModule, FormsModule, ReactiveFormsModule, MatButtonModule, MatIconModule ],
  templateUrl: './form-field-senha.component.html',
  styleUrl: './form-field-senha.component.css'
})
export class FormFieldSenhaComponent {
  @Input() formControl!: FormControl;
  @Input() autocomplete!: String;

  hideSenha = signal(true);
  showHideSenha(event: MouseEvent) {
    this.hideSenha.set(!this.hideSenha());
    event.stopPropagation();
  }
}
