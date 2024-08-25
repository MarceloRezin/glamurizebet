import { Component } from '@angular/core';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { FormFieldEmailComponent } from "../form-field-email/form-field-email.component";
import { FormFieldSenhaComponent } from "../form-field-senha/form-field-senha.component";


@Component({
  selector: 'app-login',
  standalone: true,
  imports: [RouterLink, FormsModule, ReactiveFormsModule, MatButtonModule, FormFieldEmailComponent, FormFieldSenhaComponent],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  login = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    senha: new FormControl('', [Validators.required, Validators.minLength(5)]),
  });

  logar() {
    if(this.login.invalid) {
      return;
    }

    //TODO: Post para o server
  }
}
