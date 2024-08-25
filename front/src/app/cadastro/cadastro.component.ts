import { Component, signal } from '@angular/core';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { RouterLink } from '@angular/router';
import { FormFieldNomeComponent } from "../form-field-nome/form-field-nome.component";
import { FormFieldEmailComponent } from '../form-field-email/form-field-email.component';
import { FormFieldSenhaComponent } from '../form-field-senha/form-field-senha.component';

@Component({
  selector: 'app-cadastro',
  standalone: true,
  imports: [RouterLink, FormsModule, ReactiveFormsModule, MatButtonModule, FormFieldNomeComponent, FormFieldEmailComponent, FormFieldSenhaComponent],
  templateUrl: './cadastro.component.html',
  styleUrl: './cadastro.component.css'
})
export class CadastroComponent {

  cadastro = new FormGroup({
    nome: new FormControl('', [Validators.required, Validators.minLength(3)]),
    email: new FormControl('', [Validators.required, Validators.email]),
    senha: new FormControl('', [Validators.required, Validators.minLength(5)]),
  });

  hideSenha = signal(true);
  showHideSenha(event: MouseEvent) {
    this.hideSenha.set(!this.hideSenha());
    event.stopPropagation();
  }

  cadastrar() {
    if(this.cadastro.invalid) {
      return;
    }

    //TODO: Post para o server
  }

}
