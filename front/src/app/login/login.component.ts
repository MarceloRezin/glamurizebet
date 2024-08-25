import { Component, signal } from '@angular/core';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import {MatIconModule} from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { FormFieldEmailComponent } from "../form-field-email/form-field-email.component";


@Component({
  selector: 'app-login',
  standalone: true,
  imports: [RouterLink, MatFormFieldModule, MatInputModule, FormsModule, ReactiveFormsModule, MatIconModule, MatButtonModule, FormFieldEmailComponent],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  login = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    senha: new FormControl('', [Validators.required, Validators.minLength(5)]),
  });

  hideSenha = signal(true);
  showHideSenha(event: MouseEvent) {
    this.hideSenha.set(!this.hideSenha());
    event.stopPropagation();
  }

  logar() {
    if(this.login.invalid) {
      return;
    }

    //TODO: Post para o server
  }
}
